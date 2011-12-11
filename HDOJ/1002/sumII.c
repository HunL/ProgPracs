/*
 * =====================================================================================
 *
 *       Filename:  sumII.c
 *
 *    Description:  大数相加。通过while(scanf("%d",&N))语句实现输入，太NB了。
 *                  参考链接：http://www.cnblogs.com/cpoint/archive/2011/04/13/2015210.html
 *
 *        Version:  1.0
 *        Created:  07/19/2011 08:35:34 AM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Jiehui Li, ljiehui0826@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */

#include<stdio.h>
#include<string.h>
#include<stdlib.h>

#define MAX 1001

void add(char a[],char b[])
{

  int temp = 0,flg = 0;
  int len_a = strlen(a);
  int len_b = strlen(b);
  int i = len_a;
  int j = len_b;
  char sum[MAX] = {' '};

  for(;i > 0;i--)
  {
    if(j > 0)
    {
      temp = a[i - 1] + b[j - 1] + flg - 96;
      j--;
    }
    else
      temp = a[i - 1] + flg - 48;
    if(temp >= 10)
    {
      flg = 1;
    }
    temp %= 10;
    sum[i] = temp + 48;
  }

  if(flg == 1)
  {
    sum[0] = 49;
  }

  i = 0;
  while(i <= len_a)
  {
    if(sum[i] != ' ')
      printf("%c",sum[i]);
    i++;
  }
  printf("\n");

}

int main()
{

  int N,i;
  while(scanf("%d",&N))
  {
    for(i = 1;i <= N;i++)
     {
       char a[MAX],b[MAX];
       scanf("%s %s",a,b);
       int LEN_a = strlen(a),LEN_b = strlen(b);
       printf("Case %d:\n",i);
       if(LEN_a >= LEN_b)
          add(a,b);
       else
          add(b,a);
       if(i != N)
         printf("\n");
     }
  }
  return 0;
}
