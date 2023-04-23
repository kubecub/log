#!/bin/bash
# Copyright © 2023 KubeCub. All rights reserved.
#
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.


BIN_DIR="/bin"

if [[ ! -d "$BIN_DIR" ]]; then
  mkdir "$BIN_DIR"
fi

go build -o "$BIN_DIR/binary1" ./cmd/binary1
go build -o "$BIN_DIR/binary2" ./cmd/binary2
go build -o "$BIN_DIR/binary3" ./cmd/binary3