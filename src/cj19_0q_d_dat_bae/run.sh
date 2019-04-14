#!/bin/bash
echo running test set ${1-0}
go build . && rm f0 && mkfifo f0 ; python testing_tool.py ${1-0} < f0 | ./cj19_0q_d > f0
