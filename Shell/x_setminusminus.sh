#!/bin/bash

set a b c
echo $1
echo $2
echo $3

echo $@


#set -- testarg "$@"
set -- testarg $@
echo $1
echo $2
echo $3
echo $4

echo $@
