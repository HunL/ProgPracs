//求余，找规律。
#include<iostream>
//#include<math>

using namespace std;

int main()
{
  int a,b,n;
  //int f1 = 1,f2 = 1;
  while(cin >> a >> b >> n && n != 0) 
  {
    //如何递归？
 //   f = (a * f1 + b * f2) % 7;  

    int f1 = 1,f2 = 1;
    int t = 3,f;
    while(t <= n)
    {
      f = (a * f1 + b * f2) % 7;  
      f2 = f1;
      f1 = f;
      t++;
    }
  cout << f << endl;
    //n--;
  }
  //cout << f << endl;

  return 0;
}

