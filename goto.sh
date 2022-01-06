#!/bin/zsh

set -e
workspace=$(cd "$(dirname "$0")" && pwd -P)

action="$1"

{
  case "$action" in
  [1]|"go")
    cd "$workspace/study/gosdk"
    ;;
  [2]|"k8s")
    cd "$workspace/study/k8sopt"
    ;;
  "21"|"cr-replicas")
    cd "$workspace/study/k8sopt/cr-replicas"
    ;;
  [3]|"opa")
    cd "$workspace/study/k8sopt"
    ;;
  *)
    cat <<EOF
======= 感觉这里有个bug, 虽然可以跳转, 但脚本结束, 进程也会结束. 还待查看 =======
[Usage]:
  1. source goto.sh go        =>    cd meetlife/study/gosdk
  2. source goto.sh k8s       =>    cd meetlife/study/k8sopt
    21.source goto.sh k8s     =>    cd meetlife/study/k8sopt/cr-replicas
  3. source goto.sh opa       =>    cd meetlife/study/opabundle

======= 别名配置 =======
[Alias]:
alias goto=". ./goto.sh"

you can input:
goto 1
goto 2
goto 21
goto 3
=============================================================================
EOF
    ;;
  esac
}