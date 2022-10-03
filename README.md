# Rasptube
Do you ever think to yourself: `I have this SICK raspotify setup. But I want to listen to these remixes/compilations 
that are not available on spotify and I am too lazy to VNC into my raspberry and open the browser.`? Worry no more, I got you! 

## Server-Client configuration
If you want to run the player on your pi and want to control it remotely without having to ssh into your pi you can use the server-client configuraiton:
- [server](https://github.com/IgorPidik/rasptube-server)
- [client](https://github.com/IgorPidik/rasptube-client)

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
