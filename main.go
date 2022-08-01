package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Youtube downloader START")
	readURLFromFile("urls") //to lazy to implement a commandline parser just going to hardcode it
}

func readURLFromFile(fileName string) {
	readFile, err := os.Open(fileName)
	CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	//threads := 0
	//for fileScanner.Scan() {
	//	threads += 1
	//}
	//wg := new(sync.WaitGroup)
	//wg.Add(threads)
	for fileScanner.Scan() {
		ss := strings.Split(fileScanner.Text(), " ") //splits argument into dirname and youtube link
		Download(ss[1], ss[0])                       //apparently yt-dlp deadlocks when bieng multithreaded
		//go Download(ss[0], ss[1])                   //destroy every network you are in using a only file with over 10links
	}
	readFile.Close()
}

func Download(url string, foldername string) {
	fmt.Println("Downloading: ", foldername)
	out, e := exec.Command("cmd.exe", "/c", "mkdir", foldername).Output() //why windows just why do you have to make things harder
	CheckError(e)
	fmt.Printf("%s", out)
	cmd := exec.Command("..\\yt-dlp", "-f", "ba", "-x", "--audio-format", "mp3", url, "-o", "%(title)s.%(ext)s") //yt-dlp -f 'ba' -x --audio-format mp3 $2 -o '%(title)s.%(ext)s' ORIGINAL COMMAND
	cmd.Dir = ".\\" + foldername
	e = cmd.Run()
	CheckError(e)
	fmt.Printf("%s", out)
	//defer wg.Done()
}

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}
