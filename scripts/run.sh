#!/bin/bash

# this is to allow to run swag as a command
# export PATH=$(go env GOPATH)/bin:$PATH   

# generate/update swagger docs
swag init

go run main.go