package main

import (
    I "github.com/jtdv01/goconcurrency/intro"
    Barry "github.com/jtdv01/goconcurrency/barycenter"
    "os"
    "fmt"
	"strconv"
)

func main(){
    args := os.Args
    if (len(args) < 2) {
        fmt.Println(fmt.Errorf("Not enough args!"))
        os.Exit(1)
    }

    task := args[1]

    /**
        Chapter 2: Introduction
        * Examples on basic go routines
        * Examples on Channels
    **/
    if task == "goRoutines" {
        I.MainGoRoutines()
    } else if task == "sharedMemory" {
        I.MainSharedMemory()
    } else if task == "channels"{
        I.MainChannels()
    } else if task == "singleChannel"{
        I.MainSingleChannel()
    } else if task == "bufferredChannels"{
        I.MainBufferedChannels()
    } else if task == "nonBlockingCakeFactory"{
        I.MainNonBlocking()
    }


    /**
        Chapter 3: Data Parallelism
    **/
    if task == "generateBarrycenter" {
	    nBodies, err := strconv.Atoi(os.Args[2])
        if err != nil {
            os.Exit(1)
        }
        Barry.NaiveBarycenterDataset(nBodies)
    }
}
