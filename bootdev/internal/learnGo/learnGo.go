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
			{
				Number: 2,
				Title:  "The Initial Statement of an If Block",
				Runner: c2_l2,
			},
			{
				Number: 4,
				Title:  "Switch",
				Runner: c2_l4,
			},
			{
				Number: 5,
				Title:  "Calculate Balance",
				Runner: c2_l5,
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
			{
				Number: 2,
				Title:  "Multiple Parameters",
				Runner: c3_l2,
			},
			{
				Number: 3,
				Title:  "Unit Test Lessons",
				Runner: c3_l3,
			},
			{
				Number: 7,
				Title:  "Passing Variables by Value",
				Runner: c3_l7,
			},
			{
				Number: 8,
				Title:  "Ignoring Return Values",
				Runner: c3_l8,
			},
			{
				Number: 9,
				Title:  "Named Return Values",
				Runner: c3_l9,
			},
			{
				Number: 12,
				Title:  "Explicit Returns",
				Runner: c3_l12,
			},
			{
				Number: 15,
				Title:  "Functions As Values",
				Runner: c3_l15,
			},
			{
				Number: 16,
				Title:  "Anonymous Functions",
				Runner: c3_l16,
			},
			{
				Number: 17,
				Title:  "Defer",
				Runner: c3_l17,
			},
			{
				Number: 18,
				Title:  "Block Scope",
				Runner: c3_l18,
			},
			{
				Number: 19,
				Title:  "Processing Orders",
				Runner: c3_l19,
			},
			{
				Number: 20,
				Title:  "Closures",
				Runner: c3_l20,
			},
			{
				Number: 21,
				Title:  "Currying",
				Runner: c3_l21,
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
