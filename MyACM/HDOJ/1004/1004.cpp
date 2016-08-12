//http://www.iiacm.net/archives/97

#include<iostream>
#include<string>
#include<string.h>
#include<stdlib.h>
 
using namespace std;

//学会使用这种方法
//使用库函数qsort，现在才看到的库函数，人家2B学校的在大一就分析透彻了，NND，我要付出怎样的努力才能追上啊！！
int cmp(const void *x,const void *y)
{
  return strcmp((char*)x,(char*)y);
}

int main()
{
  int i,j,N,max,tmp;
  //定义一个字符串数组
  char color[1001][16];
  //cin >> N;
  while(cin >> N && N != 0)
  {
    for(i = 0;i < N;i++)
    //while(N)
    {
       cin >> color[i];
       //cout << color[i] << endl;
       //N--;
       //i++;
       //用以下语句输入N是错误的
       //if(i == N-1)
       //{
       //  cin >> N;
       //}
    }
    qsort(color,N,sizeof(color[0]),cmp);
    max =  i = j = 0;//max用来记录最多的颜色的个数
    //for(i = 0;i < N;i++)
    while(1)//要学会这种方法！
    {
      if(!strcmp(color[i],color[j]))
        j++;
        //cout << color[i] << " " << color[j] << endl;
      else
      {
        if(j - i > max)
        {
          max = j - i;
          //i = j;
          tmp = i;

        }
        i = j;//attention!j递增的基础！
      }
      if(j == N)
      {
        if(j - i > max)
        {
          max = j - i;
          tmp = i;
        }
        break;
      }
    }
    //cout << max << " " << color[tmp] << endl;
    cout << color[tmp] << endl;
  }
  //cout << s;

  return 0;
}
