#include <iostream>

int main()
{
	for(unsigned int i = 1; i < 101; ++i)
	{
		cout << (i % 3 == 0 && i % 5 == 0 ? "FizzBuzz" : i % 3 == 0 ? "Fizz" : i % 5 == 0 ? "Buzz" : std::to_string(i)) << endl;
	}
	
	return 0;
}
