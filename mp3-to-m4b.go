package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	audioBook := os.Args[1]
	fmt.Println(audioBook)

	audioBookName := filepath.Base(audioBook)

	files, err := ioutil.ReadDir(audioBook)
	if err != nil {
		fmt.Println(err)
	}

	var parts []string
	for _, f := range files {
		if strings.ToLower(filepath.Ext(f.Name())) != ".mp3" {
			continue
		}

		parts = append(parts, audioBook+f.Name())
	}

	joinCmd := exec.Command("ffmpeg", "-i", "concat:"+strings.Join(parts, "|"), "-acodec", "copy", audioBookName+".mp3")
	joinCmd.Stdout = os.Stdout
	joinCmd.Stderr = os.Stderr

	err = joinCmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	convertCmd := exec.Command("ffmpeg", "-i", audioBookName+".mp3", "-c:a", "aac", audioBookName+".m4b")
	convertCmd.Stdout = os.Stdout
	convertCmd.Stderr = os.Stderr

	err = convertCmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(audioBookName + ".mp3")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(audioBookName + ".mp3")
	if err != nil {
		fmt.Println(err)
	}
}
