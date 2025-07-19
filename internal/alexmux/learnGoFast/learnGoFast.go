package learnGoFast

import "going/internal/alexmux"

var video = alexmux.Video{
	Title: "Learn GO Fast: Full Tutorial",
	Tutorials: []alexmux.Tutorial{
		{
			Number: 1,
			Title:  "Hello World",
			Runner: RunT1,
		},
		{
			Number: 2,
			Title:  "Variables, Constants, and Basic Data Types",
			Runner: RunT2,
		},
		{
			Number: 3,
			Title:  "Functions and Control Structures",
			Runner: RunT3,
		},
		{
			Number: 4,
			Title:  "Arrays, Slices, Maps and Loops",
			Runner: RunT4,
		},
		{
			Number: 5,
			Title:  "Strings, Runes and Bytes",
			Runner: RunT5,
		},
		{
			Number: 6,
			Title:  "Structs and Interfaces",
			Runner: RunT6,
		},
		{
			Number: 7,
			Title:  "Pointers",
			Runner: RunT7,
		},
		{
			Number: 8,
			Title:  "Goroutines",
			Runner: RunT8,
		},
		{
			Number: 9,
			Title:  "Channels",
			Runner: RunT9,
		},
		{
			Number: 10,
			Title:  "Generics",
			Runner: RunT10,
		},
	},
}

func GetVideo() alexmux.Video {
	return video
}
