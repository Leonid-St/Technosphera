package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func GetInputValue(
	wg *sync.WaitGroup,
	inputValueStart chan bool,
	die chan bool,
	pathTest string,
	cmd *exec.Cmd) ( buffer chan []byte) {
	buffer = make(chan []byte)
	go func() {
		for {
			select {

				case msg := <-inputValueStart:
				fmt.Println("msg := <-inputValueStart , msg := ", msg)
				if msg {

				b, err := os.ReadFile(pathTest)
				if err != nil {
					log.Print(err)
				}
				fmt.Println("fmt.Println(string(b))",string(b))
				buffer <- b
				}
				case <-die:
				fmt.Println("read done")
				wg.Done()
				close(die)
				return 	
			}
			
		}
			
				
			//fmt.Println(string(dataStdIn[:n]))

			//stdin.Write(dataStdIn[:n])
			//fmt.Println(string(dataStdIn[:n]))

	}()
	return  buffer
}

func uploadValue(
	wg *sync.WaitGroup,
	uploadValueStart chan bool,
	die chan bool,
	inputValueBuffer *[]byte,
	cmd *exec.Cmd) <-chan string {
	output := make(chan string)
	go func() {
		for {
			select {
			case msg :=   <-uploadValueStart:
	fmt.Println("msg := <-uploadValueStart , msg := ", msg)
	if msg {
		a := []byte("Hello from test program:- value1")
				*inputValueBuffer = append(*inputValueBuffer, a...)
				// n,err := cmd.Stdout.Write(*inputValueBuffer)
				// if err!= nil {
            //    fmt.Println(err)
            // }
				// fmt.Println("upload,:",n)
				output <- "uploaded file"
	}
	case <-die:
		fmt.Println("uploadInputBuffer done")
				wg.Done()
				close(die)
				return
			}
		}
	}()
	return output
}

func downloadValue(
	wg *sync.WaitGroup,
	downloadValueStart chan bool,
	die chan bool,
	dataStdOut *[]byte,
	cmd *exec.Cmd) <-chan bool {
	output := make(chan bool)
	go func() {
		for
		{
msg:= <-downloadValueStart 
	fmt.Println("msg := <-downloadValueStart , msg := ", msg)
	if msg {
			n, err := cmd.Stdin.Read(*dataStdOut)
				if err != nil {
					fmt.Println(err.Error())
					fmt.Println("n:", n)
					fmt.Println("dataStdOut", string((*dataStdOut)[:n]))	
					fmt.Println("download done")
				wg.Done()
				return
	}
		}
		}
	

	}()
	return output
}

func Reader(pathProgram string, pathTest string) {
	var wg sync.WaitGroup

	cmd := exec.Command(`go`, "run", "pathProgram")

	inputValueBuffer := make([]byte, 2048)

	dataStdOut := make([]byte, 2048)

	wg.Add(3)

	readInputValueStart := make(chan bool)
	uploadValueStart := make(chan bool)
	downloadValueStart := make(chan bool)

	die1 := make(chan bool)
	die2 := make(chan bool)
	die3 := make(chan bool)

	buffer := GetInputValue(&wg, readInputValueStart, die1, pathTest, cmd)
	readInputValueStart <- true
	var b []byte 
//Loop1:
	// for {
	// 	select {
	// 	case <-readDone:
			time.Sleep(time.Second * 1)
			bytes := <- buffer
			fmt.Println(bytes," - ",string(bytes))
			b = make([]byte, len(bytes))
			copy(b,bytes)
			fmt.Println(b," - ",string(b))
			die1 <- true

			//break Loop1
	// 	}
	// }

	fmt.Println("  - buffer from file- ", string(b))

	uploadInputBufferDone := uploadValue(&wg, uploadValueStart, die2, &inputValueBuffer, cmd)

	uploadValueStart <- true


		 msg:= <-uploadInputBufferDone

			fmt.Println("153",msg)
	fmt.Println("  inputValueBuffer: ",&inputValueBuffer)
	time.Sleep(time.Second * 1)
	die2 <- true
	
	downloadValueBufferDone := downloadValue(&wg, downloadValueStart, die3, &dataStdOut, cmd)

	downloadValueStart <- true
	fmt.Println("fmt.Println( )<-downloadValueBufferDone", <-downloadValueBufferDone)
			die3 <- true
			

	fmt.Println(" dataStdOut -:", string(dataStdOut))

	wg.Wait()

	fmt.Println("dataStdOut:", string(dataStdOut))

}
