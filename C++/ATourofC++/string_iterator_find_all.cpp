#include<string>
#include<vector>
#include<iostream>

using namespace std;

vector<string::iterator> find_all(string& s, char& c) {
    vector<string::iterator> res;
    // for (char* p = s.begin();p!=end();++p) {
    for (auto p = s.begin();p!=s.end();++p) {
        if (*p==c)
            res.push_back(p);
    }
    return res;
}

int main() {
    string s = "this is a test string.";
    char c = 'i';
    vector<string::iterator> vec = find_all(s, c);
    for (const auto& p: vec) {
        cout << *p << endl;
        if (*p!=c)
            cerr << "a bug\n";
    }
}