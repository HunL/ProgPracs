def sub_sort(array,low,high):
	key = array[low]
	print('key : ',key)
	while low < high:
		while low < high and array[high] >= key:
			print('high:', high,array[high])
			high -= 1
		while low < high and array[high] < key:
			print('cccc high: ',high,array[high])
			array[low] = array[high]
			low += 1
			array[high] = array[low]
	array[low] = key
	return low
	
def qs(array,low,high):
	if low < high:
		k = sub_sort(array,low,high)
		qs(array,low,k)
		qs(array,k+1,high)
		
if __name__ == '__main__':
	array = [2,5,6,7,8,9,1,4,98,99]
	print(array)
	qs(array,0,len(array)-1)
	print(array)