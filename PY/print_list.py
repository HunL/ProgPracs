print '>>>>>print directly>>>>>'
print range(0,13)

print '>>>>>print list>>>>>'
print [x * x for x in range(1,8)]

print '>>>>>print condition filter>>>>>'
print [x * x for x in range(1,11) if x % 2 != 0]

print '>>>>>print multiple nested'
print [m + n for m in 'abc' for n in '123']
