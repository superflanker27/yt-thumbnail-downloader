package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var videoLink string

func Ui() {
	color.Magenta("-----------------------------------")
	color.Magenta("  YT video thumbnail downloader")
	color.Magenta("-----------------------------------")
	fmt.Println()
	fmt.Println("Example URL -> https://www.youtube.com/watch?v=VIkIssdfnAC")
	fmt.Print(color.MagentaString("Enter url of video ->"), " ")
	fmt.Scan(&videoLink)
}

func GetThumbnail() {
	videoID := strings.ReplaceAll(videoLink, "https://www.youtube.com/watch?v=", "")
	getThumb := "https://i.ytimg.com/vi/" + videoID + "/hqdefault.jpg"
	fagart := exec.Command("powershell.exe", "wget", getThumb, "-OutFile", "thumb.jpg")
	fagart.Run()
}

func main() {
	Ui()
	GetThumbnail()
}
