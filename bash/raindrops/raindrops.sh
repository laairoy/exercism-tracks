#!/usr/bin/env bash

result=""

if [ $# -le 0 ]; then
  exit 1
fi

# if [ `expr $1 % 3` == 0 ]; then
#  result+="Pling"
#fi
#if [ `expr $1 % 5` == 0 ]; then
#  result+="Plang"
#fi
#if [ `expr $1 % 7` == 0 ]; then
#  result+="Plong"
#fi

(( $1 % 3 )) || result+=Pling
(( $1 % 5 )) || result+=Plang
(( $1 % 7 )) || result+=Plong

echo ${result:-$1}
