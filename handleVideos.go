package main

import (
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
	"log"
	"path/filepath"
	"strings"
)

var pathVideos string = "/home/luchifer120/projects/vodtube/server/videos/"

var outFileName string = "output.m3u8"

func ConvertVideo(inputFileName string) {

	nameFile := strings.Split(inputFileName, "\n")[0]
	pathOutput := strings.Split(inputFileName, ".")[0]
	folderOutput := strings.Split(pathOutput, "/")

	getFolder := len(folderOutput) - 1

	output := filepath.Join(pathVideos, folderOutput[getFolder], outFileName)

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(nameFile, output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("convert video...")

	trans.MediaFile().SetOutputFormat("hls")

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Printf("convert video... %.3f%%\r", msg.Progress)
	}

	fmt.Println("Finish convert Video...")

	err = <-done
}

/*
* * take image of video
 */

var outputImgPath string = "/home/luchifer120/projects/vodtube/server/images/"

func TakeScreenOfVideo(inputFileName string) {

	nameFile := strings.Split(inputFileName, "\n")[0]
	pathVideo := strings.Split(nameFile, ".")[0]
	listPaths := strings.Split(pathVideo, "/")

	getName := len(listPaths) - 1

	nameImg := listPaths[getName] + ".jpg"

	output := filepath.Join(outputImgPath, nameImg)

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(nameFile, output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("take screenshot to video...")

	trans.MediaFile().SetSeekTimeInput("00:02:30")
	trans.MediaFile().SetVframes(1)

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println(msg)
	}

	fmt.Println("Finish taked img of video...")

	err = <-done
}
