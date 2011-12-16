#include<iostream>

using namespace std;

int main()
{
  int t,m,n,sum,tmp1,tmp2;
  cin >> t;
  while(t > 0)
  {
    cin >> n >> m;
    tmp1 = (n + 1)*n/2;
    tmp2 = (m + 1)*m/2;
    sum = tmp1*tmp2;
    cout << sum << endl;
    t--;
  }
  
  return 0;
}
