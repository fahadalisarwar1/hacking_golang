package ExecuteCommandWindows

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Command struct{
	CmdOutput string
	CmdError string
}

func ExecuteCommandRemotelyWindows(connection net.Conn)(err error){

	// send command from server
	// execute command remotely
	// receive back results or error
	// dir pwd date
	// stop
	// special condition will be "stop"


	reader := bufio.NewReader(os.Stdin)

	commandloop := true

	for commandloop{

		fmt.Print(">> ")
		command, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(command))
		if command == "stop\n"{
			commandloop = false
			continue
		}

		cmdStruct := &Command{}

		decoder := gob.NewDecoder(connection)
		err = decoder.Decode(cmdStruct)
		if err != nil{
			fmt.Println(err)
			continue
		}

		fmt.Println(cmdStruct.CmdOutput)
		if cmdStruct.CmdError != ""{
			fmt.Println(cmdStruct.CmdError)
		}

	}
	return
}
