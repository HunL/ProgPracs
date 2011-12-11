#include "stdio.h"
#include "stdlib.h"

#define S 20

struct peo
{
            char name[S];
            int get;
            int have;
            int m;
}p[11];

int n;

int scmp(char *s1,char *s2)
{
	    int i=0;
            if(strlen(s1)!=strlen(s2))
			        return 0;
	    while(s1[i]!='\0') {
				 if(s1[i]!=s2[i])
				            return 0;
				 i++;
	     }
	     return 1;
}

int getname(char *s)
{
	    int i;
	    for(i=0;i<n;i++){
	          if(scmp(s,p[i].name))
			                return i;
	    }
}

void input()
{
	    int i,j,t;
	    int now;
	    scanf("%d",&n);
	    for(i=0;i<n;i++){
			 scanf("%s",p[i].name);
			 getchar();
			 p[i].get=0;
	    }
}

void work()
{
	    int i,j,t;
	    char str[S];
    	    int now,tmp;
	    for(t=0;t<n;t++){
		         scanf("%s",str);
		         getchar();
		         now=getname(str);
		         scanf("%d%d",&p[now].have,&p[now].m);
		         getchar();
		         for(i=0;i<p[now].m;i++){
				      scanf("%s",str);
				      getchar();
				      tmp=getname(str);
				      p[tmp].get+=p[now].have/p[now].m;
			 }
		         if(p[now].m)
 			p[now].get+=p[now].have%p[now].m;
	    }
}

void output()
{
	    int i;
	    for(i=0;i<n;i++){
		            printf("%s %d\n",p[i].name,p[i].get-p[i].have);
	    }
}

int main()
{
	    freopen("gift1.in","r",stdin);
	    freopen("gift1.out","w",stdout);
	    input();
	    work();
	    output();
	    exit(0);
}
