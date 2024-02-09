package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

type TestCase struct {
	input             string
	expectedOutput    string
	answerFromProgram string
}

func Reader(pathProgram string, pathInput string, pathOutput string) {
	//b, err := os.ReadFile(pathInput)
	// if err != nil {
	// 	panic(err)
	// }
	fileInput, err := os.Open(pathInput)
	if err != nil {
		panic("29" + err.Error())
	}
	//var reader *bufio.Reader
	readerInput := bufio.NewReader(fileInput)
	var input string
	var inputBuffer []byte
	for {
		oneByte, err := readerInput.ReadByte() //ReadBytes('\n')
		//fmt.Println("fmt.Println(bytes)", bytes)
		if err != nil {
			if err == io.EOF {
				//log.Print("EOF", " - ", oneByte, " - ", inputBuffer[len(inputBuffer)-1] == '\n', "\n")
				if inputBuffer[len(inputBuffer)-1] == '\n' {
					inputBuffer = inputBuffer[:len(inputBuffer)-1]
				}
			} else {
				panic(" err != nil && err != io.EOF " + err.Error())
			}
			break
		}
		// if err != nil && err != io.EOF {
		// 	panic(" err != nil && err != io.EOF " + err.Error())
		// }
		if len(inputBuffer) == 0 {
			inputBuffer = make([]byte, 1) //len(bytes))
			//opy(inputBuffer, bytes)
			// for i := 0; i < len(inputBuffer); i++ {
			// 	inputBuffer[i] = bytes[i]
			// }
			inputBuffer[0] = oneByte
		} else {
			// tmpBuf := make([]byte, len(inputBuffer)+len(bytes))
			// //copy(tmpBuf, inputBuffer)
			// for i := 0; i < len(inputBuffer); i++ {
			// 	tmpBuf[i] = inputBuffer[i]
			// }
			// inputBuffer = make([]byte, len(inputBuffer)+len(bytes))
			// //copy(inputBuffer, tmpBuf)
			// for i := 0; i < len(tmpBuf); i++ {
			// 	inputBuffer[i] = tmpBuf[i]
			// }
			inputBuffer = append(inputBuffer, oneByte)
		}
	}
	input = string(inputBuffer)
	//fmt.Println("--")
	//fmt.Print("Input file:\n", input, "\n")
	//fmt.Println("--")
	//--//
	fileOutput, err := os.Open(pathOutput)
	if err != nil {
		panic(err)
	}
	readerOutput := bufio.NewReader(fileOutput)
	var output string
	for {
		str, err := readerOutput.ReadByte()
		if err != nil {
			if err == io.EOF {
				if output[len(output)-1] == '\n' {
					output = output[:len(output)-1]
				}
				//log.Print("EOF", " - ", str, " - ", output[len(output)-1] == '\n', "\n")
			} else {
				panic(" err != nil && err != io.EOF " + err.Error())
			}
			break
		}
		// if err != nil && err != io.EOF {
		// 	panic(" err != nil && err != io.EOF " + err.Error())
		// }
		output += string(str)
	}
	// fmt.Println("--")
	// fmt.Print("Output file:\n", output, "\n")
	// fmt.Println("--")
	//--//
	cmd := exec.Command("go", "run", pathProgram)
	// var in *bufio.Reader
	//var writerProgram *bufio.Writer
	// in = bufio.NewReader(os.Stdin)
	writerToProgram, err := cmd.StdinPipe() //bufio.NewWriter(cmd.Stdout)
	if err != nil {
		panic("error cmd.StdinPipe()  " + err.Error())
	}
	readerFromProgram, err := cmd.StdoutPipe() //bufio.NewReader(cmd.Stdin)
	if err != nil {
		panic("error cmd.StdoutPipe()  " + err.Error())
	}
	err = cmd.Start()
	if err != nil {
		panic("error cmd.Start()  " + err.Error())
	}
	var programAnswer []byte = make([]byte, 16384)
	n, err := writerToProgram.Write(append(inputBuffer, '\n'))
	if err != nil {
		panic("writerToProgram.Write" + err.Error())
	}
	fmt.Println("Write ", n, "bytes")
	time.Sleep(time.Second * 2)
	//	for {
	bytes, err := readerFromProgram.Read(programAnswer) //ReadBytes('\n')
	if err != nil {
		panic("readerFromProgram.Read" + err.Error())
	}
	fmt.Println("Read", bytes, "bytes")
	// 	if err != nil {
	// 		break
	// 	}
	// 	if len(programAnswer) == 0 {
	// 		programAnswer = make([]byte, len(bytes))
	// 		for i := 0; i < len(programAnswer); i++ {
	// 			programAnswer[i] = bytes[i]
	// 		}
	// 	} else {
	// 		programAnswer = append(programAnswer, bytes...)
	// 	}
	// }
	var realAnswer = make([]byte, bytes)
	copy(realAnswer, programAnswer[:bytes])
	if string(realAnswer[:len(realAnswer)-1]) == output {
		fmt.Println("Test passed")
	} else {
		ts := TestCase{
			input:          input,
			expectedOutput: output,
		}
		LogTestCase(ts)
		fmt.Println("Test failed")
	}
	fileInput.Close()
	fileOutput.Close()
	//cmd.Wait()
}

func LogTestCase(testCase TestCase) {
	fmt.Println("--")
	fmt.Println(testCase.input)
	fmt.Println("--")
	fmt.Println(testCase.expectedOutput)
	fmt.Println("--")
	fmt.Println(testCase.answerFromProgram)
	fmt.Println("--")
}
