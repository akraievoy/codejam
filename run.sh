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
go build -o bin/${CJ_TASK} codejam.go

echo
echo
echo '----------============= SAMPLE IN/OUT FILES =============----------'
for TEST_IN in `ls -1 inout/ | grep '.in$'`; do
  TEST_BASE=`echo $TEST_IN | sed -e 's/\.in$//g'`
  echo "running $TEST_IN -> $TEST_BASE.${CJ_TIME}.out..."
  bin/${CJ_TASK} inout/$TEST_IN > inout/${TEST_BASE}.${CJ_TIME}.out
  diff inout/${TEST_BASE}.${CJ_TIME}.out inout/${TEST_BASE}.out
done

mkdir -p $CJ_HOME/$CJ_TASK/$CJ_TIME

echo
echo
echo '----------============= ZIPPING SOLUTION =============----------'
zip -r \
  $CJ_HOME/$CJ_TASK/$CJ_TIME/$CJ_TASK-$CJ_TIME.zip \
  *.go

unzip -l $CJ_HOME/$CJ_TASK/$CJ_TIME/$CJ_TASK-$CJ_TIME.zip

echo
echo
echo '----------============= LISTENING FOR DOWNLOADS =============----------'
echo watching for new downloads at $CJ_HOME/$CJ_TASK/$CJ_TIME && \
  fileschanged --show=created --recursive --timeout=2 $CJ_HOME/$CJ_TASK/$CJ_TIME | \
  xargs -L1 -iIN bash -c "if [[ IN = *.in ]] ; then echo IN && bin/${CJ_TASK} IN > IN.out ; fi"
