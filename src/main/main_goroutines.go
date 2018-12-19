package main

import (
    "time"
    G "goconcurrency/goroutines"
)

func main(){
    for i:=0; i< 10; i++{
        go G.SomeGoRoutine(i)
    }
    time.Sleep(1 * time.Second)
}
