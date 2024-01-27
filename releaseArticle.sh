#!/bin/bash

set -ex
pushd docs
git pull origin master
popd

hugo

pushd docs
git add *
git commit -m"New Article" -a
git push origin HEAD:master
popd
