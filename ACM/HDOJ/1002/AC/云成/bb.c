/*
 * =====================================================================================
 *
 *       Filename:  bb.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  12/11/2011 02:29:39 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Qingping Hou (houqp), dave2008713@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */
#include<stdio.h>

int main()
{
  int k,N;
  int res,ex,flg;

  scanf("%d %d",&k,&N);
  ex = 1;
  res = 0;
  while(N > 0)
  {
    flg = N%2;
    res += flg*ex;
    N = N/2;
    ex = ex*k;
  }
  printf("%d\n",res);

  return 0;
}
