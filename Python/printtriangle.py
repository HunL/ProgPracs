#coding=utf-8
import sys

len = int(sys.argv[1])

def printtriangle():
	x = 1
	for i in range(1, len+1):
		for j in range(1, i+1):
			print(2*x-1, end=' ')
			x += 1
		print(end='\n')

printtriangle()