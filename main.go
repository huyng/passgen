package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func main() {
	// Set up command-line flags
	var (
		length       int
		useClip      bool
		useSpecial   bool
		useLowerOnly bool
		showHelp     bool
	)

	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.IntVar(&length, "n", 10, "Password length (minimum 4)")
	flag.BoolVar(&useClip, "c", false, "Copy password to clipboard")
	flag.BoolVar(&useSpecial, "s", false, "Use special characters")
	flag.BoolVar(&useLowerOnly, "l", false, "Lowercase only")
	flag.Parse()

	if showHelp {
		printHelp()
		return
	}

	if length < 4 {
		fmt.Println("Error: Password length must be at least 4 characters")
		printHelp()
		os.Exit(1)
	}

	password, err := generatePassword(length, useSpecial, useLowerOnly)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	if useClip {
		err := copyToClipboard(password)
		if err != nil {
			fmt.Printf("Generated password: %s\n", password)
			fmt.Printf("Warning: Could not copy to clipboard: %v\n", err)
			return
		}
		fmt.Println("Password has been generated and copied to clipboard")
	} else {
		fmt.Printf("Generated password: %s\n", password)
	}
}

func printHelp() {
	fmt.Println("Password Generator - Creates secure random passwords")
	fmt.Println("Usage: passgen [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  -n N    Set password length (default: 10, minimum: 4)")
	fmt.Println("  -c      Copy the generated password to the clipboard")
	fmt.Println("  -s      Include special characters in the password")
	fmt.Println("  -l      Use only lowercase letters in the password")
	fmt.Println("  -h      Show this help message")
	fmt.Println("\nBy default, the password is printed to stdout.")
}

func generatePassword(length int, useSpecial bool, lowerOnly bool) (string, error) {
	allChars := lowerChars + digitChars
	sets := []string{lowerChars, digitChars}

	if !lowerOnly {
		allChars += upperChars
		sets = append(sets, upperChars)
	}
	if useSpecial {
		allChars += specialChars
		sets = append(sets, specialChars)
	}
	buf := make([]byte, length)

	// Ensure we have at least one character from each character set
	for i, set := range sets {
		char, err := randomChar(set)
		if err != nil {
			return "", err
		}
		buf[i] = char
	}

	// Fill remaining characters randomly
	for i := len(sets); i < length; i++ {
		char, err := randomChar(allChars)
		if err != nil {
			return "", err
		}
		buf[i] = char
	}

	// Shuffle the characters
	for i := range buf {
		j, err := randInt(i, len(buf))
		if err != nil {
			return "", err
		}
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf), nil
}

func randomChar(chars string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, err
	}
	return chars[n.Int64()], nil
}

func randInt(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return min + int(n.Int64()), nil
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("neither xclip nor xsel available")
		}
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
