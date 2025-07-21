package main

import (
	bootdev "bootdev/internal"
	"bootdev/internal/learnGo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Experimenting with cli menu")
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type screenOneModel struct {
	choices []bootdev.Course
	cursor  int
	choice  bootdev.Course
}

func initialModel() screenOneModel {
	//choices := make(map[int]bootdev.Course)
	//choices[0] = learnGo.LearnGoCourse

	return screenOneModel{
		choices: []bootdev.Course{
			learnGo.LearnGoCourse,
		},
	}
}

func (m screenOneModel) Init() tea.Cmd {
	return nil
}

func (m screenOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.choice = m.choices[m.cursor]
			m2 := screenTwo(m)
			return m2, m2.Init()
		default:
			return m, nil
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m screenOneModel) View() string {
	s := "Going: bootdev courses \n"
	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}
	return s
}

type screenTwoModel struct {
	course  bootdev.Course
	choices []bootdev.Chapter
	cursor  int
	choice  bootdev.Chapter
}

func (m screenTwoModel) Init() tea.Cmd {
	return nil
}

func screenTwo(m screenOneModel) screenTwoModel {
	var course bootdev.Course = m.choice
	/*chapters := make(map[int]bootdev.Chapter)
	for i, chapter := range course.Chapters {
		chapters[i] = chapter
	}*/
	return screenTwoModel{
		course:  course,
		choices: course.Chapters,
	}
}

func (m screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.choice = m.choices[m.cursor]
			m2 := screenThree(m)
			return m2, m2.Init()
			//return m, nil
		case "backspace":
			m2 := initialModel()
			return m2, initialModel().Init()
		default:
			return m, nil
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m screenTwoModel) View() string {
	s := "Going: Boot.dev courses \n"
	s += fmt.Sprintf("%s course chapters: \n", m.course.Title)
	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s %s (%d lessons)\n", cursor, "#"+strconv.Itoa(choice.Number), choice.Title, len(choice.Lessons))
	}
	return s
}

type screenThreeModel struct {
	course  bootdev.Course
	chapter bootdev.Chapter
	choices []bootdev.Lesson
	cursor  int
	choice  bootdev.Lesson
}

func (m screenThreeModel) Init() tea.Cmd {
	return nil
}

func screenThree(m screenTwoModel) screenThreeModel {
	var course bootdev.Course = m.course
	var chapter bootdev.Chapter = m.choice
	/*chapters := make(map[int]bootdev.Chapter)
	for i, chapter := range course.Chapters {
		chapters[i] = chapter
	}*/
	return screenThreeModel{
		course:  course,
		chapter: chapter,
		choices: chapter.Lessons,
	}
}

func (m screenThreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.choice = m.choices[m.cursor]
			m2 := screenFour(m)
			return m2, m2.Init()
			//return m, nil
		case "backspace":
			m2 := screenTwo(screenOneModel{
				choice: m.course,
			})
			return m2, m2.Init()
		default:
			return m, nil
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m screenThreeModel) View() string {
	s := "Going: Boot.dev courses \n"
	s += fmt.Sprintf("%s course chapter: %s\n", m.course.Title, m.chapter.Title)
	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, "#"+strconv.Itoa(choice.Number), choice.Title)
	}
	return s
}

type screenFourModel struct {
	course  bootdev.Course
	chapter bootdev.Chapter
	lesson  bootdev.Lesson
	choices []string
	cursor  int
	choice  int
}

func (m screenFourModel) Init() tea.Cmd {
	return nil
}

func screenFour(m screenThreeModel) screenFourModel {
	var course bootdev.Course = m.course
	var chapter bootdev.Chapter = m.chapter
	var lesson bootdev.Lesson = m.choice
	choices := []string{
		"Run Lesson",
		"Go Back",
	}
	return screenFourModel{
		course:  course,
		chapter: chapter,
		lesson:  lesson,
		choices: choices,
	}
}

func (m screenFourModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.choices[m.cursor] {
			case "Run Lesson":
				m2 := screenFive(m)
				return m2, m2.Init()
				//return m, nil
			case "Go Back":
				m2 := screenThree(screenTwoModel{
					course: m.course,
					choice: m.chapter,
				})
				return m2, m2.Init()
			default:
				return m, nil
			}
		case "backspace":
			m2 := screenThree(screenTwoModel{
				course: m.course,
				choice: m.chapter,
			})
			return m2, m2.Init()
		default:
			return m, nil
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m screenFourModel) View() string {
	s := "Going: Boot.dev courses \n"
	s += fmt.Sprintf("%s course chapter: %s\n", m.course.Title, m.chapter.Title)
	s += fmt.Sprintf("%s chapter lesson: #%d - %s\n", m.chapter.Title, m.lesson.Number, m.lesson.Title)
	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s \n", cursor, choice)
	}
	return s
}

