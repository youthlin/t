#!/usr/bin/env bash
rm -rf output
go env
go mod tidy

RED='\033[1;31m'   #红
GREEN='\033[1;32m' #绿
RES='\033[0m'
OUT='xtemplate'

if [ "${ENV}" == "dev" ]; then
    echo -e "Compiling in ${GREEN}develop${RES} mode."
    go build -gcflags="all=-N -l" -o output/${OUT}
else
    go build -o output/${OUT}
fi

# 打印编译结果
RET=$?
if [ $RET == 0 ] ; then
    echo -e "build ${GREEN}success${RES}"
else
    echo -e "build ${RED}failed${RES}"
    exit $RET
fi
