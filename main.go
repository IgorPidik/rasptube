package main

import (
	"errors"
	"fmt"
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/schollz/progressbar/v3"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getProgressBar() *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		100,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(10),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)
	bar.RenderBlank()
	return bar
}

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("missing url"))
	}
	requestedUrl := os.Args[1]

	ytPlayer, ytPlayerErr := NewYoutubePlayer(nil, nil)
	check(ytPlayerErr)
	defer ytPlayer.Release()

	playErr := ytPlayer.Play(requestedUrl)
	check(playErr)

	bar := getProgressBar()
	ytPlayer.VLCPlayerEventManager.Attach(vlc.MediaPlayerPositionChanged, func(event vlc.Event, i interface{}) {
		duration, _ := ytPlayer.CurrentMedia.Duration()
		position, _ := ytPlayer.VLCPlayer.MediaPosition()
		durationSeconds := duration.Seconds()
		currentSeconds := int(durationSeconds * float64(position))
		bar.ChangeMax(int(durationSeconds))
		bar.Set(currentSeconds)
	}, nil)

	fmt.Scanln()
}