// A message containing a line of output from the lesson.
type lessonOutputChunkMsg string

// A message to signal that the lesson's Runner has finished executing.
type lessonStreamFinishedMsg struct{}

// listenForLessonOutputCmd waits for the next line of output from the lesson's
// runner and sends it to the Update loop.
func listenForLessonOutputCmd(ch chan string) tea.Cmd {
	return func() tea.Msg {
		// Wait for a message on the channel. If the channel is closed,
		// 'ok' will be false.
		if output, ok := <-ch; ok {
			return lessonOutputChunkMsg(output)
		}
		// If the channel is closed, it means the runner is finished.
		return lessonStreamFinishedMsg{}
	}
}

func streamLessonOutput(runner func()) (chan string, tea.Cmd) {
	ch := make(chan string)
	var wg sync.WaitGroup

	// The main goroutine (G1) will run the process.
	go func() {
		// When G1 is completely done, it will close the channel.
		defer close(ch)

		r, w, _ := os.Pipe()

		// We have one goroutine (the reader) to wait for.
		wg.Add(1)

		// The reader goroutine (G2) reads from the pipe.
		go func() {
			// When G2 exits, signal that it's done.
			defer wg.Done()
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				// Send the line of output to the channel.
				ch <- scanner.Text()
			}
		}()

		// The writer part, happening back in G1.
		originalStdout := os.Stdout
		os.Stdout = w

		runner()

		os.Stdout = originalStdout
		// Close the writer end of the pipe. This will cause the reader's
		// scanner to stop, which will terminate G2.
		w.Close()

		// Wait for the reader goroutine (G2) to finish completely.
		// This prevents the race condition.
		wg.Wait()
	}()

	return ch, listenForLessonOutputCmd(ch)
}

type screenFiveModel struct {
	course      bootdev.Course
	chapter     bootdev.Chapter
	lesson      bootdev.Lesson
	choices     []string
	cursor      int
	output      string
	outputChan  chan string // To receive streaming output
	isStreaming bool        // To track if the process is running
}

func screenFive(m screenFourModel) screenFiveModel {
	return screenFiveModel{
		course:      m.course,
		chapter:     m.chapter,
		lesson:      m.lesson,
		choices:     []string{"Run Again", "Go Back"},
		isStreaming: true, // It starts streaming immediately
	}
}

func (m screenFiveModel) Init() tea.Cmd {
	// Start the streaming process.
	var cmd tea.Cmd
	m.outputChan, cmd = streamLessonOutput(m.lesson.Runner)
	return cmd
}

func (m screenFiveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// A new line of output has arrived.
	case lessonOutputChunkMsg:
		m.output += string(msg) + "\n"
		// Return the listener command again to wait for the next line.
		return m, listenForLessonOutputCmd(m.outputChan)

	// The stream has finished.
	case lessonStreamFinishedMsg:
		m.isStreaming = false
		// No more commands are needed, we just wait for user input.
		return m, nil

	case tea.KeyMsg:
		// Don't process key presses while the lesson is running.
		if m.isStreaming {
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.choices[m.cursor] {
			case "Run Again":
				// Clear output and start a new stream.
				m.output = ""
				m.isStreaming = true
				var cmd tea.Cmd
				m.outputChan, cmd = streamLessonOutput(m.lesson.Runner)
				return m, cmd
			case "Go Back":
				m2 := screenFour(screenThreeModel{
					course:  m.course,
					chapter: m.chapter,
					choice:  m.lesson,
				})
				return m2, m2.Init()
			}
		}
	}
	return m, nil
}
func (m screenFiveModel) View() string {
	s := fmt.Sprintf("Running Lesson: #%d - %s\n\n", m.lesson.Number, m.lesson.Title)

	s += m.output

	if m.isStreaming {
		s += "...\n" // A subtle indicator that it's still running.
	}

	s += "\n"

	// Only show choices after the stream is finished.
	if !m.isStreaming {
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s \n", cursor, choice)
		}
	}

	return s
}
