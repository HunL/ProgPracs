#include<string.h>
#include<stdio.h>
char a[21][100001],b[21][100001],f[21][1000001];int c[1000000];
int main()
{
    int i,j,d,l1,l2,m;
    scanf("%d",&i);
        getchar();
        for(j=1;j<=i;j++)
        {    
            scanf("%s%s",a[j],b[j]);
l1=strlen(a[j])-1;
l2=strlen(b[j])-1;
for(m=1;m<=100000;m++)
{
    c[m]=0;
}
m=0;
while(l1>=0&&l2>=0)
{
    c[++m]+=a[j][l1]-'0'+b[j][l2]-'0';
    if(c[m]>=10)
    {
        c[m+1]+=c[m]/10;
        c[m]=c[m]%10;
    }
l1--;l2--;
}
if(l1+1==0&&l2+1==0&&c[m+1]!=0)++m;
    while(l1>=0)
    {
    c[++m]+=(int)a[j][l1]-48;l1--;
}
while(l2>=0)
    {
    c[++m]+=(int)b[j][l2]-48;l2--;
}            
            for(d=m;d>=1;d--)
                {
                f[j][m-d]=c[d]+'0';
            }
        }    
        for(j=1;j<=i;j++)
            {
            printf("Case %d:\n",j);
            printf("%s + %s = ",a[j],b[j]);
            printf("%s\n",f[j]);
                if(j!=i)printf("\n");
            }
}
