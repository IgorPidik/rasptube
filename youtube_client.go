package main

import (
	"errors"
	"github.com/kkdai/youtube/v2"
	"sort"
	"strings"
)

type YoutubeClient struct {
	youtube.Client
}

type Video youtube.Video

type AudioFormats []youtube.Format

var AudioQuality = map[string]int{
	"AUDIO_QUALITY_LOW":    0,
	"AUDIO_QUALITY_MEDIUM": 1,
	"AUDIO_QUALITY_HIGH":   2,
}

func (v *Video) getAudioFormats() AudioFormats {
	var audioFormats []youtube.Format
	for _, format := range v.Formats {
		if strings.Contains(format.MimeType, "audio/") && format.AudioQuality != "" {
			audioFormats = append(audioFormats, format)
		}
	}
	return AudioFormats(audioFormats)
}

func (a AudioFormats) sortAudioQuality() AudioFormats {
	sort.Slice(a, func(i, j int) bool {
		if a[i].Bitrate > a[j].Bitrate {
			return true
		}

		return AudioQuality[a[i].AudioQuality] > AudioQuality[a[j].AudioQuality]
	})
	return a
}

func (a AudioFormats) getBestAudioQualityFormat() *youtube.Format {
	sortedAudio := a.sortAudioQuality()
	if len(sortedAudio) == 0 {
		return nil
	}

	return &sortedAudio[0]
}

func (c *YoutubeClient) GetBestAudioStreamUrl(videoUrl string) (string, error) {
	videoInfo, err := c.GetVideo(videoUrl)
	if err != nil {
		return "", err
	}
	video := Video(*videoInfo)
	bestAudioFormat := video.getAudioFormats().getBestAudioQualityFormat()
	if bestAudioFormat == nil {
		return "", errors.New("no audio format available")
	}

	return c.GetStreamURL(videoInfo, bestAudioFormat)
}
