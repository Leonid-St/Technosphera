package main

import (
	"bufio"
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
	cmd *exec.Cmd) (buffer chan []byte) {
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
					fmt.Println("fmt.Println(string(b))", string(b))
					buffer <- b
				}
			case <-die:
				fmt.Println("read done")
				wg.Done()
				close(buffer)
				close(die)
				return
			}

		}

		//fmt.Println(string(dataStdIn[:n]))

		//stdin.Write(dataStdIn[:n])
		//fmt.Println(string(dataStdIn[:n]))

	}()
	return buffer
}

func uploadValue(
	wg *sync.WaitGroup,
	uploadValueStart chan bool,
	die chan bool,
	inputValueBuffer []byte,
	cmd *exec.Cmd) <-chan string {
	output := make(chan string)
	go func() {
		for {
			select {
			case msg := <-uploadValueStart:
				fmt.Println("msg := <-uploadValueStart , msg := ", msg)
				if msg {
					// a := []byte("Hello from test program:- value1")
					// *inputValueBuffer = append(*inputValueBuffer, a...)
					var HELLO = inputValueBuffer
					fmt.Println("HELLO", string(HELLO))
					stdin, err := cmd.StdinPipe()
					if nil != err {
						log.Fatalf("Error obtaining stdin: %s", err.Error())
					}
					n, err := stdin.Write(inputValueBuffer)
					if nil != err {
						log.Fatalf("Error writing to stdin: %s", err.Error())
					}
					fmt.Println("upload,:", n)
					output <- "uploaded file"
				}
			case <-die:
				fmt.Println("uploadInputBuffer done")
				wg.Done()
				close(output)
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
	cmd *exec.Cmd) <-chan string {
	output := make(chan string)

	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)

	go func() {
		for {
			select {
			case msg := <-downloadValueStart:
				fmt.Println("msg := <-downloadValueStart, msg := ", msg)
				if msg {
					var stdOut []byte

					for scanner.Scan() {
						fmt.Println("scanner.Text():-",scanner.Text())
						fmt.Println("scanner.Bytes()", scanner.Bytes())

						res := make([]byte, len(stdOut)+len(scanner.Bytes()))

						res = append(stdOut, scanner.Bytes()...)

						stdOut = make([]byte, len(stdOut)+len(scanner.Bytes()))

						copy(stdOut, res)
					}
					copy(*dataStdOut, stdOut)
					output <- "downloaded "
				}
			case <-die:
				fmt.Println("downloadValue done")
				wg.Done()
				close(output)
				close(die)
				return
			}
		}

	}()
	return output
}

func Reader(pathProgram string, pathTest string) {
	var wg sync.WaitGroup

	cmd := exec.Command(`go`, "run", "pathProgram")

	var dataStdOut *[]byte

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
	time.Sleep(time.Second * 1)
	bytes := <-buffer
	fmt.Println(bytes, " - ", string(bytes))
	b = make([]byte, len(bytes))
	copy(b, bytes)
	fmt.Println(b, " - ", string(b))
	die1 <- true

	fmt.Println("  - buffer from file- ", string(b))

	uploadInputBufferDone := uploadValue(&wg, uploadValueStart, die2, b, cmd)

	uploadValueStart <- true

	msg := <-uploadInputBufferDone

	fmt.Println(msg)

	//time.Sleep(time.Second * 1)

	die2 <- true

	fmt.Println("start downloadValue")
	downloadValueBufferDone := downloadValue(&wg, downloadValueStart, die3, dataStdOut, cmd)

	downloadValueStart <- true

	fmt.Println(<-downloadValueBufferDone)

	die3 <- true

	fmt.Println(" dataStdOut -:", string(*dataStdOut))

	wg.Wait()

	fmt.Println("dataStdOut:", string(*dataStdOut))

}
