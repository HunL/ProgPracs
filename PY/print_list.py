print '>>>>> print directly >>>>>'
print range(0,13)

print '>>>>> list >>>>>'
print [x * x for x in range(1,8)]

print '>>>>> condition filter >>>>>'
print [x * x for x in range(1,11) if x % 2 != 0]

print '>>>>> multiple nested >>>>>'
print [m + n for m in 'abc' for n in '123']

print '>>>>> enumerate >>>>>'
L = ['Adam', 'Lisa', 'Bart', 'Paul']
for index,value in enumerate(L):
    print index,'-',value



print '>>>>> iteration of dict >>>>>'
d = {'Adam':95, 'Lisa':80, 'Bart':60}

print '>>>>> directly >>>>>'
print d.values()

print '>>>>> normal iteration >>>>>'
for v in d.values():
    print v

print '>>>>> iteration for save memory >>>>>'
for v in d.itervalues():
    print v 

print '>>>>> k-v >>>>>'
print d.items()

print '>>>>> k-v iteration >>>>>'
for k,v in d.iteritems():
    print k,':',v
