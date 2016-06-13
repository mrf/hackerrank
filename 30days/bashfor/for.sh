#!/bin/bash
for i in {1..99}
do
  (( comparison = i%2))
  if [ $comparison -eq 0 ]; then
    echo $i
  fi
done
