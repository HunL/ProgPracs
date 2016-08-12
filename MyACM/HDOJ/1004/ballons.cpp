/*
 *要学会使用数据结构，参考http://www.iiacm.net/archives/97，学会使用qsort,strcmp
 * */
#include<iostream>
#include<string>
#include<cstdlib>//for system function qsort
#include<cstring>//for system function strcmp

//#define MAX 15

using namespace std;

int cmp(const void * x,const void * y)
{
  return strcmp((char *)x,(char *)y);//若x和y相同则返回0
}

int main()
{
  int N;
  cin >> N;
//  string a[MAX];
  string a;
  while(N > 0)
  {
    cin >> a;
    cout << a << endl;
    N--;
    if(N == 0)
    {
      cin >> N;
    }
  }
  return 0;
}
