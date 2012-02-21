//学会STL中的map！！

#include<iostream>
#include<map>
#include<string>

using namespace std;

int main()
{
  //map<string, int> color;
  map<string, int> ballons;
  map<string, int>::iterator a;

  int i,N;

  while(cin >> N && N != 0)
  {
      ballons.clear();//清空容器
    //map<string, int> ballons;//定义一个名为ballons的map对象
    //map<string, int>::iterator a;//a是一个pair<const string,int>类型的对象

    //如何将测试用例输入到map中？
    for(i = 0;i < N;i++)
    {
      // cout << 2;
      string color;
      cin >> color;

      //ballons.clear();//清空容器

      a = ballons.find(color);
      if(a == ballons.end())
      {
      // cout << 1;
         //ballons.insert(make_pair(color,1));
         ballons.insert(map<string, int>::value_type(color,1));
         //ballons[color]++;
      }
      else
        (a->second)++;
    }

      //cout << a->first << " " << a->second << endl;

      a = ballons.begin();
      string maxstr = a->first;//注意！输出个数最多的颜色！
      int maxnum = a->second;
      while(a != ballons.end())
      {
        if(a->second > maxnum)
        {
          maxnum = a->second;
          maxstr = a->first;
        }
        a++;
      }

      //cout << maxstr << endl;
    //}
      cout << maxstr << endl;
   // }
  }

  return 0;
}
