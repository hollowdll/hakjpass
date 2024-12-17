#!/bin/bash

# This generates the Go Protobuf source codes from the proto files
protoc --go_out=. --go_opt=paths=source_relative pb/*.proto
