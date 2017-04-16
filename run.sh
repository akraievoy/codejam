#!/bin/bash

set -o nounset
set -o errexit

export CJ_HOME=${1-$HOME/Downloads/codejam}

export CJ_TIME=`date +%y%m%d_%H%M%S`

export CJ_TASK=`git rev-parse --abbrev-ref HEAD | sed -Ee 's:/(go|java)$::g' | sed -Ee 's:[^/]+/::g'`

echo
echo
echo '----------============= GO TEST =============----------'
bash -c "GOPATH=`pwd`:$GOPATH && cd src && pwd && env | grep GOPATH && go test"

echo
echo
echo '----------============= GO BUILD =============----------'
bash -c "GOPATH=`pwd`:$GOPATH && cd src && pwd && go build -o ../bin/${CJ_TASK}"

echo
echo
echo '----------============= SAMPLE IN/OUT FILES =============----------'
for TEST_OUT in `ls -1 inout/ | grep -E '[0-9]+_[0-9]+.out$'` ; do
  echo "deleting $TEST_OUT" && rm inout/$TEST_OUT
done;
for TEST_IN in `ls -1 inout/ | grep '.in$' | sort`; do
  TEST_BASE=`echo $TEST_IN | sed -e 's/\.in$//g'`
  echo "running $TEST_IN -> $TEST_BASE.${CJ_TIME}.out..."
  time bin/${CJ_TASK} inout/$TEST_IN > inout/${TEST_BASE}.${CJ_TIME}.out
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
