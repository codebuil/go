package main

import "fmt" 

func main() {

    fmt.Printf( "\x1bc\x1b[42;30m\n %s\n",retstr())

}

func retstr() string{
    return "hello world"
}