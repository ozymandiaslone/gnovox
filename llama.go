package main

import (
  "fmt"
  "os"
  "bufio"
  "io"
  "strings"
  "time"
  
  llama "github.com/go-skynet/go-llama.cpp"
)

var (
  threads = 8
  tokens = 1024
  gpulayers = 0
  modelPath = "./models/wizardLM-7B.ggmlv3.q4_0.bin"
  sentenceQueue []string 
  sentenceBuilder strings.Builder
)

func init() {
  go processQueue()
}

func Inference() {
  l, err := llama.New(modelPath)
  if err != nil {
    fmt.Println("Loading model failed:", err.Error())
    os.Exit(1)
  }
  fmt.Printf("Model loaded succesfully. \n")
  reader := bufio.NewReader(os.Stdin)

  for {
    text := readMultiLineInput(reader)
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
  if strings.ContainsAny(token, ".!?") {
    sentenceQueue = append(sentenceQueue, sentenceBuilder.String())
    sentenceBuilder.Reset()
  }
  return true
}

func readMultiLineInput(reader *bufio.Reader) string {
	var lines []string
	fmt.Print(">>> ")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Printf("Reading the prompt failed: %s", err)
			os.Exit(1)
		}
    lines = append(lines, line)
    break
	}

	text := strings.Join(lines, "")
	fmt.Println("Sending", text)
	return text
}

func processQueue() {
  for {
   	if len(sentenceQueue) > 0 {
      if len(sentenceQueue) > 1 {
        fmt.Println("We got more than one in the queeueue!!!!!!")
      }
   		sentence := sentenceQueue[0]
      Say(sentence)
   		sentenceQueue = sentenceQueue[1:]
   	}
    time.Sleep(15 * time.Millisecond)
  }
}


