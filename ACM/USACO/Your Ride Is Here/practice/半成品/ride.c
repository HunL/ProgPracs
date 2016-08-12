/*
ID:ljiehui1
LANG:C
PROG:ride
*/
#include<stdio.h>
main(){
	FILE *fin = fopen("ride.in","r");
	FILE *fout = fopen("a.out","w");
 	char a,b,c,d,e,f;
	fscanf(fin,"%c %c %c %c %c %c",&a,&b,&c,&d,&e,&f);
	fprintf(fout,"%d\n",(a-64)*(b-64)*(c-64)*(d-64)*(e-64)*(f-64));
        exit(0);
}
