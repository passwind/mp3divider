package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strconv"
)

var mp3parts = []string{
	"00:00",
	"01:49",
	"03:24",
	"05:11",
	"06:15",
	"07:50",
	"09:50",
	"11:48",
	"13:02",
	"14:39",
	"16:16",
	"17:47",
	"18:50",
	"20:33",
	"22:05",
	"23:38",
	"25:22",
	"27:09",
	"28:56",
	"30:39",
	"32:25",
	"34:03",
	"35:38",
	"37:44",
	"39:02",
	"42:57",
	"44:58",
	"48:27", // the total length
}

func main() {
	var inputFile string
	var outPath string

	flag.StringVar(&inputFile, "input", "京版小英二下课本完整音轨.mp3", "input file")
	flag.StringVar(&outPath, "outdir", ".", "output path")
	flag.Parse()

	for i := 0; i < len(mp3parts)-1; i++ {
		start := mp3parts[i]
		end := mp3parts[i+1]

		// ffmpeg -i 京版小英二下课本完整音轨.mp3 -ss 00:03:24 -to 00:05:11 -acodec copy L3.mp3
		command := "-i " + inputFile + " -ss 00:" + start + " -to 00:" + end + " -acodec copy " + outPath + "/L" + strconv.Itoa(i+1) + ".mp3"

		cmd := exec.Command("/usr/local/sbin/ffmpeg", command)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("exit with : %s", out.String())
	}
}
