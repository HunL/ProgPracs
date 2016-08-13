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

/*

bool is_contain(char a, char* s)
{

	char dest = 'A';
	char* k = "AB";

	int i;
//	if(strstr(k,dest) == NULL)
	for(i=0;i<strlen(s);i++)
	{
		if(s[i]==a)
		{
			printf("000\n");
			return true;
		}
	}

	return false;
}
*/

int main()//todo:time count
{
	char* str = "hello jeffluckyboy";
//	printf("len = %d\n", strlen(str));

//	char* diffChar = "helo jfuckyb";
	char* diffChar = "";
	diffChar = (char*)malloc(strlen(str)+1);
	if(diffChar==NULL)
		fprintf(stderr, "malloc failed...！\n");
//	printf("diffChar %s\n", diffChar);

	int i,j;
	for(i=0;i<strlen(str);i++)
	{
		char k = str[i];
		int sum = 0;

//		if(is_contain(k, diffChar))
//			continue;
		printf("diffChar %s\n", diffChar);
		if(strchr(diffChar, k)==NULL)
		{
			printf("diffChar null\n");
			for(j=i;j<strlen(str);j++)
			{

				if(str[j]==k)
					sum+=1;
			}
//			*(diffChar+sizeof(diffChar)/sizeof(int)-1+i)=k;
			diffChar[i]=k;//?
			printf("k = %c, sum = %d\n", k, sum);
		} else {
			printf("not null, k = %c\n", k);
			//continue;
		}
	}
	free(diffChar);

	/*char* dest = "A";
	char* k = "B";
	char* tmp;
	tmp = (char*)malloc(strlen(dest)+strlen(k)+1);
	if(tmp==NULL)
		fprintf(stderr, "malloc failed...\n");

	strcpy(tmp,dest);
	strcat(tmp,k);
	printf("dest %s\n", dest);
	printf("k %s\n", k);
	printf("tmp %s\n", tmp);

	free(tmp);*/
	/*
	bool b;
	b = is_contain('A',"AB");
	printf("b=%d\n",b);*/

	return 0;
}





