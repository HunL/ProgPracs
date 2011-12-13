/*
 * 第二种做法：时间复杂度0(N^2)   Time Limit Exceeded	
*/
#include<stdio.h>
#include<limits.h>//for INT_MIN

#define MAX 100000
//#define MAX 100
int a[MAX],n;
int first,last;

int maxSum(int a[],int n)
{
  int i,j;
  int maximum = INT_MIN;
//  printf("%d\n",maximum);
  int sum;
  for(i = 0;i < n;i++)
  {
      sum = 0;//每一轮循环，子串应该从第i位开始重新计算
    for(j = i;j < n;j++)
    {
        sum += a[j];
      if(sum > maximum)
      {
        maximum = sum;
        first = i+1;
        last = j+1;
      }
    }
  }
  return maximum;
}

int main()
{ 
  int i,N,T;
  int tmp = 0;
  scanf("%d",&T);
  while(T > 0)
  {
    scanf("%d",&N);
    for(i = 0;i < N;i++)
    {
      scanf("%d",&a[i]);
    }
    T--;
//    for(i = 0;i < N;i++)
//    {
//      printf("%d ",a[i]);
//    }
//    printf("\n");
    printf("Case %d:\n",++tmp);
    printf("%d",maxSum(a,N));
    printf(" %d %d\n",first,last);
    if(T > 0)
      printf("\n");
  }
  return 0;
}
