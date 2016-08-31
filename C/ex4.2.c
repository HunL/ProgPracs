/*
 * ex4.2.c
 *
 *  Created on: 2016Äê8ÔÂ31ÈÕ
 *      Author: j
 */

#include "stdio.h"
#include "stdlib.h"

int main(void)
{
	printf("1\n2\n");
	int member;
	for(member = 3; member < 101; member += 2)
	{
		int divisor;
		for(divisor = 3; divisor <= member/2; divisor += 2)
		{
			if(member % divisor == 0)
				break;
		}

		if(member/2 <= divisor)
			printf("%d\n", member);
	}

	printf("\npress any key to exit...");
	getchar();
	return EXIT_SUCCESS;
}


