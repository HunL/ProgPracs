#include <iostream>

using namespace std;

void fore_inc() {
    int v1[10] = {0,1,2,3,4,5,6,7,8,9};
    int v2[10];

    for (auto int i=0;i!=10;++i)
        v2[i]=v1[i];

    for (auto int i=0;i!=10;++i)
        cout << "v1: " << v1[i] << ", v2: " << v2[i] << endl;

    cout << endl;
}

void back_inc() {
    int v1[10] = {0,1,2,3,4,5,6,7,8,9};
    int v2[10];

    for (auto int i=0;i!=10;i++)
        v2[i]=v1[i];

    for (auto int i=0;i!=10;i++)
        cout << "back v1: " << v1[i] << ", v2: " << v2[i] << endl;

    cout << endl;
}

int main() {
    //fore_inc();
    back_inc();
}
