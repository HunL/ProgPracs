//example by:http://tuodeyi.blog.163.com/blog/static/17651458020114604138143/

// source code of submission 750618, Zhongshan University Online Judge System
#include <iostream>
#include <string>
#include <cstring>
using namespace std ;

//消除结果中前面的0 
string clear(string s1)
{
    if(s1 == "") return s1 = "0" ;

    int zero = 0 ;
    for(int i = 0 ; i < s1.size() ; i ++)
    {
        if(s1[i] == '0')
            zero ++ ;
        else break ;
    }
    s1 = s1.substr ( zero , s1.size ()-zero ) ;
    if(s1 == "") s1 = "0" ;
    return s1 ;
}


string addition(string s1, string s2) 
{
    int i ;

    while(s1.size() < s2.size() ) s1 = '0' + s1 ;
    while(s2.size() < s1.size() ) s2 = '0' + s2 ;

    s1 = '0' + s1 ;
    s2 = '0' + s2 ;
    for(i = s1.size () - 1 ; i > 0  ; i--)
    {
        s1[i] = s1[i] + s2[i] - '0' ;
        if(s1[i] - '0' > 1)
        {
            s1[i] = s1[i] - 2 ;
            s1[i-1] = s1[i-1] + 1 ;
        }
    }

    return s1 = clear(s1) ;
}


int main()
{
    int k ;
    cin >> k ;
    int c = 1 ;
    while(k--)
    {
        string str1 , str2 ;
        cin >> str1>> str2 ;
        string str3 = addition(str1,str2) ;
        cout << c++ << " "<< str3 << endl;
    }
}
