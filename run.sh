#!/bin/bash

set -o nounset
set -o errexit

function pretty_echo {
    local CAPTION=$1
    local MESSAGE=$2
    local COLOR=${3-blue}
    if [ "$COLOR" == "green" ]; then
        echo -e "\033[1;37m${CAPTION}\033[0m :: \033[1;32m${MESSAGE}\033[0m"
    else
        echo -e "\033[1;37m${CAPTION}\033[0m :: \033[1;36m${MESSAGE}\033[0m"
    fi
}

clear
pretty_echo "INIT" "DELETING TEST OUTPUTS" "green"
for TEST_OUT in `find -iname '*.actual.out'` ; do
    rm -v ${TEST_OUT}
done;
echo

if [ "${1-}" = "" ]; then
    echo "Usage:
    run.sh TASK_RELPATH

Valid values for TASK_RELPATH are: "
    find src -type d | tail -n +2 | sed -E 's/^/ * /'
    exit 1;
fi

if [ "${GOROOT-}" = "" ]; then
    export GOROOT=/home/ak/bin/go-1.7.4
fi

export GOPATH="$( pwd )"
TASK_RELPATH=$1
TASK_BINARY="$( echo $1 | sed -e 's:^src/::g' | sed -Ee 's:/$::g' )"

export CJ_TIME="$(date +%y%m%d_%H%M%S)"

pretty_echo "${TASK_RELPATH}" "go test with GOROOT=${GOROOT} GOPATH=${GOPATH}"
bash -c "GOROOT=${GOROOT} && GOPATH=${GOPATH} && cd ${TASK_RELPATH} && time ${GOROOT}/bin/go test"

pretty_echo "${TASK_RELPATH}" "go build with GOROOT=${GOROOT} GOPATH=${GOPATH}"
bash -c "GOROOT=${GOROOT} && GOPATH=${GOPATH} && time ${GOROOT}/bin/go build -o bin/${TASK_BINARY} ${TASK_RELPATH}/Solution.go"

for TEST_IN in `find ${TASK_RELPATH} -iname '*.in' | sort -n`; do
    pretty_echo "${TASK_RELPATH}" "Comparing output for ${TEST_IN}"

    TEST_BASE="$(echo $TEST_IN | sed -e 's/\.in$//g')"

    time bin/${TASK_BINARY} \
        $TEST_IN > \
        ${TEST_BASE}.${CJ_TIME}.actual.out

    diff -y --suppress-common-lines \
        ${TEST_BASE}.out \
        ${TEST_BASE}.${CJ_TIME}.actual.out
done

pretty_echo "${TASK_RELPATH}" "READY FOR UPLOAD" "green"