package main

import (
  htgotts "github.com/hegedustibor/htgo-tts"
  handlers "github.com/hegedustibor/htgo-tts/handlers"
  voices "github.com/hegedustibor/htgo-tts/voices"
  "fmt"
  "os"
  "path/filepath"
)

var (
  speech = htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
)

func deleteMP3() {
  dirPath := "audio"
  files, err := os.ReadDir(dirPath)
  if err != nil {
    fmt.Println("Error: ", err)
  }
  for _, file := range files {
    err := os.Remove(filepath.Join(dirPath, file.Name()))
    if err != nil {
      fmt.Println("Error deleting file, ", err)
    }
  }
}

func Say(phrase string) {
  speech.Speak(phrase)
  go deleteMP3()
}
