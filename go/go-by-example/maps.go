package main

import "fmt"

func main(){
  mint := make(map[string]int)

  mint["first"] = 1
  mint["second"] = 2

  fmt.Println("Mint: ", mint)
  fmt.Println("First: ", mint["first"])
  fmt.Println("Second: ", mint["second"])

  mstring["first_name"] := "Guilherme"
  mstring["last_name"] = "Guilherme"
  fmt.Println("Mstring lenght: ", len(mstring))
  fmt.Println("Full name: ", mstring["first_name"], mstring["last_name"])

}
