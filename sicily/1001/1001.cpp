#include<iostream>  
 using namespace std;  
 #include<string>  
 int main()  
 {  
     string s;  
     while(cin>>s&&s!="0")  
     {  
         long long count=0;  
         long long f1=0,f2=0;  
         int i;  
         for( i=0;i<s.size();i++)  
         {  
             if(i==0) {f1=1;count=1;}  
         if(i==1){  
                 if((s[0]-'0')*10+s[1]-'0'<=26&&(s[0]-'0')*10+s[1]-'0'>10&&s[i]!='0'){f2=2;count=2;}else {f2=1;count=1;}  
             }  
         if(s[i]=='0'&&i>1)  
         {  
         count=f1;  
         f1=f2;  
         f2=count;  
         }  
         if(i>1&&s[i]!='0')  
             {  
                 if((s[i-1]-'0')*10+s[i]-'0'<=26&&(s[i-1]-'0')*10+s[i]-'0'>10)  
                 {  
                     count=f2+f1;  
                     f1=f2;  
                     f2=count;  
                 }  
                 else{  
                         count=f2;  
                         f1=f2;  
                         f2=count;  
                     }  
             }  
         }  
         cout<<count<<endl;  
     }  
     return 0;  
 }
