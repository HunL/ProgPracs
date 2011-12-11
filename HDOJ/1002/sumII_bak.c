/*
 * =====================================================================================
 *
 *       Filename:  sumII.c
 *
 *    Description:  无法处理大数相加。
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

int main()
{
  freopen("sumII.in","r",stdin);
  freopen("sumII.out","w",stdout);

  int a,b,c,i=0;
  scanf("%d",&a);
  for(i = 1;(scanf("%d %d",&b,&c)!=EOF);i++)
  {
    printf("Case %d:\n",i);
    printf("%d + %d = %d\n\n",b,c,b+c);
  }

  fclose(stdin);
  fclose(stdout);

  return 0;
}
