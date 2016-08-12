/*
 * =====================================================================================
 *
 *       Filename:  binAdd.cpp
 *
 *    Description:  二进制数相加。
 *                  sicily的题目用来练习C++。嗯。
 *
 *        Version:  1.0
 *        Created:  07/28/2011 04:46:56 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Jiehui Li, ljiehui0826@gmail.com
 *        Company:  
 *
 * =====================================================================================
 */

#include<iostream>
#include<string>

using namespace std;

/*清除结果前面的0*/
string zeroClean(string a)
{
	if(a == "") return a = "0";
	int zero = 0;
	for(unsigned int i = 0;i < a.size();i++)
	{
		if(a[i] == '0')
			zero++;
		else break;
	}

	a = a.substr(zero,a.size()-zero);

	return a;
}

string binAdd(string a ,string b)
{
	unsigned int i;
	while(a.size() < b.size()) 
	{
		a = '0' + a;
	}

	while(b.size() < a.size()) 
	{
		b = '0' + b;
	}

	a = '0' + a;
	b = '0' + b;

	for(i = a.size()-1;i > 0;i--)
	{
		a[i] = a[i] + b[i] - '0';
/*		if(a[i] > 2)
		{
			a[i] = a[i] + b[i] -'1';
			a[i-1] = a[i-1] + b[i-1] + '1';
		}
*/
    if(a[i] - '0' > 1)
    {
      a[i] = a[i] - 2;
      a[i-1] = a[i-1] + 1;
    }
	}
	a = zeroClean(a);

	return a;
}

int main()
{
	int k;
  int c = 1;
  cin>>k;
  string str1,str2;
//  while(K > 0)
  while(k)
  {
//    cout<<c++<<endl;
    cout<<"请输入两个二进制字符串："<<endl;
  	cin>>str1>>str2;
	  cout<<c++<<" 字符串"<<str1<<"和"<<str2<<"之和为："<<binAdd(str1,str2)<<endl<<endl;
    k--;
  }

	return 0;
}
