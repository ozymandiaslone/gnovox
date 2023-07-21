module github.com/ozymandiaslone/gnovox

go 1.18

require (
	github.com/go-skynet/go-llama.cpp v0.0.0-20230719203055-f3a6ee0ef53d
	github.com/hegedustibor/htgo-tts v0.0.0-20230402053941-cd8d1a158135
)

require (
	github.com/hajimehoshi/go-mp3 v0.3.3 // indirect
	github.com/hajimehoshi/oto/v2 v2.2.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
)

replace github.com/go-skynet/go-llama.cpp => ./go-llama.cpp/
