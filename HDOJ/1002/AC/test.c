/*
 * =====================================================================================
 *
 *       Filename:  test.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  12/11/2011 12:10:48 AM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Qingping Hou (houqp), dave2008713@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

void print(char a[])
{
  int len_a = strlen(a);
  int i = len_a;
  
  for(;i > 5;i--)
  {
    printf("%c",a[i]);
  }
}
 
int main()
{
  char a[] = "abcdefghijklmn";
  print(a);
  return 0;
}
