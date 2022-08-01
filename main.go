package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Youtube downloader START")
	Download("https://www.youtube.com/watch?v=UaUa_0qPPgc", "test")
}
func Download(url string, foldername string) {
	out, e := exec.Command("cmd.exe", "/c", "mkdir", foldername).Output() //why windows just why do you have to make things harder
	CheckError(e)
	fmt.Printf("%s", out)
	cmd := exec.Command("..\\yt-dlp", "-f", "ba", "-x", "--audio-format", "mp3", url, "-o", "%(title)s.%(ext)s") //yt-dlp -f 'ba' -x --audio-format mp3 $2 -o '%(title)s.%(ext)s' ORIGINAL COMMAND
	cmd.Dir = ".\\" + foldername
	e = cmd.Run()
	CheckError(e)
	fmt.Printf("%s", out)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}
