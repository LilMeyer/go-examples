package examples

import (
  "fmt"
)

func ExampleDefer() {
  fmt.Println("start")
  defer fmt.Println("defered")
  fmt.Println("end")
  // Output: start
  // end
  // defered
}
