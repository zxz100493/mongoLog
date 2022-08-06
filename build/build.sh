#!/bin/bash
set -u
set -e
set -x
set -v

docker build -t apps:v1 ../ > /tmp/build.log
