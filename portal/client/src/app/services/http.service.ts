import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import { Subject } from 'rxjs/Subject';
import { Observable } from 'RxJS/Rx';
import { User } from '../models/user.model';
import { Team } from '../models/team.model';
import { Organization } from '../models/organization.model';
import { Member } from '../models/member.model';
import { TeamResource } from '../models/team-resource.model';
import { DockerStack } from '../docker-stacks/models/docker-stack.model';
import { DockerService } from '../docker-stacks/models/docker-service.model';
import { DockerContainer } from '../docker-stacks/models/docker-container.model';
import { StatsRequest } from '../metrics/models/stats-request.model';
import { GraphHistoricData } from '../metrics/models/graph-historic-data.model';
import * as d3 from 'd3';
import 'rxjs/add/operator/retrywhen';
import 'rxjs/add/operator/scan';
import 'rxjs/add/operator/delay';

const httpRetryDelay = 200
const httpRetryNumber = 3

@Injectable()
export class HttpService {
  private token = ""
  onHttpError = new Subject();
  parseTime = d3.timeParse("%Y-%m-%dT%H:%M:%S");
  addr = "http://localhost:8080/v1"

  constructor(private http : Http) {}

  users() {
    return this.httpGet("/users")
      .map((res : Response) => {
        const data = res.json()
        let list : User[] = []
        for (let item of data.users) {
          let user = new User(item.name, item.email, "User")
          user.verified = item.is_verified
          list.push(user)
        }
        return list
      }
    );
  }

  createOrganization(org : Organization) {
    return this.httpPost("/organizations", {name: org.name, email: org.email});
  }

  deleteOrganization(org : Organization) {
    return this.httpDelete("/organizations/"+org.name);
  }

  addUserToOrganization(org : Organization, member : Member) {
    return this.httpPost("/organizations/"+org.name+"/members", {organization_name: org.name, user_name: member.userName});
  }

  removeUserFromOrganization(org : Organization, member : Member) {
    return this.httpDelete("/organizations/"+org.name+"/members/"+member.userName);
  }

  createTeam(org : Organization, team : Team) {
    return this.httpPost("/organizations/"+org.name+"/teams", {organization_name: org.name, team_name: team.name});
  }

  deleteTeam(org : Organization, team : Team) {
    return this.httpDelete("/organizations/"+org.name+"/teams/"+team.name);
  }

  addUserToTeam(org : Organization, team : Team, member : Member) {
    return this.httpPost("/organizations/"+org.name+"/teams/"+team.name+"/members", {organization_name: org.name, team_name: team.name, user_name: member.userName});
  }

  removeUserFromTeam(org : Organization, team : Team, member : Member) {
    return this.httpDelete("/organizations/"+org.name+"/teams/"+team.name+"/members/"+member.userName);
  }

  userOrganization(user : User) {
    //return this.http.get(this.addr+"/users/"+user.name+"/organizations", { headers: this.setHeaders() })
    return this.httpGet("/users/"+user.name+"/organizations")
      .map((res : Response) => {
        const data = res.json()
        //console.log("data")
        //console.log(data)
        let list : Organization[] = []
        for (let org of data.organizations) {
          let newOrg = new Organization(org.name, org.email)
          if (org.members) {
            for (let mem of org.members) {
              newOrg.members.push(new Member(mem.name, mem.role))
            }
          }
          if (org.teams) {
            for (let team of org.teams) {
              let newTeam = new Team(team.name)
              for (let mname of team.members) {
                newTeam.members.push(new Member(mname, 0))
              }
              newOrg.teams.push(newTeam)
            }
          }
          list.push(newOrg)
        }
        console.log(list)
        return list
      }
    )
  }

  login(username : string, pwd : string) {
    return this.httpPost("/login", {name: username, password: pwd})
  }

  stacks() {
    return this.httpGet("/stacks")
    .map((res : Response) => {
      const data = res.json()
      let list : DockerStack[] = []
      if (data.entries) {
        for (let item of data.entries) {
          let stack = new DockerStack(
            item.stack.id,
            item.stack.name,
            item.services,
            item.stack.owner.name,
            item.stack.owner.type
          )
          list.push(stack)
        }
      }
      return list
    })
  }

