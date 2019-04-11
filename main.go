package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// 京版小英二下课本完整音轨.mp3
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

// 京版小英四下课本完整音轨
var mp3parts1 = []string{
	"00:00",
	"03:30",
	"06:08",
	"07:58",
	"10:30",
	"13:19",
	"15:33",
	"17:47",
	"19:45",
	"23:19",
	"25:10",
	"27:00",
	"29:18",
	"32:50",
	"35:56",
	"38:10",
	"41:25",
	"43:13",
	"45:25",
	"47:25",
	"49:25",
}

func main() {
	var inputFile string
	var outPath string

	flag.StringVar(&inputFile, "input", "京版小英四下课本完整音轨.mp3", "input file")
	flag.StringVar(&outPath, "outdir", ".", "output path")
	flag.Parse()

	parts := mp3parts1

	for i := 0; i < len(parts)-1; i++ {
		start := parts[i]
		end := parts[i+1]

		if len(start) == 5 {
			start = "00:" + start
		}
		if len(end) == 5 {
			end = "00:" + end
		}

		// ffmpeg -i 京版小英二下课本完整音轨.mp3 -ss 00:03:24 -to 00:05:11 -acodec copy L3.mp3
		cmd := exec.Command("/usr/local/sbin/ffmpeg",
			"-i", inputFile,
			"-ss", start,
			"-to", end,
			"-acodec", "copy",
			fmt.Sprintf("%s/L%02d.mp3", outPath, (i+1)))
		var out bytes.Buffer
		cmd.Stdout = &out

		var outErr bytes.Buffer
		cmd.Stderr = &outErr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("error output: %s", outErr.String())
			log.Fatal(err)
		}
		fmt.Printf("exit with : %s", out.String())
	}
}
