package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ClientBase() {
	//open connection:
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dial:", err.Error())
		return
	}

	//inputReader := bufio.NewReader(os.Stdin)

	//send info to server until Quit
	for {
		//content, _ := inputReader.ReadString('\n')
		inputContent := <-ReadStdin()
		if inputContent == "Q" {
			conn.Close()
			return
		}

		_, err := conn.Write([]byte(inputContent))
		if err != nil {
			fmt.Println("Error Write:", err.Error())
			return
		}
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Receive data from service:", string(buf[:length]))
	}
}

func ReadStdin() chan string {
	cb := make(chan string)
	sc := bufio.NewScanner(os.Stdin)
	go func() {

		if sc.Scan() {
			cb <- sc.Text()
		}
	}()

	return cb
}

func main() {
	ClientBase()
}
