/*
 * checksum.c
 *
 *  Created on: 2016��8��17��
 *      Author: j
 */

#include<stdio.h>

void main(void)
{
	signed char sum = -1;
	char ch;

	while((ch = getchar()) != '\n')
	{
		printf("%c", ch);
		sum += ch;
	}

	sum += '\n';
	printf(" checksum = %d \npress any key to exit...", sum);
	getchar();
}
