#!/bin/bash
function abort(){
    echo "Your change doesn't follow PaddleCloud's code style." 1>&2
    echo "Please use pre-commit to reformat your code and git push again." 1>&2
    exit 1
}

trap 'abort' 0
set -e

apt-get remove clang-format
apt-get install -y clang-format-3.8

cd $TRAVIS_BUILD_DIR
export PATH=/usr/bin:$PATH
pre-commit install
pre-commit  --version

if ! pre-commit run -a ; then
  git diff  --exit-code
  exit 1
fi

trap : 0
