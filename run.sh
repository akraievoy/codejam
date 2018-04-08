#!/bin/bash

set -o nounset
set -o errexit

if [ "${1-}" = "" ]; then
    echo "Usage:
    run.sh TASK_NAME

Valid values for TASK_NAME are: "
    ls -1 src | sed -E 's/^/ * /'
    exit 1;
fi

if [ "${GOROOT-}" = "" ]; then
    export GOROOT=/home/ak/bin/go-1.7.4
fi

export GO_PATH="$( pwd )"
export CJ_TASK="$1"
export CJ_TIME="$(date +%y%m%d_%H%M%S)"

echo -e "
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
 \033[1;36m*\033[0m go test: GOROOT=${GOROOT} GO_PATH=${GO_PATH}
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"
bash -c "cd src/${CJ_TASK}/ && time ${GOROOT}/bin/go test"

echo -e "
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
 \033[1;36m*\033[0m go build: GOROOT=${GOROOT} GO_PATH=${GO_PATH}
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"
bash -c "time ${GOROOT}/bin/go build -o bin/${CJ_TASK} src/${CJ_TASK}/Solution.go"

echo ""
for TEST_OUT in `ls -1 inout/${CJ_TASK}/ | grep -E '[0-9]+_[0-9]+.actual.out$'` ; do
    echo "deleting $TEST_OUT" && rm inout/${CJ_TASK}/$TEST_OUT
done;

for TEST_IN in `ls -1 inout/${CJ_TASK}/ | grep '.in$' | sort -n`; do
    echo -e "
    --=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
         \033[36m*\033[0m Comparing output for ${TEST_IN}
    --=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"

    TEST_BASE="$(echo $TEST_IN | sed -e 's/\.in$//g')"

    time bin/${CJ_TASK} \
        inout/${CJ_TASK}/$TEST_IN > \
        inout/${CJ_TASK}/${TEST_BASE}.${CJ_TIME}.actual.out

    diff -y --suppress-common-lines \
        inout/${CJ_TASK}/${TEST_BASE}.out \
        inout/${CJ_TASK}/${TEST_BASE}.${CJ_TIME}.actual.out
done

if which fileschanged ; then

    echo -e "
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
 \033[1;36m*\033[0m creating solution archive -- required for pre-2018 versions of CodeJam
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"

    zip -r \
      $CJ_TASK-$CJ_TIME.zip \
      src/${CJ_TASK}

    unzip -l ${CJ_TASK}-$CJ_TIME.zip

    echo -e "
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
    \033[1;32mecho watching for new inputs at inout/${CJ_TASK} -- for pre-2018 CodeJam\033[0m
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"

    fileschanged --show=created --recursive --timeout=2 inout/${CJ_TASK} | \
        xargs -L1 -iIN bash -c "if [[ IN = *.in ]] ; then echo IN && bin/${CJ_TASK} IN > IN.out ; fi"

else

    echo -e "
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--
            \033[1;32m---================<[ READY FOR UPLOAD ]>================---\033[0m
--=[\033[1;37m${CJ_TASK}\033[0m]=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--"

fi
