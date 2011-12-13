#include<iostream>
using namespace std;

int main()
{
      int k, n;
      int res, ex,flag;
      while(cin>>k>>n)
      {
            ex=1;res=0;
            while(n>0)
                 {
                      flag = n%2;
                      res+=flag*ex;
                      n=n/2;
                      ex=ex*k;
                 }
             cout<<res<<endl;
      }
}
