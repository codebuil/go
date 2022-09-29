package main

import "fmt" 

func main() {
    for n:=0 ; n<16;n++{
    fmt.Printf( "\x1bc\x1b[42;30m\n %d\n",power(2,n))
    }
}

func power( a int, b int) int{
   aa:=a
   if b==0{
       return 1
   }
   if b==1{
       return a
   }
   for n:=1 ; n<b;n++{
       aa=aa*a
   }
   return aa
}
