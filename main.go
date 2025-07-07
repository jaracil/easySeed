package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	var wordCount string
	var creationMethod string
	var passPhrase1 string
	var passPhrase2 string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select 12 or 24 words").
				Options(
					huh.NewOption("12 words", "12"),
					huh.NewOption("24 words", "24"),
				).
				Value(&wordCount),

			huh.NewSelect[string]().
				Title("Select creation method").
				Options(
					huh.NewOption("Use pass phrase", "passphrase"),
					huh.NewOption("Create random seed", "random"),
				).
				Value(&creationMethod),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	var entropy []byte
	if creationMethod == "passphrase" {
		passwordForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter pass phrase").
					EchoMode(huh.EchoModePassword).
					Value(&passPhrase1),

				huh.NewInput().
					Title("Enter pass phrase again to confirm").
					EchoMode(huh.EchoModePassword).
					Value(&passPhrase2),
			),
		)

		err := passwordForm.Run()
		if err != nil {
			log.Fatal(err)
		}

		if passPhrase1 != passPhrase2 {
			fmt.Println("Pass phrases do not match")
			return
		}

		hash256 := sha256.Sum256([]byte(passPhrase1))

		if wordCount == "12" {
			entropy = hash256[:16]
		} else {
			entropy = hash256[:32]
		}
	} else {
		if wordCount == "12" {
			entropy, err = bip39.NewEntropy(128)
		} else {
			entropy, err = bip39.NewEntropy(256)
		}
		if err != nil {
			fmt.Println("Error generating random entropy")
			return
		}
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	words := strings.Fields(mnemonic)
	
	// Start interactive word reveal
	model := initialModel(words)
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	words    []string
	cursor   int
	revealed []bool
	finished bool
}

func initialModel(words []string) model {
	return model{
		words:    words,
		cursor:   0,
		revealed: make([]bool, len(words)),
		finished: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				// Hide current word before moving
				m.revealed[m.cursor] = false
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.words)-1 {
				// Hide current word before moving
				m.revealed[m.cursor] = false
				m.cursor++
			}
		case " ":
			// Toggle reveal for current word
			m.revealed[m.cursor] = !m.revealed[m.cursor]
		case "h":
			// Hide all words
			for i := range m.revealed {
				m.revealed[i] = false
			}
		case "enter":
			// Finish and exit
			m.finished = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	// Define styles
	numberStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("14")) // Cyan
	wordStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))   // Green
	hiddenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))  // Gray
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("13")) // Magenta
	
	s := "\n🔐 Seed Phrase Interactive Viewer\n\n"
	
	for i, word := range m.words {
		var line string
		
		// Number part
		numberText := fmt.Sprintf("%02d:", i+1)
		if i == m.cursor {
			numberText = cursorStyle.Render("► " + numberText)
		} else {
			numberText = "  " + numberStyle.Render(numberText)
		}
		
		// Word part
		var wordText string
		if m.revealed[i] {
			wordText = wordStyle.Render(word)
		} else {
			wordText = hiddenStyle.Render("••••••••")
		}
		
		line = fmt.Sprintf("%s %s", numberText, wordText)
		s += line + "\n"
	}
	
	s += "\n📋 Controls:\n"
	s += "  ↑/↓ or k/j: Navigate\n"
	s += "  Space: Reveal/Hide current word\n"
	s += "  h: Hide all words\n"
	s += "  Enter: Finish\n"
	s += "  q or Ctrl+C: Quit\n"
	
	return s
}
