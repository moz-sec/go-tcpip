#!/usr/bin/env bash

set -euxo pipefail

if [ "$(whoami)" != "root" ]; then
  echo "Require root privilege"
  exit 1
fi

ip tuntap add mode tap user "$USER" name tap0
ip addr add 192.0.2.1/24 dev tap0
ip link set tap0 up

ip addr show
