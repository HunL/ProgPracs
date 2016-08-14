/*
 * 1.8.2.inputoutput.c
 *
 *  Created on: 2016年8月14日
 *      Author: j
 */

//eclipse cdt怎么调出控制台进行输入

#include<stdio.h>

int main(void)
{
	int line;
	char ch;
	int at_beginning;

	line = 0;
	at_beginning = 1;

	while((ch = getchar()) != EOF)
	{
		if(at_beginning == 1)
		{
			at_beginning = 0;
			line+=1;
			printf("%d: ", line);
		}

		putchar(ch);
		if(ch == '\n')
			at_beginning = 1;
	}

	return 0;
}
