#!/bin/sh
#
# Update vim-jp/redirects by cron.
#
# Usage: cron-update.sh {WORKDIR}

set -e

REPO="git@github.com:vim-jp/redirects.git"

DIR=$1 ; shift
USER_NAME="redirects cron updater"
USER_EMAIL="redirects+cron%$(hostname -s)@vim-jp.org"

if [ ! -d "$DIR" ] ; then
  parent=$(dirname "$DIR")
  echo $parent
  if [ ! -d "$parent" ] ; then
    mkdir -p "$parent"
  fi
  git clone -b gh-pages --depth 50 "$REPO" "$DIR"
  git config push.default simple
  git config user.email "$USER_EMAIL"
  git config user.name "$USER_NAME"
fi

cd "$DIR"
go run _scripts/vim_jp-redirects-update/main.go
git add --update
if ! git diff --quiet HEAD ; then
  git commit -m "Updated by cron on $(hostname -s) at $(date "+%Y/%m/%d %H:%M %Z")"
  #git push --quiet
fi
