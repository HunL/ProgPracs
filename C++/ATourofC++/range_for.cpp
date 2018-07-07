#include <iostream>

using namespace std;

int main() {
    int v[] = {0,1,2,3,4,5,6,7,8,9};

    for (auto x:v)
        cout << x << endl;

    for (auto x:{10,21,32,43,54,65})
        cout << x << endl;
}
