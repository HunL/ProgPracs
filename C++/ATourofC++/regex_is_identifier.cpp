#include<iostream>
#include<regex>
#include<sstream>

using namespace std;

bool is_identifier(const string& s) {
    // regex pat {"R[_[:alpha:]\\w]"};
    regex pat {R"([_[:alpha:]]\w*)"};
    return regex_match(s,pat);
}

int main() {
    // bool res;
    // res = is_identifier("_n6");
    bool res = is_identifier("6n");
    cout << "true: " << boolalpha << true << endl;
    cout << "false: " << false << endl;
    cout << "res: " << noboolalpha << res << endl;
}