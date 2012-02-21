#include<iostream>

#define MAX 100

using namespace std;

int size;
int i,j,t;
int a[MAX];

int partition(int first,int last)
{
    int mid;
    int pivot = a[first];
    while(first < last)
    {
        while(first < last && a[last] >= pivot)
            last--;
        if(first < last)
            a[first++] = a[last];
        while(first < last && a[first] <= pivot)
            first++;
        if(first < last)
            a[last--] = a[first];
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
    cin >> size;
    for(t = 0;t < size;t++)
    {
        cin >> a[t];
    }
    quicksort(0,size-1);
    for(i = 0;i < size-1;i++)
    {
//        for(j = size-1;j > i;j--)
//        {
            if(a[i] == a[i+1])
            {
//                for(k = j+1;k < size;k++)
//                {
                    int tmp;
                    tmp = a[i+1];
                    a[i] = tmp;
//                }
                size--;
            }
//        }
    }
    
    if(a[size] == a[size-1])
    {
        int tmp;
        tmp = a[size];
        a[size-1] = tmp;
        size--;
    }
    
    cout << size << endl;
    for(i = 0;i < size;i++)
    {
        cout << a[i] << " ";
    }
    cout << endl;
    
    return 0;
}
