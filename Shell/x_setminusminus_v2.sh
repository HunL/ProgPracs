#!/bin/bash

set -- -z a b c
echo $1
echo $2
echo $3

echo $@

# It's "good practice" to get into the habit of using -- when you're not sure what is following :-)
set -z a b c
echo $1
echo $2
echo $3
echo $4

echo $@
