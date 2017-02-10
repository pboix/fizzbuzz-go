package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "bytes"
  "strings"
  "sort"
)

var conditions = map[int]string { 3: "Fizz", 5: "Buzz", 7: "Razz" }
var buffer bytes.Buffer

func fizzbuzzrazz(number int) string {
  buffer.Reset()
  var keys []int
  for k := range conditions {
    keys = append(keys, k)
  }
    sort.Ints(keys)
    for  _, k := range keys {
      if number % k == 0 {
        buffer.WriteString(conditions[k])
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
  if err != nil || number == 0 {
    fmt.Println("You better enter a valid number...")
  } else {
    fmt.Println(fizzbuzzrazz(number))
  }
  main()
}
