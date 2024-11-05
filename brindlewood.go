package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "strings"
    "os/exec"
    "log"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run main.go <COMPANY> <ROLE>")
        os.Exit(1)
    }

    role := os.Args[1]
    company := os.Args[2]
    address := os.Args[3]

    // Open the base.txt file
    inputFile, err := os.Open("base.txt")
    if err != nil {
        fmt.Printf("Error opening base.txt: %v\n", err)
        os.Exit(1)
    }
    defer inputFile.Close()

    // Read the file contents
    scanner := bufio.NewScanner(inputFile)
    var content strings.Builder
    for scanner.Scan() {
        content.WriteString(scanner.Text() + "\n")
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading base.txt: %v\n", err)
        os.Exit(1)
    }

    // Perform the find and replace
    newText := strings.ReplaceAll(content.String(), "__COMPANY__", company)
    newText = strings.ReplaceAll(newText, "__ROLE__", role)
    newText = strings.ReplaceAll(newText, "__ADDRESS__", address)

/*
    // Write the new content to cover_letter.txt
    err = ioutil.WriteFile("cover_letter.txt", []byte(newText), 0644)
    if err != nil {
        fmt.Printf("Error writing to cover_letter.txt: %v\n", err)
        os.Exit(1)
    }
*/

    // Copying the cover letter to clipboard
		// Create a command to run 'wl-copy'
    cmd := exec.Command("wl-copy")

    // Use a bytes.NewReader to send the text as input to 'wl-copy'
    cmd.Stdin = bytes.NewReader([]byte(newText))

    // Execute the command
    err = cmd.Run()
    if err != nil {
        log.Fatalf("Failed to copy text to clipboard: %v", err)
    }
  

    fmt.Println("cover_letter copie to clipboard.")
}

