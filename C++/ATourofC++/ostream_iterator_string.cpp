#include<iostream>
#include<iterator>

using namespace std;

ostream_iterator<string> oo {cout};

int main() {
    *oo = "hello,";
    ++oo;
    *oo = "world!\n";
}