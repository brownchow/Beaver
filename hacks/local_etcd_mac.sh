#!/bin/bash

curl -OL https://github.com/etcd-io/etcd/releases/download/v3.4.14/etcd-v3.4.14-darwin-amd64.zip
unzip etcd-v3.4.14-darwin-amd64.zip

./etcd-v3.4.14-darwin-amd64/etcd

# curl -L http://127.0.0.1:2379/version
