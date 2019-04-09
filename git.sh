#!/usr/bin/env bash
git add .

echo -n "git commit message:..." ---:
read message
git commit -m "$message"
git push origin master
#exit