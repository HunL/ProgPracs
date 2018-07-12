#include<iostream>
#include<vector>
// #include<iterator>

using namespace std;

vector<int> read_ints(istream& is) {
    vector<int> res;
    int i;
    while(is >> i)
        res.push_back(i);
    return res;
}

int main() {
    vector<int> res = read_ints(cin);
    const string s = "main res: ";
    cout << s;
    // cout << res << endl;
    // copy (res.begin(), res.end(), ostream_iterator<int> (cout, "\n"));
    for (auto v: res){
        cout << v << ' ';
    }
    cout << endl;
}