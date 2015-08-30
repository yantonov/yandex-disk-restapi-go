#!/bin/sh

SCRIPT_DIR=$(cd `dirname $0` && pwd)

cd $SCRIPT_DIR/..

go test github.com/yantonov/yandex-disk-restapi-go/test
