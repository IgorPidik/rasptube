package main

import (
	"fmt"
	vlc "github.com/adrg/libvlc-go/v3"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ytClient := YoutubeClient{}
	url, streamUrlErr := ytClient.GetBestAudioStreamUrl("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	fmt.Println(streamUrlErr)

	errVlc := vlc.Init("--quiet", "--no-xlib")
	fmt.Println(errVlc)
	player, _ := vlc.NewPlayer()
	_, e := player.LoadMediaFromURL(url)
	fmt.Println(e)
	ee := player.Play()
	fmt.Println(ee)
	var first string

	// Taking input from user
	fmt.Scanln(&first)

	player.SetMediaPosition(0.5)
	fmt.Scanln(&first)
}
