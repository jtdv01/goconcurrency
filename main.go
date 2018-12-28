package main

import (
    I "github.com/jtdv01/goconcurrency/intro"
    Barycenter "github.com/jtdv01/goconcurrency/barycenter"
    BooksHandler "github.com/jtdv01/goconcurrency/booksHandler"
    "os"
    "fmt"
    "strconv"
    Newsclient "github.com/jtdv01/goconcurrency/newsclient"
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
    } else if task == "parallelism"{
        I.MainParallelism()
    } else if task == "serialTasks"{
        I.MainSerialTasks()
    } else if task == "booksHandler"{
        BooksHandler.Main()
    }


    /**
        Chapter 3: Data Parallelism
    **/
    if task == "generateBarycenter" {
        // Create these datasets using `make barycenter_datasets`
	    nBodies, err := strconv.Atoi(os.Args[2])
        if err != nil {
            os.Exit(1)
        }
        Barycenter.GenerateBarycenterDatasets(nBodies)
    } else if task == "naiveBarycenter"{
	    filename := os.Args[2]
        fmt.Println("Reading ", filename)
        Barycenter.NaiveBarycenter(filename)
    } else if task == "parallelBarycenter"{
	    filename := os.Args[2]
        fmt.Println("Reading ", filename)
        Barycenter.ParallelBarycenter(filename)
    } else if task == "nonConcurrentNewsclient" {
        Newsclient.NonconcurrentNewsclient()
    } else if task == "concurrentNewsclient" {
        Newsclient.ConcurrentNewsclient()
    }

}
