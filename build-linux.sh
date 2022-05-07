#!/bin/bash

go build -ldflags="-s -w" -o tmp-code . && upx tmp-code -9