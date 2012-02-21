#include <map>
#include <string>
#include <iostream>
using namespace std;
 
int main()
{
    int n;
 
    while(cin>>n && (n!=0))
    {
        map<string, int> balloons;
        map<string, int>::iterator it;
 
        for(int i=0; i<n; i++)
        {
            string color;
            cin>>color;
 
            it = balloons.find(color);
            if(it == balloons.end())
            {
                balloons.insert(map<string, int>::value_type(color,1));
            }
            else
            {
                (it->second)++;
            }
        }
 
        it = balloons.begin();
        string maxcolor = it->first;
        int maxnum = it->second;
 
        for(++it; it != balloons.end(); ++it)
        {
            if(it->second > maxnum)
            {
                maxnum = it->second;
                maxcolor = it->first;
            }
        }
 
        cout<<maxcolor<<endl;
    }
 
    return 0;
}
