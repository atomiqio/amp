package cli

import (
	"fmt"
	"github.com/appcelerator/amp/api/auth"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var (
	ampConfigFolder = filepath.Join(".config", "amp")
	ampTokenFile    = "token"
)

// SaveToken saves the authentication token to file
func SaveToken(metadata metadata.MD) error {
	// Extract token from header
	tokens := metadata[auth.TokenKey]
	if len(tokens) == 0 {
		return fmt.Errorf("invalid token")
	}
	token := tokens[0]
	if token == "" {
		return fmt.Errorf("invalid token")
	}

	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("cannot get current user")
	}
	if err := os.MkdirAll(filepath.Join(usr.HomeDir, ampConfigFolder), os.ModePerm); err != nil {
		return fmt.Errorf("cannot create folder")
	}
	if err := ioutil.WriteFile(filepath.Join(usr.HomeDir, ampConfigFolder, ampTokenFile), []byte(token), 0600); err != nil {
		return fmt.Errorf("cannot write token")
	}
	return nil
}

// ReadToken reads the authentication token from file
func ReadToken() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("cannot get current user")
	}
	data, err := ioutil.ReadFile(filepath.Join(usr.HomeDir, ampConfigFolder, ampTokenFile))
	if err != nil {
		return "", fmt.Errorf("cannot read token")
	}
	return string(data), nil
}

// GetLoginCredentials returns the login credentials
func GetLoginCredentials() *auth.LoginCredentials {
	token, err := ReadToken()
	if err != nil {
		return &auth.LoginCredentials{Token: ""}
	}
	return &auth.LoginCredentials{Token: token}
}
