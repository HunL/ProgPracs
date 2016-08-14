/*
 * countNum.c：统计字符串中各字符出现的次数
 *
 *  Created on: 2016年8月13日
 *      Author: j
 */
#include "stdio.h"
#include "string.h"
#include "stdbool.h"
#include "stdlib.h"

// 字符无法转成字符串？
#if 0

int main()//todo:time count
{
	char* str = "hello jeffluckyboy";

	char* diffChar = "";
	diffChar = (char*)malloc(strlen(str)+1);
	if(diffChar==NULL)
		fprintf(stderr, "malloc failed...！\n");

	int i,j;
	for(i=0;i<strlen(str);i++)
	{
		char k = str[i];
		int sum = 0;

		printf("diffChar %s\n", diffChar);
		if(strchr(diffChar, k)==NULL)
		{
			printf("diffChar null\n");
			for(j=i;j<strlen(str);j++)
			{

				if(str[j]==k)
					sum+=1;
			}
			diffChar[i]=k;//?
			printf("k = %c, sum = %d\n", k, sum);
		} else {
			printf("not null, k = %c\n", k);
		}
	}
	free(diffChar);

	return 0;
}

#endif



