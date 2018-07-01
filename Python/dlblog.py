from urllib.request import urlretrieve
from urllib.request import urlopen
import re
import urllib.request

def getContent(url):
	content = urlopen(url).read()
	return content

def getUrls(url):
	content = getContent(url)
	regrex = re.compile(r'title=.+target=.+href="(.+?\.html)"')
	urls = regrex.findall(content.decode('utf-8'))
	print(urls)
	return urls
	
def storePages(url, page):
	urls = getUrls(url)
	cnt = 1
	for url in urls:
		local = 'G://work//fuyun//dl//blog//'+'page'+str(page)+'_No.'+str(cnt)+'.html'
		urlretrieve(url, local)
		cnt += 1
	
def storeBlog():
	for i in range(1, 10):
		url = 'http://blog.sina.com.cn/s/articlelist_1414927881_0_'+str(i)+'.html'
		storePages(url, i)
		
storeBlog()