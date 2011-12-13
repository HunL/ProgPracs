#include<iostream>
#include<cstdlib>

#define MAX 100

using namespace std;
  
int t;
int i,j,k;
//int size = 10;
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
  //initialize the array
  for(t = 0;t < size;t++)
  {
    cin>>a[t];
  }

  //display the original array
//  cout << "original array is :\n";
//  for(t = 0;t < size;t++)
//  {
//    cout << a[t] << " ";
//  } 
//  cout << endl;

  //quick sort
//  cout << "\nQuick sorting...\n" << endl;
  quicksort(0,size-1);

  //display the sorted array
//  cout << "sorted array is :\n";
//  for(t = 0;t < size;t++)
//  { 
//    cout << a[t] << " ";
//  }
//  cout << endl;

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
  for(t = 0;t < size;t++)
    cout << a[t] << " ";
  cout << endl;
  return 0;
}
