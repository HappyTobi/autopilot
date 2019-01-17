#!/bin/bash
# vim: set ft=sh

set -e

cd $GOPATH/src/github.com/happytobi/autopilot

govendor install github.com/happytobi/autopilot/vendor/github.com/onsi/ginkgo/ginkgo

ginkgo -r "$@"
