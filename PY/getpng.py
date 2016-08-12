#coding=utf-8
from urllib.request import urlopen
from urllib.request import urlretrieve
import re
#import urllib

def getHtml(url):
	page = urlopen(url)
	html = page.read()
	return html
	
def getImg(html):
	imgre = re.compile(r'img src="(.+?\.png)"')
	#imglist = re.findall(imgre, html.decode('utf-8'))
	imglist = imgre.findall(html.decode('utf-8'))
	#return imglist
	x = 0
	for imgurl in imglist:
		#urllib.urlretrieve(imgurl, '%s.png' % x)
		urlretrieve(imgurl, '%s.png' % x)
		x+=1
	
html = getHtml("http://www.cnblogs.com/skynet/archive/2011/05/03/2035105.html")

print(getImg(html))