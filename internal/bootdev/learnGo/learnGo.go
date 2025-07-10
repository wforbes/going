package learnGo

import "going/internal/bootdev"

var learnGoChapters = []bootdev.Chapter{
	{
		Number: 1,
		Title:  "Variables",
		Runner: runChapter1Lessons,
	},
	{
		Number: 2,
		Title:  "Conditionals",
		Runner: runChapter2Lessons,
	},
}

var LearnGoCourse = bootdev.Course{
	Title:    "Learn Go",
	Chapters: learnGoChapters,
}
