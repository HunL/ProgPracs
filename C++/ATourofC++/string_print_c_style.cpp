#include<string>
#include<iostream>

using namespace std;

void print(const string& s) {
    printf("For people who like printf: %s\n", s.c_str());
    cout << "For people who like streams: " << s << '\n';
}

int main() {
    print("this is a string for test.");
}