 #include <iostream>  
    #include<algorithm>  
    #include<map>  
    #include<string>  
    #include<cstdio>  
    using namespace std;  
    int main()  
    {  
        //freopen("1.txt","r",stdin);  
        int n;  
        while(scanf("%d",&n)&&n)  
        {  
            int max=0;  
            map<string,int> mp;  
            map<string,int>::iterator p;  
            string s;  
            for(int i=0;i<n;i++)  
            {  
                cin>>s;  
                mp[s]++;  
            }  
            for(p=mp.begin();p!=mp.end();p++)  
            {  
                if(p->second>max)  
                {  
                    s=p->first;  
                    max=p->second;  
                }     
            }  
            cout<<s<<endl;  
        }  
    }  
