package util

import (
    "bytes"
    "fmt"
    "os/exec"
    "regexp"
    "strings"
    "github.com/acarl005/stripansi"
)

// Uses the Ollama CLI to generate a summary.
func GenerateOllamaSummary(studentID int, name string, age int, email string) (string, error) {
    inputText := fmt.Sprintf(
        "Please provide a concise summary of the following student information: ID: %d, Name: %s, Age: %d, Email: %s.",
        studentID, name, age, email,
    )

    cmd := exec.Command("ollama", "run", "llama3:latest")
    cmd.Stdin = strings.NewReader(inputText)

    var outBuf, errBuf bytes.Buffer
    cmd.Stdout = &outBuf
    cmd.Stderr = &errBuf

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("error generating summary: %v, stderr: %s", err, errBuf.String())
    }

    combinedOutput := outBuf.String() + errBuf.String()
    cleanedOutput := removeEscapeSequences(combinedOutput)

    if cleanedOutput == "" {
        return "Sorry, no summary could be generated. Please try again.", nil
    }

    return cleanedOutput, nil
}

func removeEscapeSequences(str string) string {
    cleaned := stripansi.Strip(str)
    reControl := regexp.MustCompile(`[\x00-\x1F\x7F]+`)
    cleaned = reControl.ReplaceAllString(cleaned, "")
    cleaned = strings.TrimSpace(cleaned)
    return cleaned
}



