#!/bin/bash

java \
  -cp build/libs/codejam-1.0.jar \
  TaskGen ctm 40000 40000 > inout/large.txt

time java -Xmx512m \
  -cp build/libs/codejam-1.0.jar \
  SolutionBrute < inout/large.txt > inout/large.brute.txt

while true; do

  time go/Solution < inout/large.txt > inout/large.go.txt

  if diff inout/large.go.txt inout/large.brute.txt ; then
    echo no differences, trying again
  else
    exit 1;
  fi

done
