#include<iostream>
#include<algorithm>
#include<vector>
#include<list>

using namespace std;

template<typename T>
using Iterator = typename T::iterator;

template<typename C, typename V>
vector<Iterator<C>> find_all(C& c, V v) {
    vector<Iterator<C>> res;

    // for (auto& p:c) {
    // for (auto& p=c.begin();p!=c.end();++p){
    for (auto p=c.begin();p!=c.end();++p){
        if(*p==v)
            res.push_back(p);
    }        
    return res;
}

int main() {
    list<double> ld {1.1,2.2,3.3};
    vector<Iterator<list<double>>> resld = find_all(ld, 2.2);
    for (auto& x: resld) {
        cout << *x << endl;
        if(*x!=2.2) 
            cerr << "ld bug\n";
    }
    cout << "list<double> no bug\n";

    string s {"this is a test string."};
    vector<Iterator<string>> resstr = find_all(s, 'i');
    for (auto& x : resstr) {
        cout << *x << endl;
        if(*x!='i')
            cerr << "string bug\n";
    }
    cout << "string no bug\n";

    vector<string> vs {"red", "green", "yellow", "red"};
    vector<Iterator<vector<string>>> resvs = find_all(vs, "red");
    for (auto& x : resvs) {
        cout << *x << endl;
        if(*x!="red")
            cerr << "vector string bug\n";
    }
    cout << "vector string no bug\n";

}