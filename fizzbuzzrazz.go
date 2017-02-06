package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "bytes"
  "strings"
)

var conditions = map[int]string { 3: "Fizz", 5: "Buzz", 7: "Razz" }
var buffer bytes.Buffer

func fizzbuzzrazz(number int) string {
  buffer.Reset()
  for key, value := range conditions {
    if number % key == 0 {
      buffer.WriteString(value)
    }
  }

  if buffer.String() == "" {
    return strconv.Itoa(number) + " is not divisible by any of the options!"
  } else {
    return buffer.String()
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Welcome to the FizzBuzz challenge! Please, enter a number: ")
  text, _ := reader.ReadString('\n')
  number, err := strconv.Atoi(strings.TrimSpace(text))
  if err != nil {
    fmt.Println("This is not a number...")
  } else {
    fmt.Println(fizzbuzzrazz(number))
  }
  main()
}
