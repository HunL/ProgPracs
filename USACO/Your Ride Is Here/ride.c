/*
ID:ljiehui1
LANG:C
PROG:ride
*/
#include<stdio.h>
int change(char c)
{
	int k=c;
	return (k-64);
}
int sum(char a[])
{
	int msum=1,j;
	int i=0;
	char c=a[i];
	while(c!='\0')
	{
		j=change(c);
		msum*=j;
		i++;
		c=a[i];
	}
	return msum;
}
int main(){
      FILE *fin = fopen("ride.in","r");
      FILE *fout = fopen("ride.out","w");
      char group[10],comet[10];
      fscanf(fin,"%s%s",group,comet);
      int gsum,csum;
      gsum=sum(group);
      csum=sum(comet);
      if(gsum%47==csum%47) fprintf(fout,"%s\n","GO");
      else  fprintf(fout,"%s\n","STAY");
      fclose(fin);
      fclose(fout);
      return 0;
}

