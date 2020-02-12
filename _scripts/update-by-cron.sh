#!/bin/sh
#
# Update vim-jp/redirects by cron.
#
# Usage: update-by-cron.sh {WORKDIR}

set -e

# REPO="git@github.com:vim-jp/redirects.git"
REPO="git@github.com:0Delta/redirects.git"

DIR=$1 ; shift
USER_NAME="redirects cron updater"
USER_EMAIL="redirects+cron%$(hostname -s)@vim-jp.org"

# Setup working directory and cd to it.
if [ -d "$DIR" ] ; then
  cd "$DIR"
  git checkout -q gh-pages
  git fetch -q -p
  git merge -q --ff-only @{u}
else
  parent=$(dirname "$DIR")
  if [ ! -d "$parent" ] ; then
    mkdir -p "$parent"
  fi
  git clone -b gh-pages --depth 50 --quiet "$REPO" "$DIR"
  cd "$DIR"
  git config push.default simple
  git config user.email "$USER_EMAIL"
  git config user.name "$USER_NAME"
fi

# Update repository.
cmd=$(which vim_jp-redirects-update)
if [ "x$cmd" = "x" ] ; then
  go run _scripts/vim_jp-redirects-update/main.go
else
  "$cmd"
fi
git add --update

# Commit changes.
if ! git diff --quiet HEAD ; then
  git commit -m "Updated by cron on $(hostname -s) at $(date "+%Y/%m/%d %H:%M %Z")"
  git push --quiet
fi
