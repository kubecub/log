#!/bin/bash
# Copyright © 2023 KubeCub. All rights reserved.
#
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.


BIN_DIR="/bin"

if [[ ! -d "$BIN_DIR" ]]; then
  mkdir "$BIN_DIR"
fi

DIR_1="./example/example.go"
DIR_2="./example/simple/simple.go"
DIR_3="./example/context/main.go"
DIR_4="./example/vlevel/v_level.go"

printf "===========> Building binaries for %s"
go build -o "$BIN_DIR/$DIR_1" $DIR_1
go build -o "$BIN_DIR/$DIR_2" $DIR_2
go build -o "$BIN_DIR/$DIR_3" $DIR_3
go build -o "$BIN_DIR/$DIR_4" $DIR_4