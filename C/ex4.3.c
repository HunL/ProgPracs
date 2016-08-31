/*
 * ex4.3.c
 *
 *  Created on: 2016Äê9ÔÂ1ÈÕ
 *      Author: j
 */

#include "stdio.h"
#include "stdlib.h"

int main(void)
{
	int i,j,k;
	printf("please input the three sides of the triangle:\n");
	scanf("%d %d %d", &i, &j, &k);
	printf("three sides are: %d, %d, %d\n",i, j, k);

	if((i==j) && (j==k))
	{
		printf("Equilateral Triangle!\n");
	} else if(((i==j) && (i!=k)) || ((i==k) && (j!=k)) || ((j==k) && (i!=k)))
		{
			printf("Equicrural Triangle!\n");
		} else if((i!=j) && (i!=k) && (j!=k))
				{
					if((i+j<k) || (i+k<j) || (k+j<i))
						{
							printf("Not A Triangle\n");
						} else {
							printf("Normal Triangle\n");
						}
				}

	printf("\npress any key to exit...");
	getchar();
	return EXIT_SUCCESS;
}


