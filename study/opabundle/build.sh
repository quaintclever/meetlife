#!/bin/zsh
set -e
workspace=$(cd "$(dirname "$0")" && pwd -P)

action="$1"

{
  case "$action" in
  "start")
    cd "$workspace/opasrv" || exit
    make run
    cd "$workspace/opa" || exit
    make run
    ;;
  *)
    cat <<EOF

Usage:
  opa start: start
EOF
    ;;
  esac
}

