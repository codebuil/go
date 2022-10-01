package main 
import (
    "fmt"
    
)
var(
    m uint64=1
    n uint64=0
)
func main() {
   
       
   
   
     
          fmt.Printf( "\x1b[42;30m\n \n",) 
     for n:=0 ; n<19;n++{
         fmt.Println(m)
         fmt.Printf("h%x\n",m)
         m=mults(m,10)
          
     }
          
  } 
func mults( a uint64, b uint64)uint64{ 

        
                return a*b
                
}
