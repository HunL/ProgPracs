def sub_sort(array,low,high):
	key = array[low]
	pivot = low
	while low < high:
		while low < high and array[high] >= key:
			print('555', high,low,array[high],array)
			high -= 1
		while low < high and array[low] <= key:
			print('888', high,low,array[low],array)
			low += 1
		while low < high and array[high] < key and array[low] > key:
			print('111', high,low,array[high])
			tmp = array[high]
			array[high] = array[low]
			array[low] = tmp
			if high > low+1:
				high -= 1
				low += 1
			else:
				print('2000',high,low,array)
				break
	tmp2 = array[low]
	array[low] = key
	array[pivot] = tmp2
	return low
	
def qs(array,low,high):
	if low < high:
		k = sub_sort(array,low,high)
		qs(array,low,k)
		qs(array,k+1,high)
	
if __name__ == '__main__':
	array = [4,6,7,9,3,2,5,1,99,98]
	print(array)
	qs(array,0,len(array)-1)
	print(array)
		


# arr = [3,1,2,4,8,6,9,10]
# arrlen = len(arr)
#print(arrlen)


# def divide(start, end):
	# i = start
	# j = end
	# pivot = 0
	# while i < j:
		# if arr[j] > arr[pivot]:
			# j -= 1
		# if arr[i] < arr[pivot]:
			# i += 1
		# print('arr[j]:',arr[j])
		# print('arr[i]:',arr[i])
	# else:
		# if i == j:
			# tmp = arr[j-1]
			# arr[j-1] = arr[pivot]
			# arr[pivot] = tmp
			# print('j:ddd')
			# print('i:ddd')
			# pivot = j-1
			# print(pivot, arr[pivot])
			# return pivot
		# else:
			# tmp = arr[j]
			# arr[j] = arr[i]
			# arr[i] = tmp
			# print('j:ccc')
			# print('i:ccc')
			
# p = divide(1, arrlen-1)

# print(arr)
