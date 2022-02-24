#! /usr/bin/env bash
set -x

make clear

if [ "$1" != "" ]; then
    make gen4repo repo=gop-app-doc
else
    make gen
fi

cd public
git init
git add -A
git commit -m 'deploy'

if [ "$1" != "" ]; then
    git push -f https://gitee.com/rdor/gop-app-doc master:master
else
    git push -f https://github.com/zrcoder/gop-app-doc master:master
fi