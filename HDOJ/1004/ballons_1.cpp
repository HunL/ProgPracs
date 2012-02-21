/*
 *要学会使用STL，参考http://betabone.com/archives/166.html和http://www.iiacm.net/archives/97
 * */
#include<iostream>
#include<string>

//#define MAX 15

using namespace std;

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
