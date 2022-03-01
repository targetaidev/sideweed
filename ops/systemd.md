# Systemd service for sideweed

Systemd script for sideweed load balancer.

## Installation

- Systemd script is configured to run the binary from /usr/local/bin/
- Download the binary from https://github.com/targetaidev/sideweed/releases

## Create default configuration

```sh
$ cat <<EOF >> /etc/default/sideweed
# sideweed options
SIDEWEED_OPTIONS="--health-path=/health --address :8000"

# sideweed sites
SIDEWEED_SITES="http://172.17.0.{11...18}:9000 http://172.18.0.{11...18}:9000"

EOF
```

## systemctl

Copy sideweed.service to /etc/systemd/system directory.

## Enable startup on boot

```sh
systemctl enable sideweed.service
```

## Note
Replace User=nobody and Group=nobody file with your local setup user.
