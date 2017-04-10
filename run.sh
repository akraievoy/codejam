#!/bin/bash

set -o nounset
set -o errexit

export CJ_HOME=${1-$HOME/Downloads/codejam}

export CJ_TIME=`date +%y%m%d_%H%M%S`

export CJ_TASK=`git rev-parse --abbrev-ref HEAD | sed -Ee 's:/(go|java)$::g' | sed -Ee 's:[^/]+/::g'`

echo
echo
echo '----------============= GO TEST =============----------'
go test

echo
echo
echo '----------============= GO BUILD =============----------'
go build -o bin/${CJ_TASK} solution.go

echo
echo
echo '----------============= SAMPLE IN/OUT FILES =============----------'
for TEST_IN in `ls -1 inout/ | grep '.in$'`; do
  TEST_BASE=`echo $TEST_IN | sed -e 's/\.in$//g'`
  echo "running $TEST_IN -> $TEST_BASE.${CJ_TIME}.out..."
  bin/${CJ_TASK} inout/$TEST_IN > inout/${TEST_BASE}.${CJ_TIME}.out
  diff inout/${TEST_BASE}.${CJ_TIME}.out inout/${TEST_BASE}.out
done

echo
echo
echo
echo
echo
echo
echo '----------===========<[ READY FOR UPLOAD ]>===========----------'
echo '----------==========<[[ READY FOR UPLOAD ]]>==========----------'
echo '----------===========<[ READY FOR UPLOAD ]>===========----------'
