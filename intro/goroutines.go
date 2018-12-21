package intro

import (
    "fmt"
    "time"
)

func SomeGoRoutine(i int){
    fmt.Printf("Routine %d\n", i)
}

func MainGoRoutines(){
    for i:=0; i< 10; i++{
        SomeGoRoutine(i)
    }
    /**
    This is just a simple time slee, there are no wait groups
    **/
    time.Sleep(1 * time.Second)
}