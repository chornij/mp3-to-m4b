package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	audioBook := os.Args[1]
	log.Printf("Audiobook: %s\n\n", audioBook)

	audioBookName := filepath.Base(audioBook)

	parts, err := getParts(audioBook, "files")
	if err != nil {
		log.Fatalln("Failed to get audiobook parts", err)
	}

	partsDirs, err := getParts(audioBook, "dirs")
	if err != nil {
		log.Fatalln("Failed to get audiobook parts", err)
	} else {
		parts = append(parts, partsDirs...)
	}

	log.Println("Files order:")
	for _, p := range parts {
		fmt.Println(p)
	}

	joinCmd := exec.Command("ffmpeg", "-i", "concat:"+strings.Join(parts, "|"), "-acodec", "copy", audioBookName+".mp3")
	joinCmd.Stdout = os.Stdout
	joinCmd.Stderr = os.Stderr

	err = joinCmd.Run()

	defer func() {
		err = os.Remove(audioBookName + ".mp3")
		if err != nil {
			log.Fatalln("Failed to remove mp3", err)
		}
	}()

	if err != nil {
		log.Fatalln("Failed to concat mp3", err)
	}

	convertCmd := exec.Command("ffmpeg", "-i", audioBookName+".mp3", "-vn", "-c:a", "aac", audioBookName+".m4b")
	convertCmd.Stdout = os.Stdout
	convertCmd.Stderr = os.Stderr

	err = convertCmd.Run()
	if err != nil {
		err = os.Remove(audioBookName + ".m4b")
		if err != nil {
			log.Fatalln("Failed to remove corrupt m4b", err)
		}

		log.Fatalln("Failed to convert m4b", err)
	}
}

func getParts(dir string, objType string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var parts []string
	for _, f := range files {
		if objType != "dirs" && !f.IsDir() && strings.ToLower(filepath.Ext(f.Name())) == ".mp3" {
			parts = append(parts, path.Join(dir, f.Name()))
		} else if objType == "dirs" && f.IsDir() {
			subParts, err := getParts(path.Join(dir, f.Name()), "files")
			if err != nil {
				return nil, err
			}

			subPartsDirs, err := getParts(path.Join(dir, f.Name()), "dirs")
			if err != nil {
				return nil, err
			} else {
				subParts = append(subParts, subPartsDirs...)
			}

			parts = append(parts, subParts...)
		}
	}

	return parts, nil
}
