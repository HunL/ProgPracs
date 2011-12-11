/*
 * =====================================================================================
 *
 *       Filename:  sum.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  07/16/2011 12:08:00 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Jiehui Li, ljiehui0826@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */

#include<stdio.h>

int main()
{
  freopen("sum.in","r",stdin);
  freopen("sum.out","w",stdout);

  int i,a,sum=0;
   while(scanf("%d",&a)!=EOF)
  {
      for(i = 1;i <= a;i++)
      {
        sum += i;
      }
      printf("%d\n\n",sum);
      sum = 0;//i love u,simple mistake.
  }
  
  fclose(stdin);
  fclose(stdout);

  return 0;
}
