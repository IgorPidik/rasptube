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
		length, _ := ytPlayer.VLCPlayer.MediaLength()
		position, _ := ytPlayer.VLCPlayer.MediaPosition()
		currentMilliseconds := int(position * float32(length))
		bar.ChangeMax(length)
		bar.Set(currentMilliseconds)
	}, nil)

	ytPlayer.VLCPlayerEventManager.Attach(vlc.MediaPlayerEndReached, func(event vlc.Event, i interface{}) {
		// use go routine to make sure the execution finishes, even after vlc kills the thread
		go func() {
			bar.Finish()
			ytPlayer.Release()
			os.Exit(0)
		}()
	}, nil)

	fmt.Scanln()
}
