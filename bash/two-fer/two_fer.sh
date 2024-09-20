#!/usr/bin/env bash

main () {
#  if [ ${#1} -gt 0 ]; then
#    name=$1
#  else 
#    name=you
#  fi
  
#  echo "One for $name, one for me."

echo "One for ${1-you}, one for me." # Default Value Parameter Expansion

}

main "$@"
