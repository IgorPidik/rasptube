package main

import (
	"fmt"
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/schollz/progressbar/v3"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ytClient := YoutubeClient{}
	url, streamUrlErr := ytClient.GetBestAudioStreamUrl("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	check(streamUrlErr)

	errVlc := vlc.Init("--quiet", "--no-xlib")
	check(errVlc)
	player, playerErr := vlc.NewPlayer()
	check(playerErr)

	playerEM, emErr := player.EventManager()
	check(emErr)

	media, mediaErr := player.LoadMediaFromURL(url)
	check(mediaErr)

	bar := progressbar.NewOptions64(
		100,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(10),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)

	playerEM.Attach(vlc.MediaPlayerPositionChanged, func(event vlc.Event, i interface{}) {
		duration, _ := media.Duration()
		position, _ := player.MediaPosition()
		durationSeconds := int(duration.Seconds())
		currentSeconds := int(float32(durationSeconds) * position)
		bar.ChangeMax(durationSeconds)
		bar.Set(currentSeconds)
	}, nil)

	playErr := player.Play()
	check(playErr)

	var first string
	fmt.Scanln(&first)

	media.Release()
	player.Release()
	vlc.Release()
}
