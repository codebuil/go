package main

import "fmt" 

func main() {

    fmt.Printf( "\x1bc\x1b[42;30m\n %d\n",mins(40,20))

}

func mins( a int, b int) int{
    if (a>b){
        
    return a
    }
    return b
}
