#include<iostream>
#include<vector>
#include<list>
#include<algorithm>

using namespace std;

struct Entry {
    string name;
    int number;
};

vector<Entry> phone_book = {
    {"David Hume", 123456},
    {"Karl Popper", 234567},
    {"Bertrand Arthur William Russell", 345678}
};

// 重定义<
bool operator<(const Entry& x, const Entry& y) {
    return x.name < y.name;
}

// 重定义==
bool operator==(const Entry& x, const Entry& y) {
    return x.name == y.name;
}

list<Entry> f(vector<Entry>& v) {
    list<Entry> res;
    sort(v.begin(),v.end());
    unique_copy(v.begin(),v.end(),back_inserter(res));
    return res;
}

int main() {
    list<Entry> lst = f(phone_book);
    for (const auto& x: lst) {
        cout << x.name << x.number << endl;
    }
}