#include "include/face.h"
#include <iostream>
#include <cstring>

using namespace std;

int main(int argc, char **argv) {
    auto s = getFace();
    for (int i = 1; i <= 10; i++) {
        cout << "hello" << endl;
    }
    cout << s << endl;
    cout << strlen(s) << endl;
    return 0;
}
