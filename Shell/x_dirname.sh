#!/bin/bash

# 输入：../Shell/x_dirname.sh
# 输出：../Shell
DIR=`dirname $0`
echo $DIR

# 输出：a
DIR1=`dirname a/b`
echo $DIR1
