package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func ValidateCookies(cookies string) error {
	if !strings.Contains(cookies, "wp_logged_in_cookie=") {
		return fmt.Errorf("invalid cookies: wp_logged_in_cookie is required")
	}
	return nil
}

func GetPrompt() (string, error) {
	in := bufio.NewReader(os.Stdin)
	prompt, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	prompt = prompt[:len(prompt)-1] // remove last symbol (end of the string)

	return strings.TrimSpace(prompt), nil
}

func GetLargeInput() (string, error) {
	fmt.Println("Paste your cookies and press Enter twice to finish:")

	// Create a scanner with a large buffer
	const maxCapacity = 1024 * 1024 // 1MB buffer
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var input bytes.Buffer
	emptyLineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// If we get an empty line, increment counter
		if len(strings.TrimSpace(line)) == 0 {
			emptyLineCount++
			// Break after two consecutive empty lines
			if emptyLineCount >= 2 {
				break
			}
			continue
		}

		// Reset empty line counter if we get actual content
		emptyLineCount = 0

		// Add the line to our input buffer
		if input.Len() > 0 {
			input.WriteString("\n")
		}
		input.WriteString(line)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return strings.TrimSpace(input.String()), nil
}

func ReadCookiesFromFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading cookies file: %v", err)
	}
	return strings.TrimSpace(string(content)), nil
}

func ShowProgress(page int) {
	// Clear the current line
	fmt.Printf("\r")
	fmt.Printf("Processing page [%d]", page)
}

func WaitForKeyPress() {
	fmt.Println("Press any key to exit...")

	// Get the file descriptor for stdin
	fd := int(os.Stdin.Fd())

	// Save the original terminal state
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("\nError:", err)
		return
	}
	defer term.Restore(fd, oldState)

	// Read a single byte (one keypress)
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer)
}
