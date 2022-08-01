package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Internet is shit, have fun waiting")
}
func Download(url string) {
	cmd := exec.Command(".\\yt-dlp.exe")

	e := cmd.Run()
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
