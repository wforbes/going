package learnGo

import bootdev "bootdev/internal"

var LearnGoCourse = bootdev.Course{
	Title:    "Learn Go",
	Chapters: learnGoChapters,
}

var learnGoChapters = []bootdev.Chapter{
	{
		Number: 1,
		Title:  "Variables",
		Runner: runChapter1Lessons,
		Lessons: []bootdev.Lesson{
			{
				Number: 1,
				Title:  "Learn Go (for Developers)",
				Runner: c1_l1,
			},
			{
				Number: 2,
				Title:  "Basic Variables",
				Runner: c1_l2,
			},
			{
				Number: 3,
				Title:  "Short Variable Declarations",
				Runner: c1_l3,
			},
			{
				Number: 4,
				Title:  "Why Go?",
				Runner: c1_l4,
			},
			{
				Number: 8,
				Title:  "Type Sizes",
				Runner: c1_l8,
			},
			{
				Number: 11,
				Title:  "Go Is Statically Typed",
				Runner: c1_l11,
			},
			{
				Number: 14,
				Title:  "Same Line Declarations",
				Runner: c1_l14,
			},
			{
				Number: 17,
				Title:  "Constants",
				Runner: c1_l17,
			},
			{
				Number: 18,
				Title:  "Computed Constants",
				Runner: c1_l18,
			},
			{
				Number: 20,
				Title:  "Formatting Strings in Go",
				Runner: c1_l20,
			},
			{
				Number: 21,
				Title:  "Runes and Strings Encoding",
				Runner: c1_l21,
			},
			{
				Number: 25,
				Title:  "Format Practice",
				Runner: c1_l25,
			},
		},
	},
	{
		Number: 2,
		Title:  "Conditionals",
		Runner: runChapter2Lessons,
		Lessons: []bootdev.Lesson{
			{
				Number: 1,
				Title:  "Conditionals",
				Runner: c2_l1,
			},
		},
	},
	{
		Number: 3,
		Title:  "Functions",
		Runner: runChapter3Lessons,
		Lessons: []bootdev.Lesson{
			{
				Number: 1,
				Title:  "Functions",
				Runner: c3_l1,
			},
		},
	},
	{
		Number: 4,
		Title:  "Structs",
		Runner: runChapter4Lessons,
		Lessons: []bootdev.Lesson{
			{
				Number: 1,
				Title:  "Structs in Go",
				Runner: c4_l1,
			},
		},
	},
}
