#include<list>
#include<iostream>

using namespace std;

struct Entry {
    string name;
    int number;
};

list<Entry> phone_book = {
        {"David Hume", 123456},
        {"Karl Popper", 234567},
        {"Bertrand Arthur William Russell", 345678}
    };

int get_num(const string& s) {
    for (const auto& x : phone_book) 
        if (x.name == s)
            return x.number;
        return 0;
}

int main() {
    // int res = get_num("Karl Popper");
    int res = get_num("Karl");
    cout << res << endl;    
}