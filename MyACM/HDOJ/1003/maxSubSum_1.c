/*
使用朴素算法
时间复杂度：O(n^3)
*/

#include<stdio.h>
#include<stdlib.h>

/*
输入字串 --> 用求最大字串算法处理 --> 输出结果
求最大字串算法:
1、朴素算法
2、优化后算法
3、二分法
4、动态规划算法
*/

//if not define algor1..define.

int subSum(int* a,int left,int right)
{
	int i,sum = 0;
	for(i = left;i <= right;i++)
	{
		sum += a[i];
	}

	return sum;
}

int maxSubSum(int* a,int len)
{
	int i,j,curSum;
	int sum = a[0];
	for(i = 0;i < len;i++)
	{
		curSum = 0;
		for(j = i;j < len;j++)
		{
			curSum = subSum(a,i,j);

			if(curSum > sum)
			{
				sum = curSum;
			}
		}
	}
	printf("maxSubSum为：%d\n",sum);

	return sum;
}

int main()
{
	int *a;
	int i,len;
	printf("请输入数组长度：\n");
	scanf("%d",&len);
	a = (int *)malloc(len * sizeof(int));
	printf("请输入数组元素：\n");

	for(i = 0;i < len;i++)
	{
//		printf("请输入数组：\n");
		scanf("%d",a+i);
	}
	maxSubSum(a,len);

	free(a);
	return 0;
}
