/*
 * 2.8.2_check_double_brace.c
 *
 *  Created on: 2016Äê8ÔÂ15ÈÕ
 *      Author: j
 */
#include<stdio.h>

int main(void)
{
	int counter = 0;
	char ch;

	while((ch = getchar()) != EOF)
	{
		if(ch == '{')
		{
			counter = 1;
		}

		if(ch == '}')
		{
			counter -= 1;
		}
	}

	if(counter != 0)
	{
		printf("check braces fails...!\n");
	} else {
		printf("check braces succeed!\n");
	}

	getchar();
	return 0;
}
