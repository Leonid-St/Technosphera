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
	cmd *exec.Cmd) (output chan bool, buffer []byte) {
	output = make(chan bool)
	go func() {
		for {
			select {
			case <-inputValueStart:
				b, err := os.ReadFile(pathTest)
				if err != nil {
					log.Print(err)
				}
				fmt.Println(string(b))
				buffer = make([]byte, len(b))
				for i := range b {
					buffer[i] = b[i]
				}
				// time.Sleep(time.Second * 2)
				// fmt.Println(cmd.Stdout.Write(b))
				output <- true
			case <-die:
				fmt.Println("read done")
				wg.Done()
				return
			}
			//fmt.Println(string(dataStdIn[:n]))

			//stdin.Write(dataStdIn[:n])
			//fmt.Println(string(dataStdIn[:n]))
		}

	}()
	return output, buffer
}

func uploadValue(
	wg *sync.WaitGroup,
	die chan bool,
	uploadValueStart chan bool,
	inputValueBuffer *[]byte,
	cmd *exec.Cmd) <-chan bool {
	output := make(chan bool)
	go func() {
		for {
			select {
			case <-uploadValueStart:
				a := []byte("Hello from test program:- value1")
				*inputValueBuffer = append(*inputValueBuffer, a...)
				//cmd.Stdout.Write(*inputValueBuffer)
			case <-die:
				fmt.Println("uploadInputBuffer done")
				wg.Done()
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
		for {
			select {
			case <-downloadValueStart:
				n, err := cmd.Stdin.Read(*dataStdOut)
				if err != nil {
					fmt.Println(err.Error())
					fmt.Println("n:", n)
					fmt.Println("dataStdOut", string((*dataStdOut)[:n]))
				}
			case <-die:
				fmt.Println("download done")
				wg.Done()
				return
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

	readDone, buffer := GetInputValue(&wg, readInputValueStart, die1, pathTest, cmd)

	readInputValueStart <- true

Loop1:
	for {
		select {
		case <-readDone:
			time.Sleep(time.Second * 1)
			die1 <- true
			break Loop1
		}
	}

	fmt.Println("Loop1 is break  - inputValueBuffer - ", string(buffer))

	uploadInputBufferDone := uploadValue(&wg, uploadValueStart, die1, &inputValueBuffer, cmd)

	uploadValueStart <- true

Loop2:
	for {
		select {
		case <-uploadInputBufferDone:
			time.Sleep(time.Second * 1)
			die2 <- true
			break Loop2
		}
	}
	fmt.Println("Loop2 is break  ")

	downloadValueBufferDone := downloadValue(&wg, downloadValueStart, die3, &dataStdOut, cmd)

	downloadValueStart <- true
Loop3:
	for {
		select {
		case <-downloadValueBufferDone:
			die3 <- true
			break Loop3
		}
	}

	fmt.Println("Loop3 is break, dataStdOut -:", string(dataStdOut))

	wg.Wait()

	fmt.Println("dataStdOut:", string(dataStdOut))

}
