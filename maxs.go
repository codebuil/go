package main

import "fmt" 

func main() {

    fmt.Printf( "\x1bc\x1b[42;30m\n %d\n",maxs(40,20))

}

func maxs( a int, b int) int{
    if (a>b){
        
    return a
    }
    return b
}
