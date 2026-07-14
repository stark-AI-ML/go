#include <iostream>

int printNRev(int x) // rec
{
    if (x == 0)
        return x;

    std::cout << x << std::endl;

    return printNRev(x - 1);
}

int printN(int x)
{
    if (x == 0)
        return x;

    printN(x - 1);

    std::cout << x << std::endl;
    return x;
}

int main()
{

    // std::cout << "hello";

    // printNRev(10);
    std::cout << printN(10);
}