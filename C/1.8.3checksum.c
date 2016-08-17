/*
 * checksum.c
 *
 *  Created on: 2016Äê8ÔÂ17ÈÕ
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
