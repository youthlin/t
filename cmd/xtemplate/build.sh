#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
cd "${ROOT_DIR}"

rm -rf output
mkdir -p output
cp -r lang output/
go env

RED='\033[1;31m'   #红
GREEN='\033[1;32m' #绿
RES='\033[0m'
OUT='xtemplate'

if [ "${ENV:-}" == "dev" ]; then
    echo -e "Compiling in ${GREEN}develop${RES} mode."
    go build -gcflags="all=-N -l" -o "output/${OUT}"
else
    go build -o "output/${OUT}"
fi

echo -e "build ${GREEN}success${RES}"
