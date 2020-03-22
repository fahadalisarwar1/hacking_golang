package handleConnection

import (
	"net"
)

func ConnectWithVictim(IP string, port string )(connection net.Conn, err error){
	LocalAddressPort := IP + ":" + port
	listener, err := net.Listen("tcp", LocalAddressPort)
	if err != nil{
		return
	}
	connection, err = listener.Accept()
	if err != nil {
		return
	}
	return
}