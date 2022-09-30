package main

import "fmt" 

func main() {
    for n:=0.00 ; n<11.00;n++{
    fmt.Printf( "\x1bc\x1b[42;30m\n %8.2f\n",scale0_10(n*10.00))
    }
}

func scale0_10( a float64)float64{
   aa:=a
   if a<0.00 {
       aa=0.00
   }
   if a>100.00{
       aa=100.00
   }
   aa=aa/10.00
   return aa
}
