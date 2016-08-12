/*
使用分治法
时间复杂度：O(nlgn)
分：使用二分法分段
治：见_maxSubSequenceSum2描述
*/

#include<stdio.h>
#include<stdlib.h>

/*含左边界元素的子序列和中的最大值*/
int maxLeftSubSum(int *a,int left,int right)
{
	int i;
	int curSum = 0;
	int maxSum = a[left];
	for(i = left;i <= right;i++)
	{
		curSum += a[i];
		if(curSum > maxSum)
		{
			maxSum = curSum;
		}
	}

	return maxSum;
}

/*含右边界元素的子序列和中的最大值*/
int maxRightSubSum(int *a,int left,int right)
{
	int i;
	int curSum = 0;
	int maxSum = a[right];
	for(i = right;i >= left;i--)
	{
		curSum += a[i];
		if(curSum > maxSum)
		{
			maxSum = curSum;
		}
	}

	return maxSum;
}

/*求三个数中最大者*/
int max(int a,int b,int c)
{
	int numMax;
	if(a >= b && a >= c)
	{
		numMax = a;
	}else if(b >= a && b >= c)
	{
		numMax = b;
	}else if(c >= a && c >= b)
	{
		numMax = c;
	}

	return numMax;
}

/*分治法求序列某段子序列中子序列和最大值*/
int maxSubSum(int *a,int left,int right)
{
	int center;
	int leftMaxSum;
	int rightMaxSum;
	int leftBoundMaxSum;
	int rightBoundMaxSum;

	center = (left + right) >> 1;

	if(left == right)
	{
		return a[left];
	}

	leftMaxSum = maxSubSum(a,left,center);
	rightMaxSum = maxSubSum(a,center + 1,right);
	leftBoundMaxSum = maxRightSubSum(a,left,center);
	rightBoundMaxSum = maxLeftSubSum(a,center + 1,right);

 	return max(leftMaxSum,rightMaxSum,leftBoundMaxSum + rightBoundMaxSum);
}

int main()
{
	int *a,i,len;
	printf("请输入数组长度：\n");
	scanf("%d",&len);
	a = (int *)malloc(len * sizeof(int));
	printf("请输入数组元素： \n");
	for(i = 0;i < len;i++)
	{
		scanf("%d",a+i);
	}

	printf("maxSubSum = %d\n",maxSubSum(a,0,len));

	return 0;
}
