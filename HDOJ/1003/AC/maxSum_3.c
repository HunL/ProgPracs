/*
 * 此方法我认为正确，可是不能AC
 * 第三种做法：动规，时间复杂度0(N) 仍可进一步减少空间复杂度
*/
#include<stdio.h>
#include<limits.h>//for INT_MIN

#define MAX 100000
//#define MAX 100
int a[MAX],n;
int Start[MAX],All[MAX];
int first,last;
int temp;

int max(int x,int y)
{
  return (x > y)?x:y;
}

int maxSum(int a[],int n)
{
  Start[n-1] = a[n-1];
  All[n-1] = a[n-1];
  
  for(temp = n-2;temp >= 0;temp--)
  {
    Start[temp] = max(a[temp],a[temp]+Start[temp+1]);
    All[temp] = max(Start[temp],All[temp+1]);
    if(Start[temp] == a[temp])
      last = temp+1;
    if(All[temp] == Start[temp])
      first = temp+1;
  }
  return All[0];
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
