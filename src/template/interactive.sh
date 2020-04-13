#!/bin/bash
TASK_NAME=d
for test_number in 0 1 2 ; do
  python3 src/${TASK_NAME}/interactive_runner.py python3 src/${TASK_NAME}/local_testing_tool.py ${test_number} -- bin/${TASK_NAME}
done
