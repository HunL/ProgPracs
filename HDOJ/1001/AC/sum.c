/*
 * =====================================================================================
 *
 *       Filename:  sum.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  12/10/2011 11:00:21 AM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author: 
 *        Company:  
 *
 * =====================================================================================
 */
#include<stdio.h>

int main()
{
  int i,n;
  int sum;
  while(scanf("%d",&n) != EOF)
  {
    sum = 0;
    for(i=1;i<=n;i++)
    {
      sum += i;
    }

  printf("%d\n\n",sum);
  }

  return 0;
}
