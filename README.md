![build](https://github.com/targetaidev/sideweed/workflows/Build/badge.svg) ![license](https://img.shields.io/badge/license-AGPL%20V3-blue)

*sideweed* is a high-performance sidecar load-balancer. By attaching a tiny load balancer as a sidecar to each of the client application processes, you can eliminate the centralized loadbalancer bottleneck and DNS failover management. *sideweed* automatically avoids sending traffic to the failed servers by checking their health via the readiness API and HTTP error returns.

# Install

## Binary Releases

Download the latest binary from https://github.com/targetaidev/sideweed/releases and unzip a single binary file.

## Build from source

```
go install -v github.com/targetaidev/sideweed@latest
```

> You will need a working Go environment. Therefore, please follow [How to install Go](https://golang.org/doc/install).
> Minimum version required is go1.17

# Usage

```
NAME:
  sideweed - High-Performance sidecar load-balancer

USAGE:
  sideweed - [FLAGS] SITE1 [SITE2..]

FLAGS:
  --address value, -a value           listening address for sideweed (default: ":8080")
  --health-path value, -p value       health check path
  --read-health-path value, -r value  health check path for read access - valid only for failover site
  --health-port value                 health check port (default: 0)
  --health-duration value, -d value   health check duration in seconds (default: 5s)
  --log, -l                           enable logging
  --trace value, -t value             enable request tracing - valid values are [all,application,cluster] (default: "all")
  --quiet, -q                         disable console messages
  --json                              output sideweed logs and trace in json format
  --debug                             output verbose trace
  --help, -h                          show help
  --version, -v                       print the version

SITE:
  Each SITE is a comma separated list of pools of that site: http://172.17.0.{2...5},http://172.17.0.{6...9}.
  If all servers in SITE1 are down, then the traffic is routed to the next site - SITE2.

EXAMPLES:
  1. Load balance across 4 servers (http://server1:9000 to http://server4:9000)
     $ sideweed --health-path "/health" http://server{1...4}:9000

  2. Load balance across 4 servers (http://server1:9000 to http://server4:9000), listen on port 8000
     $ sideweed --health-path "/health" --address ":8000" http://server{1...4}:9000

  3. Two sites, each site having two pools, each pool having 4 servers:
     $ sideweed --health-path=/health http://site1-server{1...4}:9000,http://site1-server{5...8}:9000 \
               http://site2-server{1...4}:9000,http://site2-server{5...8}:9000

  4. Two sites, each site having two pools, each pool having 4 servers. After failover, allow read requests to site2 even if it has just read quorum:
     $ sideweed --health-path=/health --read-health-path=/health/read  http://site1-server{1...4}:9000,http://site1-server{5...8}:9000 \
               http://site2-server{1...4}:9000,http://site2-server{5...8}:9000
```

## Examples

### Load balance across a web service using DNS provided IPs
```
$ sideweed --health-path=/health http://myapp.myorg.domain
```

### Load balance across servers (http://server1:9000 to http://server4:9000)
```
$ sideweed --health-path=/health --address :8000 http://server{1...4}:9000
```

### Two sites with 4 servers each
```
$ sideweed --health-path=/health http://site1-server{1...4}:9000 http://site2-server{1...4}:9000
```
