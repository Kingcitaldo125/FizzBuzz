package main

import(
    "fmt"
    "strconv"
)

func fizz_buzz(n int) string {
  if n % 3 == 0 && n % 5 == 0 {
    return "FizzBuzz"
  } else if n % 3 == 0 {
    return "Fizz"
  } else if n % 5 == 0 {
    return "Buzz"
  } else {
    return strconv.Itoa(n)
  }
}

func main() {
  for i := 1; i < 101; i++ {
    fmt.Println(fizz_buzz(i))
  }
}
