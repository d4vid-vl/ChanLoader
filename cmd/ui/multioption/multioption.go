package multioption

import (
	"fmt"

	"ChanLoader/cmd/steps"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Change this
var (
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCCAA")).Bold(true)
	titleStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#FFCCAA")).Foreground(lipgloss.Color("#880000")).Bold(true).Padding(0, 1, 0)
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#a00000")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#aa0000"))
	descriptionStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF2DB"))
)

// A Selection represents a choice made in a multiSelect step
type Selection struct {
	Choice string
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(value string) {
	s.Choice = value
}

// A multiSelect.model contains the data for the multiSelect step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	cursor   int
	options  []steps.Item
	selected map[int]struct{}
	choice   *Selection
	header   string
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialModelMulti initializes a multiSelect step with
// the given data
func InitialModelMulti(options []steps.Item, selection *Selection, header string) model {
	return model{
		options:  options,
		selected: make(map[int]struct{}),
		choice:   selection,
		header:   titleStyle.Render(header),
	}
}

// Update is called when "things happen", it checks for
// important keystrokes to signal when to quit, change selection,
// and confirm the selection.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			if len(m.selected) == 1 {
				for selectedKey := range m.selected {
					if selectedKey >= len(m.options) {
						break // Make sure the index is within bounds
					}
					m.choice.Update(m.options[selectedKey].Title)
					m.cursor = selectedKey
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// View is called to draw the multiInput step
func (m model) View() string {
	s := m.header + "\n\n"

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			choice.Title = selectedItemStyle.Render(choice.Title)
			choice.Desc = selectedItemDescStyle.Render(choice.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("x")
		}

		title := focusedStyle.Render(choice.Title)
		description := descriptionStyle.Render(choice.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n\n", focusedStyle.Render("y"))
	return s
}
