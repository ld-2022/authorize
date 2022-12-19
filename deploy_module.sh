#!/bin/zsh
git config --global https.proxy http://127.0.0.1:7890

git config --global https.proxy https://127.0.0.1:7890

version=v0.0.1-beta
echo "发布版本: $version"
git tag $version
git push origin $version

git config --global --unset http.proxy

git config --global --unset https.proxy