#!/bin/sh
set -ex

GOBIN=gotip

"${GOBIN}" generate ./...
"${GOBIN}" run main.go