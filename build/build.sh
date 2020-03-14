#!/usr/bin/env bash

# Usage:
# ./build/build.sh 1.0

if [[ -z "${1}" ]]; then
    echo "version number is required, example: ${0} 1.0"
    exit 0
fi

VERSION="v${1}"

SOURCE_DIR=$(cd "$(dirname $(dirname "$0"))"; pwd)

DIST_DIR="${SOURCE_DIR}/dist/${VERSION}"

mkdir -p ${DIST_DIR}

CONFIG_NAME="database.json"

CONFIG_PATH="${SOURCE_DIR}/${CONFIG_NAME}.example"

STATIC_PATH="${SOURCE_DIR}/web/"

function build() {
    if [[ "${os}" == "darwin" && "${arch}" == "386" ]] ; then
        continue
    fi

    BUILD_NAME="dbbook"

    if [[ "${os}" == "windows" ]] ; then
        BUILD_NAME="${BUILD_NAME}.exe"
    fi

    BUILD_DIR_NAME="dbbook_${os}_${arch}_${VERSION}"

    echo "📦 ${BUILD_DIR_NAME} building..."

    BUILD_DIR_PATH="${DIST_DIR}/${BUILD_DIR_NAME}"

    mkdir -p ${BUILD_DIR_PATH}

    CONFIG_DIST_PATH="${BUILD_DIR_PATH}/${CONFIG_NAME}"

    cp ${CONFIG_PATH} ${CONFIG_DIST_PATH}

    STATIC_DIST_PATH="${BUILD_DIR_PATH}/web/"

    rsync -aq --delete ${STATIC_PATH} ${STATIC_DIST_PATH}

    BUILD_PATH="${BUILD_DIR_PATH}/${BUILD_NAME}"

    cd ${SOURCE_DIR} && CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -o ${BUILD_PATH}

    cd ${DIST_DIR} && zip -r -o -q "${BUILD_DIR_NAME}.zip" ${BUILD_DIR_NAME}
}

go get -v

for os in windows darwin linux freebsd ; do
    for arch in 386 amd64; do
        build
    done
done

echo "✅  All done."
