/*
 * 1.8.2.inputoutput.c
 *
 *  Created on: 2016��8��14��
 *      Author: j
 */

//eclipse cdt��ô��������̨��������->��������gcc

#include<stdio.h>

int main(void)
{
	int line;
	char ch;
	int at_beginning;

	line = 0;
	at_beginning = 1;

	printf("please input : ");
	while((ch = getchar()) != EOF)
	{
		if(at_beginning == 1)
		{
			at_beginning = 0;
			line+=1;
			printf("line %d: ", line);
		}

		putchar(ch);
		if(ch == '\n')
		{
			at_beginning = 1;
			printf("please input : ");
		}
	}

	return 0;
}
