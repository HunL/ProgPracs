/*
 * pirntLongestLine.c
 *
 *  Created on: 2016年8月17日
 *      Author: j
 */

#include<stdio.h>
#include<string.h>
#include<stdlib.h>

#define MAX_INPUT 1001

int main(void)
{
//	char str[MAX_INPUT];
	char input[MAX_INPUT];
	char longest_str[MAX_INPUT];
//	int max = 0;
	int longest_len = 0;
	int len;
//	int times = 5;

	//如何从标准输入中输入文件尾->windows下ctrl+z
//	while(times != 0 && (gets(input) != NULL))
	while(gets(input) != NULL)
	{
		len = strlen(input);

		if(len > longest_len)
		{
//			int i;
//			for(i=0;i<len;i++)
//				longest_str[i] = str[i];
			strcpy(longest_str, input);
			longest_len = len;
		}
//		printf("string %s, len %d\n", input, len);
//		printf("string %s, longest %d\n", longest_str, longest_len);
//		times--;
	}

	if(longest_len > 0)
		puts(longest_str);
	printf("\npress any key to exit...");
	getchar();

	return EXIT_SUCCESS;
}
