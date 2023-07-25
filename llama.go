package main
import (
  "fmt"
  "strings"
  "os"
  "time"
  
  llama "github.com/go-skynet/go-llama.cpp"
)

var (
  threads = 8
  tokens = 2224
  gpulayers = 0
  modelPath = "./models/wizardLM-7B.ggmlv3.q4_0.bin"
  sentenceQueue chan string
  sentenceBuilder strings.Builder
)

func init() {
  sentenceQueue = make(chan string, 100)
  go processQueue()
}

func Inference() {
  l, err := llama.New(modelPath)
  if err != nil {
    fmt.Println("Loading model failed:", err.Error())
    os.Exit(1)
  }
  fmt.Printf("Model loaded succesfully. \n")

  for {
    text := ReadMultiLineInput()
    _, err := l.Predict(text, llama.Debug, llama.SetTokenCallback(caller), llama.SetTokens(tokens), llama.SetThreads(threads), llama.SetTopK(90), llama.SetTopP(0.87))
    if err != nil {
      panic(err)
    }
    embeds, err := l.Embeddings(text)
    if err != nil {
      fmt.Printf("Embeddings: error %s \n", err.Error())
    }
    fmt.Printf("Embeddings: %v", embeds)
    fmt.Printf("\n\n")
  }
}

func caller(token string) bool {
  fmt.Print(token)
  sentenceBuilder.WriteString(token)
  if strings.ContainsAny(token, ".!?:") {
    sentenceQueue <- sentenceBuilder.String()
    sentenceBuilder.Reset()
  }
  return true
}

func processQueue() {
  for {
    select {
    case sentence := <- sentenceQueue:
      fmt.Println("speaking Sentence!")
      Say(sentence)
      time.Sleep(50* time.Millisecond)
    default:
      time.Sleep(50* time.Millisecond)
      continue
    }
  }
}


