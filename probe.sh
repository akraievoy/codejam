#!/bin/bash
while true; do
  java \
    -cp build/libs/codejam-1.0.jar \
    TaskGen ctm 40000 40000 > inout/probe.txt

  time java -Xmx512m \
    -cp build/libs/codejam-1.0.jar \
    SolutionBrute < inout/probe.txt > inout/probe.brute.txt

  time go/Solution < inout/probe.txt > inout/probe.go.txt

  if diff inout/probe.go.txt inout/probe.brute.txt ; then
    echo no differences, trying another seed
  else
    exit 1;
  fi
done
