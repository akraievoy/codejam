#!/bin/bash

set -o nounset
set -o errexit

export GO_ROOT=/home/ak/bin/go-1.7.4
export GO_PATH="$( pwd )"
export CJ_TASK="$1"
export CJ_TIME="$(date +%y%m%d_%H%M%S)"

echo "
--=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
go build: ROOT=${GO_ROOT} PATH=${GO_PATH}
--=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"
bash -c "time ${GO_ROOT}/bin/go build -o bin/${CJ_TASK} src/${CJ_TASK}/Solution.go"

for TEST_OUT in `ls -1 inout/${CJ_TASK}/ | grep -E '[0-9]+_[0-9]+.out$'` ; do
    echo "deleting $TEST_OUT" && rm inout/${CJ_TASK}/$TEST_OUT
done;

for TEST_IN in `ls -1 inout/${CJ_TASK}/ | grep '.in$' | sort -n`; do
    echo "
    --=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
    Comparing output for ${TEST_IN}
    --=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"

    TEST_BASE="$(echo $TEST_IN | sed -e 's/\.in$//g')"

    time bin/${CJ_TASK} \
        inout/${CJ_TASK}/$TEST_IN > \
        inout/${CJ_TASK}/${TEST_BASE}.${CJ_TIME}.actual.out

    diff -y --suppress-common-lines \
        inout/${CJ_TASK}/${TEST_BASE}.out \
        inout/${CJ_TASK}/${TEST_BASE}.${CJ_TIME}.actual.out
done

echo "
--=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
READY FOR UPLOAD
--=[${CJ_TASK}]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"
