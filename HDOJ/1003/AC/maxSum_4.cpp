/*
贪心算法
*/
#include<iostream>
#include<limits.h>

using namespace std;

int main()
{
  int k,T,N;
  cin >> T;
  for(k = 1;k <= T;k++)
  {
    cin >> N;
    int max = INT_MIN;
    int i,j,fig,begin,end,sum;
    for(i = j = 0,sum = 0;j < N;j++)
    {
      cin >> fig;
      sum += fig;
      if(sum > max)
      {
        max = sum;
        begin = i+1;
        end = j+1;
      }
      if(sum < 0)
      {
        i = j+1;//重新定义字串的起点
        sum = 0;
      }
    }
    cout << "Case " << k << ":" << endl;
    cout << max << " " << begin << " " << end << endl;
    if(k != T)
      cout << endl;
  }

  return 0;
}
