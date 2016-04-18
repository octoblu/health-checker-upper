#!/bin/bash

APP_NAME=health-checker-upper
TMP_DIR=$PWD/tmp
IMAGE_NAME=local/$APP_NAME

build_on_docker() {
  docker build --tag $IMAGE_NAME:built .
}

build_on_local() {
  env GOOS=linux go build -a -tags netgo -installsuffix cgo -ldflags '-w' -o "${TMP_DIR}/${APP_NAME}" .
}

build_osx_on_local() {
  env GOOS=darwin go build -a -tags netgo -installsuffix cgo -ldflags '-w' -o "${APP_NAME}" .
}

copy() {
  cp $TMP_DIR/$APP_NAME .
  cp $TMP_DIR/$APP_NAME entrypoint/
}

init() {
  rm -rf $TMP_DIR/ \
   && mkdir -p $TMP_DIR/
}

package() {
  docker build --tag $IMAGE_NAME:latest entrypoint
}

run() {
  docker run --rm \
    --volume $TMP_DIR:/export/ \
    $IMAGE_NAME:built \
      cp $APP_NAME /export
}

panic() {
  local message=$1
  echo $message
  exit 1
}

docker_build() {
  init    || panic "init failed"
  build_on_docker || panic "build_on_docker failed"
  run     || panic "run failed"
  copy    || panic "copy failed"
  package || panic "package failed"
}

local_build() {
  init    || panic "init failed"
  build_on_local || panic "build_on_local failed"
  copy    || panic "copy failed"
  package || panic "package failed"
}

osx_build() {
  init    || panic "init failed"
  build_osx_on_local || panic "build_osx_on_local failed"
}

release_build() {
  mkdir -p dist \
  && osx_build \
  && tar -czf "${APP_NAME}-osx.tar.gz" "${APP_NAME}" \
  && mv "${APP_NAME}-osx.tar.gz" dist/
  echo "Wrote dist/${APP_NAME}-osx.tar.gz"
}

main() {
  local mode="$1"
  if [ "$mode" == "local" ]; then
    echo "Local Build"
    local_build
    exit $?
  fi

  if [ "$mode" == "docker" ]; then
    echo "Docker Build"
    docker_build
    exit $?
  fi

  if [ "$mode" == "osx" ]; then
    echo "OSX Build"
    osx_build
    exit $?
  fi

  if [ "$mode" == "release" ]; then
    echo "Release Build"
    release_build
    exit $?
  fi

  echo "Usage: ./build.sh local/docker/osx/release"
  exit 1
}
main $@
