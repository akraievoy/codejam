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
  diff -y inout/${TEST_BASE}.out inout/${TEST_BASE}.${CJ_TIME}.out
done

mkdir -p $CJ_HOME/$CJ_TASK/$CJ_TIME

echo
echo
echo '----------============= ZIPPING SOLUTION =============----------'
zip -r \
  $CJ_HOME/$CJ_TASK/$CJ_TIME/$CJ_TASK-$CJ_TIME.zip \
  src

unzip -l $CJ_HOME/$CJ_TASK/$CJ_TIME/$CJ_TASK-$CJ_TIME.zip

echo
echo
echo '----------============= LISTENING FOR DOWNLOADS =============----------'
echo '----------============= LISTENING FOR DOWNLOADS =============----------'
echo '----------============= LISTENING FOR DOWNLOADS =============----------'
echo watching for new downloads at $CJ_HOME/$CJ_TASK/$CJ_TIME && \
  fileschanged --show=created --recursive --timeout=2 $CJ_HOME/$CJ_TASK/$CJ_TIME | \
  xargs -L1 -iIN bash -c "if [[ IN = *.in ]] ; then echo IN && bin/${CJ_TASK} IN > IN.out ; fi"
