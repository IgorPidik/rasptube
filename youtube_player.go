package main

import (
	vlc "github.com/adrg/libvlc-go/v3"
)

type YoutubePlayer struct {
	YTClient              *YoutubeClient
	VLCPlayer             *vlc.Player
	VLCPlayerEventManager *vlc.EventManager
	CurrentMedia          *vlc.Media
}

func NewYoutubePlayer(ytClient *YoutubeClient, vlcPlayer *vlc.Player) (*YoutubePlayer, error) {
	if ytClient == nil {
		ytClient = &YoutubeClient{}
	}

	if vlcPlayer == nil {
		if errVLCInit := vlc.Init("--quiet", "--no-xlib"); errVLCInit != nil {
			return nil, errVLCInit
		}

		player, playerErr := vlc.NewPlayer()
		if playerErr != nil {
			return nil, playerErr
		}
		vlcPlayer = player
	}

	eventManager, eventManagerErr := vlcPlayer.EventManager()
	if eventManagerErr != nil {
		return nil, eventManagerErr
	}

	return &YoutubePlayer{
		YTClient:              ytClient,
		VLCPlayer:             vlcPlayer,
		VLCPlayerEventManager: eventManager,
	}, nil
}

func (p *YoutubePlayer) Play(url string) error {
	url, streamUrlErr := p.YTClient.GetBestAudioStreamUrl(url)
	if streamUrlErr != nil {
		return streamUrlErr
	}

	media, mediaErr := p.VLCPlayer.LoadMediaFromURL(url)
	if mediaErr != nil {
		return streamUrlErr
	}

	p.CurrentMedia = media
	return p.VLCPlayer.Play()
}

func (p *YoutubePlayer) Release() {
	if media, err := p.VLCPlayer.Media(); err == nil {
		media.Release()
	}
	p.VLCPlayer.Release()
	vlc.Release()
}
