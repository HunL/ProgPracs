class Singleton(object):
	def __new__(cls, *arg, **kw):
		if not hasattr(cls, '_instance'):
			orig = super(Singleton,cls)
			cls._instance = orig.__new__(cls,*arg,**kw)
		return cls._instance 
		
class MyClass(Singleton):
	a = 1
	
one = MyClass()
two = MyClass()

two.a = 23
print(one.a)



