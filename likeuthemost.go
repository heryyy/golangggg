package main

import (
	"fmt"
	"time"
)

func printLyrics() {
	lines := []struct {
		line      string
		charDelay time.Duration
	}{
		{"That when they see you smiling, I go crazy for you", 100 * time.Millisecond},
		{"'Cause you're the one that I like, I can't deny", 100 * time.Millisecond},
		{"Every night you're on my mind", 110 * time.Millisecond},
		{"So, if I call you tonight", 100 * time.Millisecond},
		{"Will you pick up and give me your time?", 110 * time.Millisecond},
		{"Miss you every night, miss you all the time", 110 * time.Millisecond},
		{"No, I don't even know where to start", 110 * time.Millisecond},
	}

	delays := []time.Duration{
		600 * time.Millisecond, 600 * time.Millisecond, 500 * time.Millisecond, 1600 * time.Millisecond,
		600 * time.Millisecond, 700 * time.Millisecond, 700 * time.Millisecond, 1400 * time.Millisecond,
		700 * time.Millisecond, 500 * time.Millisecond, 500 * time.Millisecond, 3600 * time.Millisecond,
		500 * time.Millisecond, 600 * time.Millisecond, 600 * time.Millisecond, 600 * time.Millisecond,
		800 * time.Millisecond, 900 * time.Millisecond, 900 * time.Millisecond, 800 * time.Millisecond,
		800 * time.Millisecond, 900 * time.Millisecond, 900 * time.Millisecond, 16000 * time.Millisecond,
		1400 * time.Millisecond, 900 * time.Millisecond, 900 * time.Millisecond, 800 * time.Millisecond,
		800 * time.Millisecond, 600 * time.Millisecond, 500 * time.Millisecond, 500 * time.Millisecond,
		600 * time.Millisecond, 5000 * time.Millisecond,
	}

	for i, line := range lines {
		for _, char := range line.line {
			fmt.Print(string(char))
			time.Sleep(line.charDelay)
		}
		time.Sleep(delays[i%len(delays)])
		fmt.Println()
	}
}

func main() {
	printLyrics()
}
