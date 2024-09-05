#!/usr/bin/env bash

count=0

if [[ $# < 2 ]]; then
  echo "Usage: hamming.sh <string1> <string2>"
  exit -1
elif [[ ${#1} != ${#2} ]]; then
  echo "strands must be of equal length" 
  exit -1
fi

for (( i=0; i<${#1}; i++ )); do
  [ ${1:i:1} == ${2:i:1} ] || (( count++ ))
done
echo $count

