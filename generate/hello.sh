#!/usr/bin/env bash

protoDir="../proto/hello"
outDir="../proto/hello"
protoc -I ${protoDir}/ ${protoDir}/*proto --micro_out=${outDir} --go_out=plugins=grpc:${outDir}