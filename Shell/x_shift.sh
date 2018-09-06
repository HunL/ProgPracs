#!/bin/bash
#测试 shift 命令

until [ $# -eq 0 ]
do
    echo "第一个参数为：$1 参数个数为：$#"
    shift
done