  services(stackName : string) {
    return this.httpGet("/stacks/"+stackName+"/services")
    .map((res : Response) => {
      const data = res.json()
      let list : DockerService[] = []
      if (data.services) {
        for (let item of data.services) {
          if (item.id) {
            let serv = new DockerService(
              item.id,
              item.name,
              item.mode,
              item.replicas,
              item.imge
            )
            list.push(serv)
          }
        }
      }
      return list
    })
  }

  tasks(serviceId : string) {
    return this.httpGet("/tasks/"+serviceId)
    .map((res : Response) => {
      const data = res.json()
      let list : DockerContainer[] = []
      if (data.tasks) {
        for (let item of data.tasks) {
          if (item.id) {
            let cont = new DockerContainer(
              item.id,
              item.image,
              item.state,
              item.desired_state,
              item.node_id
            )
            list.push(cont)
          }
        }
      }
      return list
    })
  }

  organizationRessources() {
    return this.httpGet("/...")
    .map((res : Response) => {
      let list : Organization[] = []
      //
      return list
    })
  }

  stats(request : StatsRequest) {
    return this.httpPost("/stats", request)
    .map((res : Response) => {
      let data = res.json()
      //console.log(data)
      let list : GraphHistoricData[] = []
      if (data.entries) {
        for (let item of data.entries) {
          let datal : { [name:string]: number; } = {}
          if (request.stats_cpu) {
            this.setValue(datal, 'cpu-usage', item.cpu.total_usage, 1, 1)
          }
          if (request.stats_io) {
            this.setValue(datal, 'io-total', item.io.total, 1, 1)
            this.setValue(datal, 'io-write', item.io.write, 1, 1)
            this.setValue(datal, 'io-read', item.io.read, 1, 1)
          }
          if (request.stats_mem) {
            this.setValue(datal, 'mem-limit', item.mem.limit, 1, 1)
            this.setValue(datal, 'mem-maxusage', item.mem.maxusage, 1, 1)
            this.setValue(datal, 'mem-usage', item.mem.usage, 1, 1024*1024)
            this.setValue(datal, 'mem-usage-p', item.mem.usage_p, 1, 1)
          }
          if (request.stats_net) {
            this.setValue(datal, 'net-rx-bytes', item.net.rx_bytes, 1, 1)
            this.setValue(datal, 'net-rx-packets', item.net.rx_packets, 1, 1)
            this.setValue(datal, 'net-tx-bytes', item.net.tx_bytes, 1, 1)
            this.setValue(datal, 'net-tx-packets', item.net.tx_packets, 1, 1)
            this.setValue(datal, 'net-total-bytes', item.net.total_bytes, 1, 1)
          }
          list.push(
            new GraphHistoricData(
              this.parseTime(item.group),
              datal
            )
          )
        }
      }
      return list
    });
  }

  setValue(datal :{ [name:string]: number; }, name : string, val : number, mul : number, div : number) {
    if (val) {
      datal[name] = (val * mul) / div
    } else {
      datal[name] = 0
    }
  }

//--------------------------------------------------------------------------------------
// http core functions
//--------------------------------------------------------------------------------------

  private setHeaders() {
    var headers = new Headers
    headers.set('Authorization', 'amp '+this.token)
    return headers
  }

  setToken(token : string) {
    this.token = token
  }

  httpGet(url : string) : Observable<any> {
    let headers = this.setHeaders()
    return this.http.get(this.addr+url, { headers: this.setHeaders() })
      .retryWhen(e => e.scan<number>((errorCount, err) => {
        console.log("retry: "+errorCount)
        if (errorCount >= httpRetryNumber) {
            throw err;
        }
        return errorCount + 1;
      }, 0).delay(httpRetryDelay)
    )
  }

  httpDelete(url : string) : Observable<any> {
    let headers = this.setHeaders()
    return this.http.delete(this.addr+url, { headers: this.setHeaders() })
      .retryWhen(e => e.scan<number>((errorCount, err) => {
        console.log("retry: "+errorCount)
        if (errorCount >= httpRetryNumber) {
            throw err;
        }
        return errorCount + 1;
      }, 0).delay(httpRetryDelay)
    )
  }

  httpPost(url : string, data : any) : Observable<any> {
    let headers = this.setHeaders()
    return this.http.post(this.addr+url, data, { headers: this.setHeaders() })
      .retryWhen(e => e.scan<number>((errorCount, err) => {
        console.log("retry: "+errorCount)
        if (errorCount >= httpRetryNumber) {
            throw err;
        }
        return errorCount + 1;
      }, 0).delay(httpRetryDelay)
    )
  }

}