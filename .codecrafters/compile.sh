#!/bin/sh
#
# 
# This runs before .codecrafters/run.sh
#

# Exit early if any commands fail
set -e

go build -o /tmp/shell-target cmd/myshell/*.go
