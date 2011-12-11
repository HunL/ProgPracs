/*
ID:ljiehui1
PROG:gift1
LANG:C
*/
#include<stdio.h>

struc peo  //定义p[]的结构变量；
{
	char name[10];
	int MoGet;
	int MoHave;
	int PeoDi;
}p[10];

int scmp(char *s1,char *s2)  //根据名字的字符串是否一致来判断把钱分给那个人；
{
	int i;
	for(i=0;i<10;i++)
	{
		if(strlen(s1)!=strlen(s2))
			return 0;
		while(s1[i]!='\0')
		{if(s1[i]!=s2[i])
			return 0;
			i++;
		}
	}
	return 1;
}


int mod(int i,j)
{
	return (i/j);
}
int sum=0,dsum;
dsum=i%j;
main()
{
	FILE *fin = fopen("gift.in","r");
	FILE *fout = fopen("gift.out","w");
	char a[];
	fscanf(fin,"%s",a);
	fprintf(fout,"%s",a);
}
