#!/usr/bin/env bash

SOURCE_DIR=$(cd "$(dirname $(dirname "$0"))"; pwd)

DIST_DIR="${SOURCE_DIR}/dist"

mkdir -p ${DIST_DIR}

CONFIG_NAME="database.json"

CONFIG_PATH="${SOURCE_DIR}/${CONFIG_NAME}.example"
CONFIG_DIST_PATH="${DIST_DIR}/${CONFIG_NAME}"

STATIC_PATH="${SOURCE_DIR}/web/"
STATIC_DIST_PATH="${DIST_DIR}/web/"

BUILD_NAME="dbbook"
BUILD_PATH="${DIST_DIR}/${BUILD_NAME}"

if [[ ! -f ${CONFIG_DIST_PATH} ]]; then
    cp ${CONFIG_PATH} ${CONFIG_DIST_PATH}
fi

rsync -av --delete ${STATIC_PATH} ${STATIC_DIST_PATH}

cd ${SOURCE_DIR} \
&& go get -v \
&& CGO_ENABLED=0 go build -o ${BUILD_PATH}