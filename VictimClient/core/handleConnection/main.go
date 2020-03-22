package handleConnection

import (
	"net"
)

func ConnectWithServer(ServerIP string, Port string)(connection net.Conn, err error){
	ServerAddress := ServerIP + ":" + Port
	connection, err = net.Dial("tcp", ServerAddress)
	if err != nil{
		return
	}

	return
}
