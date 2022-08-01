package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Youtube downloader START")
	Download("https://www.youtube.com/watch?v=UaUa_0qPPgc")
}
func Download(url string) {
	out, e := exec.Command(".\\yt-dlp", "-f", "ba", "-x", "--audio-format", "mp3", url, "-o", "%(title)s.%(ext)s").Output() //yt-dlp -f 'ba' -x --audio-format mp3 $2 -o '%(title)s.%(ext)s' ORIGINAL COMMAND
	CheckError(e)
	fmt.Printf("%s", out)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
