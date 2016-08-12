/*
 * =====================================================================================
 *
 *       Filename:  A+B.c
 *
 *    Description:  language C
 *
 *        Version:  1.0
 *        Created:  07/16/2011 09:36:43 AM
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
  freopen("A+B.in","r",stdin);
  freopen("A+B.out","w",stdout);

  int a,b;
  while(scanf("%d %d",&a,&b)!=EOF)
  {
    printf("%d\n",a+b);
  }
  
  fclose(stdin);
  fclose(stdout);
  return 0;
}
