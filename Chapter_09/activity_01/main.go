package main

import (
	author "activity_01/author"
)

func main() {
	// create an author instance
	authorInstance := author.NewAuthor("Jane Doe", "jane@example.com")
	// write and review instance
	chapterTitle := "Introduction to Go Modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability"

	authorInstance.WriteChapter(chapterTitle, chapterContent)
	authorInstance.ReviewChapter(chapterTitle, "This chapter looks great, but let's add some more examples")
	authorInstance.FinalizeChapter(chapterTitle)
}
