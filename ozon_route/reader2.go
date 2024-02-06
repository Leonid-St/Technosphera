package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func writeDataFromFileToProgram(programPath, inputFilePath string) (*exec.Cmd, error) {
	cmd := exec.Command("go", "run", programPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return cmd, err
	}
	defer stdin.Close()

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return cmd, err
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Print("input ", input, "\n")
		_, err := fmt.Fprintln(stdin, input)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return cmd, err
		}
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return cmd, err
	}
	return cmd, nil
}

func testProgramOutputExpected(cmd *exec.Cmd, expectedOutputFilePath string) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	cmd.Start()
	b, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		log.Print(err, "\n")
	}
	var reader = bufio.NewReader(stdout)
	var buffer []string = make([]string, 1)
	var k = 0
	time.Sleep(100 * time.Millisecond)
	for {
		str, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Print("76 ", err, "\n")
		}
		fmt.Print("str ", str, "\n")
		fmt.Print("str in string :", string(str), "\n")
		if err == io.EOF {
			fmt.Print("EOF", "\n")
			break
		}
		if k == 0 {
			buffer[0] = strings.TrimSpace(string(str)) //str[:len(str)-1])
		} else {
			var tmp []string = make([]string, len(buffer))
			for i := range buffer {
				tmp[i] = buffer[i]
			}
			//copy(tmp, buffer)
			buffer = make([]string, len(buffer)+1)
			for i := range tmp {
				buffer[i] = tmp[i] // buffer[i]
			}
			//copy(buffer, tmp)
			buffer[k] = strings.TrimSpace(string(str)) //string(str[:len(str)-1])
		}
		k++
	}
	fmt.Print("buffer ", buffer, "\n")
	fmt.Print("output", "\n", strings.Join(buffer, "\n"), "\n", "err:", err, "\n")
	if err != nil && err != io.EOF {
		return err
	}

	if strings.TrimSpace(strings.Join(buffer, "\n")) == strings.TrimSpace(string(b)) {
		fmt.Print("Тест пройден!", "\n")
	} else {
		fmt.Print("Тест не пройден. Ожидаемый вывод:",
			"\n",
			strings.TrimSpace(string(b)),
			"\n", "Фактический вывод:",
			"\n",
			strings.TrimSpace(strings.Join(buffer, "\n")))
	}

	return nil
}

func GptReader(programPath string, inputFilePath string, expectedOutputFilePath string) {
	cmd, err := writeDataFromFileToProgram(programPath, inputFilePath)
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка записи данных в программу:", err)
		return
	}
	err = testProgramOutputExpected(cmd, expectedOutputFilePath)
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка тестирования программы:", err)
		return
	}
	cmd.Wait()
}
