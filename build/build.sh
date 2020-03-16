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

if [[ "${DIST_DIR}" == "" || "${DIST_DIR}" == "/" ]]; then
    echo "unexpected dist directory"
    exit 0
fi

if [[ -d "${DIST_DIR}" ]]; then
    rm -rf ${DIST_DIR}
fi

mkdir -p ${DIST_DIR}

CONFIG_NAME="database.json"

CONFIG_PATH="${SOURCE_DIR}/${CONFIG_NAME}.example"

STATIC_PATH="${SOURCE_DIR}/web/"

function build() {
    if [[ "${os}" == "darwin" && "${arch}" == "386" ]] ; then
        continue
    fi

    BUILD_NAME="dbbook"

    ARCHIVE_SUFFIX=".tar.gz"

    if [[ "${os}" == "windows" ]] ; then
        BUILD_NAME="${BUILD_NAME}.exe"
        ARCHIVE_SUFFIX=".zip"
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

    ARCHIVE_NAME="${BUILD_DIR_NAME}${ARCHIVE_SUFFIX}"

    cd ${DIST_DIR} && archive ${ARCHIVE_NAME} ${BUILD_DIR_NAME} && sha256 ${ARCHIVE_NAME}
}

function sha256() {
    echo "`openssl dgst -sha256 -r $1`" >> "sha256_checksum.txt"
}

function archive() {
    if [[ "${os}" == "windows" ]] ; then
        zip_archive $1 $2
    else
        tar.gz_archive $1 $2
    fi
}

function tar.gz_archive() {
    tar -zcf $1 $2 && rm -rf $2
}

function zip_archive() {
    zip -r -o -q -m $1 $2
}

go get -v

for os in darwin freebsd linux windows ; do
    for arch in 386 amd64; do
        build
    done
done

echo "✅  All done."
