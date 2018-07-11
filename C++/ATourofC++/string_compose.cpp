#include<string>
#include<iostream>

using namespace std;

string compose(const string& name, const string& domain) {
    return name + '@' + domain;
}

int main() {
    auto addr = compose("dmr", "bell-labs.com");
    cout << addr << endl;
}
