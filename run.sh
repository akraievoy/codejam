#!/bin/bash

set -o nounset
set -o errexit

gradle clean test jar

export CJ_HOME=${1-$HOME/Downloads/codejam}

export CJ_TASK=`git rev-parse --abbrev-ref HEAD`

export CJ_TIME=`date +%y%m%d_%H%M%S`

mkdir -p $CJ_HOME/$CJ_TASK/$CJ_TIME

zip -r \
  -x 'build/*' 'gradlew*' '.gradle/*' 'gradle/*' '*target*' '.idea*' '.git/*' '*.iml' 'inout/*' 'classes/*' '.gitignore' @ \
  $CJ_HOME/$CJ_TASK/$CJ_TIME/$CJ_TASK-$CJ_TIME.zip \
  .

echo watching for new downloads at $CJ_HOME/$CJ_TASK/$CJ_TIME and inout && \
  fileschanged --show=created --recursive --timeout=2 $CJ_HOME/$CJ_TASK/$CJ_TIME inout | \
  java -jar build/libs/codejam-1.0.jar
