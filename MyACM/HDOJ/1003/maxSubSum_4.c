/*
 * =====================================================================================
 *
 *       Filename:  maxSubSum_4.c
 *
 *    Description:  用动态规划思想解决最大子串和问题。
 *                  时间复杂度：O(n)
 *
 *        Version:  1.0
 *        Created:  07/27/2011 12:20:38 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Jiehui Li, ljiehui0826@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */

#include<stdio.h>
#include<stdlib.h>

int maxSubSum(int *a,int len)
{
  int i;
  int curSum = 0;
  int maxSum = a[0];

  for(i = 0;i < len;i++)
  {
    curSum += a[i];

    if(curSum > maxSum)
    {
      maxSum = curSum;
    }

    if(curSum < 0)
    {
      curSum = 0;
    }
  }

  printf("maxSum = %d\n",maxSum);
  return maxSum;
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
    scanf("%d",a+i);
  }

  maxSubSum(a,len);
  free(a);

  return 0;
}
