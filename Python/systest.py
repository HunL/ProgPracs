#coding=utf-8
import sys

print("script name:", sys.argv[0])
for i in range(1, len(sys.argv)):
	print("arg:", i, sys.argv[i])