#! /usr/bin/env bash
set -x

make clear dir=09-Docs
make genSite dir=09-Docs
cd 09-Docs/public
git init
git add -A
git commit -m 'deploy'
git push -f https://github.com/zrcoder/gop-app-doc master:master