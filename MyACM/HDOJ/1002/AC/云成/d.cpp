#include<iostream>
#include<cstdlib>
#define MAX 300

using namespace std;

int s[MAX];
int tmp[MAX];

  int partition(int first,int last)
{
  int mid;
  int pivot = s[first];

  while(first < last)
  {
    while(first < last && s[last] >= pivot)
      last--;//从右往左找到第一个比pivot小的数
    if(first < last)
      s[first++] = s[last];//交换
    while(first < last && s[first] <= pivot)
      first++;//从左往右找到第一个比pivot大的数
    if(first < last)
      s[last--] = s[first];//交换
  }
  
  s[first] = pivot;
  mid = first;
  return mid;
}

void quicksort(int first,int last)
{
  int mid;
  if(first < last)
  {
    mid = partition(first,last);
    quicksort(first,mid-1);
    quicksort(mid+1,last);
  }
}


int main()
{
  int i,j,n;
  int a,b,c;
//  scanf("%d",&n);
  cin >> n;
//  int s[n];
  for(i = 0;i < n;i++)
  {
//    scanf("%d %d %d",&a,&b,&c);
    cin >> a >> b >> c;
    s[i] = a+b+c;
    tmp[i] = s[i];
//    printf("%d ",s[i]);
  }
  
  quicksort(0,n-1);

//  i = 5;
//  while(i > 0)
  for(i = n-1;i > n-6;i--)
  {
for(j = 0;j < n;j++)
{
if(tmp[j] == s[i])
  cout << j+1; 
}
    cout << " " << s[i] << " "<< endl;
  }
  return 0;
}
