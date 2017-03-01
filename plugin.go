package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Recime/recime-cli/cmd"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type botimizeIo struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Type     string `json:"type"`
}

type pkg struct {
	Name string
	UID  string
}

type plugin struct {
}

func (p *plugin) Add(name string) {
	switch strings.ToLower(name) {
	case "botimize.io":
		source := fmt.Sprintf("%s/plugin", apiEndpoint)

		uid := cmd.GetUID()

		jsonBody, err := json.Marshal(&botimizeIo{
			Name:     uid,
			Platform: "generic",
			Type:     "BotimizeIO",
		})

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner

		s.Start()

		res, err := http.Post(source, "application/json; charset=utf-8", bytes.NewBuffer(jsonBody))

		defer res.Body.Close()

		s.Stop()

		check(err)

		var data struct {
			Token string `json:"accessToken"`
		}

		bytes, err := ioutil.ReadAll(res.Body)

		check(err)

		json.Unmarshal(bytes, &data)

		if len(data.Token) > 0 {
			config := cmd.Config{Key: "BOTIMIZE_API_KEY", Value: data.Token, Source: source}
			config.Save()

		} else {
			red := color.New(color.FgRed).Add(color.Bold)
			red.Println("ERROR: Failed to add the plugin. Please try again later.")

			return
		}

	default:
		panic("INFO: Unsupported Platform.")
	}

	console := color.New(color.FgHiMagenta)

	fmt.Println("")

	console.Println("INFO: Plugin Added Succesfully.")

	fmt.Println("")
}
