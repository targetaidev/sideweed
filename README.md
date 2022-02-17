![build](https://github.com/targetaidev/sidekick/workflows/CI/badge.svg) ![license](https://img.shields.io/badge/license-AGPL%20V3-blue)

*sidekick* is a high-performance sidecar load-balancer. By attaching a tiny load balancer as a sidecar to each of the client application processes, you can eliminate the centralized loadbalancer bottleneck and DNS failover management. *sidekick* automatically avoids sending traffic to the failed servers by checking their health via the readiness API and HTTP error returns.

# Architecture
![architecture](https://raw.githubusercontent.com/minio/sidekick/master/arch_sidekick.png)

# Install

## Binary Releases

| OS      | ARCH    | Binary                                                                                                 |
|:-------:|:-------:|:------------------------------------------------------------------------------------------------------:|
| Linux   | amd64   | [linux-amd64](https://github.com/minio/sidekick/releases/latest/download/sidekick-linux-amd64)         |
| Linux   | arm64   | [linux-arm64](https://github.com/minio/sidekick/releases/latest/download/sidekick-linux-arm64)         |
| Apple   | amd64   | [darwin-amd64](https://github.com/minio/sidekick/releases/latest/download/sidekick-darwin-amd64)       |

You can also verify the binary with [minisign](https://jedisct1.github.io/minisign/) by downloading the corresponding [`.minisig`](https://github.com/minio/sidekick/releases/latest) signature file. Then run:
```
minisign -Vm sidekick-<OS>-<ARCH> -P RWTx5Zr1tiHQLwG9keckT0c45M3AGeHD6IvimQHpyRywVWGbP1aVSGav
```

## Build from source

```
go install -v github.com/targetaidev/sidekick@latest
```

> You will need a working Go environment. Therefore, please follow [How to install Go](https://golang.org/doc/install).
> Minimum version required is go1.17

# Usage

```
NAME:
  sidekick - High-Performance sidecar load-balancer

USAGE:
  sidekick - [FLAGS] SITE1 [SITE2..]

FLAGS:
  --address value, -a value           listening address for sidekick (default: ":8080")
  --health-path value, -p value       health check path
  --read-health-path value, -r value  health check path for read access - valid only for failover site
  --health-port value                 health check port (default: 0)
  --health-duration value, -d value   health check duration in seconds (default: 5)
  --insecure, -i                      disable TLS certificate verification
  --log, -l                           enable logging
  --trace value, -t value             enable request tracing - valid values are [all,application,minio] (default: "all")
  --quiet, -q                         disable console messages
  --json                              output sidekick logs and trace in json format
  --debug                             output verbose trace
  --cacert value                      CA certificate to verify peer against
  --client-cert value                 client certificate file
  --client-key value                  client private key file
  --cert value                        server certificate file
  --key value                         server private key file
  --help, -h                          show help
  --version, -v                       print the version

SITE:
  Each SITE is a comma separated list of pools of that site: http://172.17.0.{2...5},http://172.17.0.{6...9}.
  If all servers in SITE1 are down, then the traffic is routed to the next site - SITE2.
```

## Examples

### Load balance across a web service using DNS provided IPs
```
$ sidekick --health-path=/ready http://myapp.myorg.dom
```

### Load balance across 4 MinIO Servers (http://minio1:9000 to http://minio4:9000)
```
$ sidekick --health-path=/minio/health/ready --address :8000 http://minio{1...4}:9000
```

### Two sites with 4 servers each
```
$ sidekick --health-path=/minio/health/ready http://site1-minio{1...4}:9000 http://site2-minio{1...4}:9000
```
