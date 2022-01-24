#!/bin/sh
set -ex

GO=gotip

"${GO}" generate ./...
"${GO}" install
calm-defusor