package main

import "fmt"
import "math"
import "time"

func evenlyDivisible(counter int, divisor int) bool {
  return math.Remainder(float64(counter), float64(divisor)) == 0
}

func fizzy(counter int) bool {
  return evenlyDivisible(counter, 3)
}

func buzzy(counter int) bool {
  return evenlyDivisible(counter, 5)
}

func fizzyAndBuzzy(counter int) bool {
  return fizzy(counter) && buzzy(counter)
}

func whatToSay(counter int) string {
  if fizzyAndBuzzy(counter) {
    return "Fizz Buzz!"
  } else if fizzy(counter) {
    return "Fizz"
  } else if buzzy(counter) {
    return "Buzz"
  }
  return fmt.Sprintf("%d", counter)
}

func main() {

  for counter := 1 ; true ; counter++ {
    fmt.Printf("%v\n", whatToSay(counter))

    time.Sleep(time.Second)
  }

}
