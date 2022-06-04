# Rasptube
Do you ever think to yourself: `I have this SICK raspotify setup. But I want to listen to these remixes/compilations 
that are not available on spotify and I am too lazy to VNC into my raspberry and open the browser.`? Worry no more, I got you! 

## Prerequisites
- [youtube-dl](https://github.com/ytdl-org/youtube-dl)
- [lib-vlc](https://github.com/adrg/libvlc-go#prerequisites)

## Build
```console
$ go build
```

## Run
```console
$ ./rasptube https://www.youtube.com/watch?v=dQw4w9WgXcQ 
```

Press `ENTER` or `CTRL-C` to exit.

## TODO V1
- [x] Play Youtube links
- [x] Show progress bar
- [ ] Error handling
- [ ] Playlists
- [ ] Terminal UI
  - [ ] Playback options

## TODO V2
- [ ] Convert to a background service
- [ ] Client terminal UI
- [ ] Remote control
