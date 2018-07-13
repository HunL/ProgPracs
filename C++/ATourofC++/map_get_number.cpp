#include<map>
#include<iostream>

using namespace std;

// struct Entry {
//     string name;
//     int number;
// };

// map<Entry> phone_book = {
map<string, int> phone_book = {
    {"David Hume", 123456},
    {"Karl Popper", 234567},
    {"Bertrand Arthur William Russell", 345678}
};

int get_number(const string& s) {
    // return phone_book.at(s);
    // return phone_book[s];
    map<string, int>::iterator key = phone_book.find(s);
    if (key!=phone_book.end())
    {
        cout << "key: " << key->first << key->second << endl;
        return key->second;
    }
    return 0;
}

int main() {
    for (const auto& x : phone_book) {
        cout << "1. first: " << x.first << " second: " << x.second << endl;
    }
    // int res = get_number("Karl Popper");
    int res = get_number("Karl");
    cout << res << endl;

    for (const auto& x : phone_book) {
        cout << "2. first: " << x.first << " second: " << x.second << endl;
    }
}