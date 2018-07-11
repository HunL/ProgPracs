#include<iostream>
#include<regex>
#include<fstream>

using namespace std;

void use() {
    ifstream in("file.txt");
    if(!in)
        cerr << "no file\n";

    cout << "aaaaaaa" << endl;
    regex pat {R"(\w{2}\s*\d{5}(-\d{4})?)"};
    // regex pat (R"(\w{2}\s*\d{5}(-\d{4})?)");
    cout << "bbbbbbb" << endl;

    int lineno = 0;
    for (string line;getline(in, line);){
        cout << "ccccccc" << endl;
        ++lineno;
        smatch matches;
        if (regex_search(line, matches, pat)){
            cout << lineno << ":" << matches[0] << '\n';
            cout << "matches size: " << matches.size() << endl;
            if (1 < matches.size() && matches[1].matched)
                cout << "\t: " << matches[1] << '\n';
        }
    }
}

int main() {
    use();
}