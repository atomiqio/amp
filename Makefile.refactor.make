SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})

# =============================================================================
# BUILD MANAGEMENT
# Variables declared here are used by this Makefile *and* are exported to
# override default values used by supporting scripts in the hack directory
# =============================================================================
export UG := $(shell echo "$$(id -u):$$(id -g)")

export VERSION := $(shell cat VERSION)
export BUILD := $(shell git rev-parse HEAD | cut -c1-8)
export LDFLAGS := "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

export OWNER := appcelerator
export REPO := github.com/$(OWNER)/amp

# COMMON DIRECTORIES
# =============================================================================
CMDDIR := cmd

# =============================================================================
# DEFAULT TARGET
# =============================================================================
all: build

# =============================================================================
# VENDOR MANAGEMENT (GLIDE)
# =============================================================================
GLIDETARGETS := vendor

$(GLIDETARGETS): glide.yaml
	@glide install || (rm -rf vendor; exit 1)
# TODO: temporary fix for trace conflict, remove when resolved
	@rm -rf vendor/github.com/docker/docker/vendor/golang.org/x/net/trace

install-deps: $(GLIDETARGETS)

.PHONY: update-deps
update-deps:
	@glide update
# TODO: temporary fix for trace conflict, remove when resolved
	@rm -rf vendor/github.com/docker/docker/vendor/golang.org/x/net/trace

.PHONY: clean-deps
clean-deps:
	@rm -rf vendor

.PHONY: cleanall-deps
# cleanall-deps will effectively causes `install-deps` to behave like `update-deps`
cleanall-deps: clean-deps
	@rm -rf .glide glide.lock

# =============================================================================
# PROTOC (PROTOCOL BUFFER COMPILER)
# Generate *.pb.go, *.pb.gw.go files in any non-excluded directory
# with *.proto files.
# =============================================================================
PROTODIRS := api cmd data tests
PROTOFILES := $(shell find $(PROTODIRS) -type f -name '*.proto')
PROTOGWFILES := $(shell find $(PROTODIRS) -type f -name '*.proto' -exec grep -l 'google.api.http' {} \;)
# Generate swagger.json files for protobuf types even if only exposed over gRPC, not REST API
PROTOTARGETS := $(PROTOFILES:.proto=.pb.go) $(PROTOGWFILES:.proto=.pb.gw.go) $(PROTOFILES:.proto=.swagger.json)

PROTOOPTS := \
	-I/go/src/ \
	-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:/go/src/ \
	--grpc-gateway_out=logtostderr=true:/go/src \
	--swagger_out=logtostderr=true:/go/src/

%.pb.go %.pb.gw.go %.swagger.json: %.proto
	@echo $<
	@protoc $(PROTOOPTS) /go/src/$(REPO)/$<

protoc: $(PROTOTARGETS)

.PHONY: clean-protoc
clean-protoc:
	@find $(PROTODIRS) \( -name "*.pb.go" -o -name "*.pb.gw.go" -o -name "*.swagger.json" \) -type f -delete

# =============================================================================
# CLEAN
# =============================================================================
.PHONY: clean cleanall
# clean doesn't remove the vendor directory since installing is time-intensive;
# you can do this explicitly: `ampmake clean-deps clean`
clean: clean-protoc clean-cli clean-server
cleanall: clean cleanall-deps

# =============================================================================
# BUILD
# =============================================================================
# When running in the amptools container, set DOCKER_CMD="sudo docker"
DOCKER_CMD ?= "docker"

build: install-deps protoc build-server build-cli

# =============================================================================
# BUILD CLI (`amp`)
# Saves binary to `cmd/amp/amp.alpine`, then builds `appcelerator/amp` image
# =============================================================================
AMP := amp
AMPBINARY=$(AMP).alpine
AMPTAG := local
AMPIMG := appcelerator/$(AMP):$(AMPTAG)
AMPBOOTDIR := bootstrap
AMPBOOTEXE := bootstrap
AMPBOOTIMG := appcelerator/$(AMP)-bootstrap:$(AMPTAG)
AMPTARGET := $(CMDDIR)/$(AMP)/$(AMPBINARY)
AMPDIRS := api/client $(CMDDIR)/$(AMP) tests
AMPSRC := $(shell find $(AMPDIRS) -type f -name '*.go')

$(AMPTARGET): $(CMDDIR)/$(AMP)/Dockerfile $(GLIDETARGETS) $(PROTOTARGETS) $(AMPSRC) $(AMPBOOTDIR)/$(AMPBOOTEXE)
	@go build -ldflags $(LDFLAGS) -o $(AMPTARGET) $(REPO)/$(CMDDIR)/$(AMP)

build-bootstrap: $(AMPBOOTDIR)/Dockerfile $(AMPBOOTDIR)/$(AMPBOOTEXE)
	@$(DOCKER_CMD) build -t $(AMPBOOTIMG) $(AMPBOOTDIR)

build-cli: $(AMPTARGET) build-bootstrap
	@$(DOCKER_CMD) build -t $(AMPIMG)  $(CMDDIR)/$(AMP) || (rm -f $(AMPTARGET); exit 1)

rebuild-cli: clean-cli build-cli

.PHONY: clean-cli
clean-cli:
	@rm -f $(AMPTARGET)
	@$(DOCKER_CMD) image rm $(AMPIMG) $(AMPBOOTIMG) || true

xbuild-cli:
	@hack/xbuild $(REPO)/bin $(AMP) $(REPO)/$(CMDDIR)/$(AMP) $(LDFLAGS)

build-cli-wrapper:
#	@hack/build4alpine $(REPO)/bin $(AMP) $(REPO)/$(CMDDIR)/$(AMP) $(LDFLAGS)
	@hack/xbuild $(REPO)/bin $(AMP) $(REPO)/$(CMDDIR)/ampwrapper

# =============================================================================
# BUILD SERVER (`amplifier`)
# Saves binary to `cmd/amplifier/amplifier.alpine`,
# then builds `appcelerator/amplifier` image
# =============================================================================
AMPL := amplifier
AMPLBINARY=$(AMPL).alpine
AMPLTAG := local
AMPLIMG := appcelerator/$(AMPL):$(AMPLTAG)
AMPLTARGET := $(CMDDIR)/$(AMPL)/$(AMPLBINARY)
AMPLDIRS := cmd/$(AMPL) api data tests
AMPLSRC := $(shell find $(AMPLDIRS) -type f -name '*.go')

$(AMPLTARGET): $(GLIDETARGETS) $(PROTOTARGETS) $(AMPLSRC)
	@go build -ldflags $(LDFLAGS) -o $(AMPLTARGET) $(REPO)/$(CMDDIR)/$(AMPL)

build-server: $(AMPLTARGET)
	@cp -f /root/.config/amp/amplifier.yaml cmd/amplifier &> /dev/null || touch cmd/amplifier/amplifier.yaml
	@$(DOCKER_CMD) build -t $(AMPLIMG) $(CMDDIR)/$(AMPL) || (rm -f $(AMPLTARGET); exit 1)
	@rm -f cmd/amplifier/amplifier.yaml

rebuild-server: clean-server build-server

.PHONY: clean-cli
clean-server:
	@rm -f $(AMPLTARGET)


dump:
	@echo $(DOCKER_CMD)

