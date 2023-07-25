package main
import (
  "bufio"
  "os"
  "fmt"
  "io"
  "strings"
)

func ReadMultiLineInput() string {
  reader := bufio.NewReader(os.Stdin)
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
  text = strings.TrimSpace(text)
  fmt.Println("MultiLineInput Detected:", text)
	return text
}
