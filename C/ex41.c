/*
 * ex41.c
 *
 *  Created on: 2016Äê8ÔÂ31ÈÕ
 *      Author: j
 */

#include "stdio.h"
#include "stdlib.h"

int main(void)
{
	int input, loop_times;
	scanf("%d %d", &input, &loop_times);
	printf("input£º%d, loop_times: %d\n", input, loop_times);

	int i = 0;
	float ai;
	while(i<loop_times)
	{
		if(i==0)
			ai = 1;
		else{
			ai = (ai+input/ai)/2;
		}
		printf("a%d: %f\n", i, ai);
		i++;
	}

	printf("\npress any key to exit...");
	getchar();
	return EXIT_SUCCESS;
}


