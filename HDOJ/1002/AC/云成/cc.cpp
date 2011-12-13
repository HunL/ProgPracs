#include<iostream>
#include<cstdlib>

#define MAX 100

using namespace std;
  
int t;
int i,j,k;
int size;
int a[MAX];

int partition(int first,int last)
{
  int mid;
  int pivot = a[first];

  while(first < last)
  {
    while(first < last && a[last] >= pivot)
      last--;//从右往左找到第一个比pivot小的数
    if(first < last)
      a[first++] = a[last];//交换
    while(first < last && a[first] <= pivot)
      first++;//从左往右找到第一个比pivot大的数
    if(first < last)
      a[last--] = a[first];//交换
  }
  
  a[first] = pivot;
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
  cin>>size;

  for(t = 0;t < size;t++)
  {
    cin>>a[t];
//   a[t] = rand()%1000;
  }
  cout << endl;

  for(i = 0;i < size;i++)
  {
    for(j = size-1;j > i;j--)
    {
      if(a[i] == a[j])
      {
        for(k = j+1;k < size;k++)
        {
          int tmp;
          tmp = a[k];
          a[k-1] = tmp;
        }
        size--;
      }
    }
  }
    cout << size << endl;

  quicksort(0,size-1);


  for(t = 0;t < size;t++)
    cout << a[t] << " ";
  cout << endl;
  return 0;
}
