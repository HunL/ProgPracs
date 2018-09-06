# *_* coding:utf-8 *_*

import time
import threading

def a():
    print ('now a start running:'), time.ctime()
    time.sleep(5)
    print ('a is ending:'), time.ctime()

def b():
    print ('now b start running:'), time.ctime()
    time.sleep(10)
    print ('b is ending:'), time.ctime()

def test():
    a1 = threading.Thread(target=a, args=())
    b1 = threading.Thread(target=b, args=())
    p = [a1, b1]
    for i in range(len(p)):
        p[i].start()
    for i in range(len(p)):
        p[i].join()

test()
