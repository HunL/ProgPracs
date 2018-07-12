#include<sstream>
#include<iostream>

using namespace std;

void test() {
    ostringstream oss;
    // string str;
    // oss << "hello: " << str << endl;
    oss << "{temperature," << scientific << 123.4567890 << "}";
    cout << oss.str() << endl;
}

int main() {
    test();
}