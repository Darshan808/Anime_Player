package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

type Response struct {
	Info struct {
		Title   string `json:"title"`
		ID      string `json:"id"`
		Episode string `json:"episode"`
	} `json:"info"`

	Stream struct {
		Multi struct {
			Main struct {
				URL string `json:"url"`
			} `json:"main"`
		} `json:"multi"`
	} `json:"stream"`
}

func main() {
	//Flags
	inTerminal := pflag.Bool("in-terminal", false, "Run in terminal mode")
	pflag.Parse()

	if len(pflag.Args()) < 2 {
		fmt.Println("Provide the name of anime and its episode number!")
		return
	}

	//Api URL
	URL := "https://api.amvstr.me/api/v2/stream/" + os.Args[1] + "/" + os.Args[2]

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println("An error occured while fetching the stream. Check anime name and episode number!", err.Error())
		return
	}

	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("Playing ", response.Info.Title, " episode ", response.Info.Episode)
	main_url := response.Stream.Multi.Main.URL

	var cmd *exec.Cmd

	if *inTerminal {
		cmd = exec.Command("mpv", "--no-config", "--vo=sixel", "--cache=yes", "--cache-secs=10", "--profile=sw-fast", "--vo-sixel-width=800", "--vo-sixel-height=400", "--framedrop=vo", main_url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd = exec.Command("mpv", main_url)
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
