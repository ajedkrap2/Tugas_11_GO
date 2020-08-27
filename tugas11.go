package main

import "fmt"
import "strconv"

func main() {
  defer conclude()
    var input1, input2 string
    fmt.Print("Masukan input pertama: ")
    fmt.Scanln(&input1)
    fmt.Print("Masukan input kedua: ")
    fmt.Scanln(&input2)

    var x, errx = strconv.ParseFloat(input1, 8)
    var y, erry = strconv.ParseFloat(input2, 8)

    if errx != nil || erry != nil {
      panic("Input Invalid")
    } else {
      // fmt.Printf("%v dan %v \n", input1, input2)
      calculate(x, y)
    }

}

func calculate(x float64, y float64){
  fmt.Printf("Hasil Penambahan: %v\n", x+y)
  fmt.Printf("Hasil Pengurangan: %v\n", x-y)
  fmt.Printf("Hasil Perkalian: %v\n", x*y)
  fmt.Printf("Hasil Pembagian: %v\n", x/y)
}

func conclude(){
  r := recover()
  if r != nil {
    fmt.Println(r)
  } else {
    fmt.Println("Success")
  }
}
