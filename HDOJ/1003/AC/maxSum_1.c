/*
第一种做法：时间复杂度0(N^3)  Time Limit Exceeded
*/
#include<stdio.h>
#include<limits.h>//for INT_MIN

#define MAX 100000

int a[MAX],n;
int first,last;

int maxSum(int a[],int n)
{
  int i,j,k;
  int maximum = INT_MIN;
  int sum;
  for(i = 0;i < n;i++)
  {
    for(j = i;j < n;j++)
    {
      sum = 0;//每一轮循环，子串应该从第j位开始重新计算
      for(k = i;k <= j;k++)
      {
        sum += a[k];
      }
      if(sum > maximum)
      {
        maximum = sum;
        first = i+1;
        last = k;
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
    printf("Case %d:\n",++tmp);
    printf("%d",maxSum(a,N));
    printf(" %d %d\n",first,last);
    if(T > 0)
      printf("\n");
  }
  return 0;
}
