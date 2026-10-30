#!/bin/bash
# Copyright © 2023 KubeCub. All rights reserved.
#
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.


set -euo pipefail

BIN_DIR="${OUTPUT_DIR:-./bin}"
mkdir -p "$BIN_DIR"

go build -mod=readonly -o "$BIN_DIR/example" ./example/example.go
go build -mod=readonly -o "$BIN_DIR/simple" ./example/simple/simple.go
go build -mod=readonly -o "$BIN_DIR/context" ./example/context/main.go
go build -mod=readonly -o "$BIN_DIR/vlevel" ./example/vlevel/v_level.go
