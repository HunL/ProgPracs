/*
 * =====================================================================================
 *
 *       Filename:  sum.c
 *
 *    Description:  大数相加
 *
 *        Version:  1.0
 *        Created:  12/10/2011 11:38:46 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Qingping Hou (houqp), dave2008713@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */
#include<stdio.h>
#include<string.h>

#define MAX 1000

void add(char a[],char b[])
{
  int len_a = strlen(a);
  int len_b = strlen(b);
  int i = len_a;
  int j = len_b;
  int temp = 0;
  int flg = 0;
  char sum[MAX] = " ";

  for(;i > 0;i--)//最后一位是a[i-1]和b[j-1]
  {
    if(j > 0)
    {
      temp = a[i-1] +b[j-1] + flg - 96; 
      j--;
    }
    else
      temp = a[i-1] + flg - 48;
    
//    if(temp > 10)//1+9不产生进位，SB
    if(temp >= 10)
    {
      flg = 1;
    }
    temp %= 10;
    sum[i] = temp + 48;//将a[i-1]和b[j-1]的个位保存在sum[i]
  }

  if(flg == 1)
  {
//    sum[0] = 1;
    sum[0] = 49;
  }

  i = 0;
  while(i <= len_a)
  {
    printf("%c",sum[i]);
    i++;
  }
}

int main()
{
  int i,n;
  char a[MAX],b[MAX];
  while(scanf("%d",&n) != EOF)
  {
    for(i=0;i < n;i++)
    {
      scanf("%s %s",a,b);
      int len_a = strlen(a),len_b = strlen(b);
      int k = len_a,j = len_b;
      printf("Case %d:\n",i+1);
      printf("%s + %s = ",a,b);
      if(k >= j)
      {
       add(a,b);
      }
      else
      {
       add(b,a);
      }
       printf("\n\n");
    }
  }
  
  return 0;
}
//int main()
//{
//  int i,n;
//  int a,b;
//  while(scanf("%d",&n) != EOF)
//  {
//    for(i=0;i < n;i++)
//    {
//      scanf("%d %d",&a,&b);
//      printf("Case %d:\n",i+1);
//      printf("%d + %d = %d\n\n",a,b,a+b);
//    }
//  }
//  return 0;
//}
