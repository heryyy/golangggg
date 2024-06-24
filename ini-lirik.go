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
		{"Tiba-tiba aku jatuh cinta", 100 * time.Millisecond},
		{"Diam-diam kau pun juga cinta", 100 * time.Millisecond},
		{"Kita berdua belum punya kekasih", 110 * time.Millisecond},
		{"Saling mendekati", 100 * time.Millisecond},

		{"Perasaan tak bisa berdusta", 110 * time.Millisecond},
		{"Bahagia terasa sempurna", 110 * time.Millisecond},
		{"Kita berdua belum punya kekasih", 110 * time.Millisecond},
		{"Tunggu apa lagi", 140 * time.Millisecond},

		{"Katakan cinta bila kau cinta", 120 * time.Millisecond},
		{"Hati ini meminta", 120 * time.Millisecond},
		{"Kau lebih dari teman berbagi", 120 * time.Millisecond},
		{"Jadi kekasihku saja", 120 * time.Millisecond},

		{"Perasaan tak bisa berdusta", 120 * time.Millisecond},
		{"Bahagia terasa sempurna", 120 * time.Millisecond},
		{"Kita berdua belum punya kekasih", 120 * time.Millisecond},
		{"Tunggu apalagi", 120 * time.Millisecond},

		{"Katakan cinta bila kau cinta", 120 * time.Millisecond},
		{"Hati ini meminta", 130 * time.Millisecond},
		{"Kau lebih dari teman berbagi", 140 * time.Millisecond},
		{"Jadi kekasihku saja", 120 * time.Millisecond},

		{"Cinta katakan cinta", 120 * time.Millisecond},
		{"Hati ini meminta", 120 * time.Millisecond},
		{"Kau lebih dari teman berbagi", 170 * time.Millisecond},
		{"Jadilah kekasihku", 110 * time.Millisecond},

		{"Katakan (Cinta bila kau cinta) cinta", 120 * time.Millisecond},
		{"Hati ini meminta", 130 * time.Millisecond},
		{"Kau lebih dari teman berbagi", 120 * time.Millisecond},
		{"Jadi kekasihku saja", 120 * time.Millisecond},

		{"Cinta katakan cinta", 120 * time.Millisecond},
		{"Hati ini meminta", 180 * time.Millisecond},
		{"Kau lebih dari teman berbagi", 170 * time.Millisecond},
		{"Jadi kekasihku saja", 150 * time.Millisecond},

		{"Jadi kekasihku saja", 150 * time.Millisecond},
		{"Jadi kekasihku saja (awas salting)", 140 * time.Millisecond},
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
