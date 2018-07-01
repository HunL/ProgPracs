#encoding=utf-8
#from urllib.request import urlopen
from urllib.request import urlretrieve
#import re
import urllib.request

def getUrl(str):
	start = str.find('href=')
	start += len('href="')
	end = str.find('.html')
	end += len('.html')
	url = str[start:end]
	#print(url)
	return url

def storePage(url):
	#page = url.urlopen(url).read()
	local = 'G://work//fuyun//dl//page//ok.html'
	urlretrieve(url, local)
	
	
str = '<a title="" target="_blank" href="http://blog.sina.com.cn/s/blog_5456160901019qow.html">'
url = getUrl(str)
storePage(url)