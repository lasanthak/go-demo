package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	tabStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(lipgloss.Color("#04B575")).
			Padding(0, 1)

	activeTabStyle = tabStyle.
			BorderForeground(lipgloss.Color("#FF5F87"))

	windowStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("#04B575")).
			Border(lipgloss.NormalBorder()).
			Padding(1, 2)

	listItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = listItemStyle.
				Foreground(lipgloss.Color("#FF5F87"))
)

type tab int

const (
	tabTasks tab = iota
	tabStats
	tabSettings
)

type Task struct {
	Title     string
	Completed bool
	CreatedAt time.Time
}

type model struct {
	activeTab tab
	width     int
	height    int

	// Tasks tab
	tasks       []Task
	taskCursor  int
	newTaskText string
	inputMode   bool

	// Stats tab
	stats map[string]int

	// Settings tab
	settings map[string]string
}

func initialModel() model {
	return model{
		activeTab: tabTasks,
		tasks: []Task{
			{Title: "Learn Go", Completed: true, CreatedAt: time.Now().Add(-2 * time.Hour)},
			{Title: "Build TUI app", Completed: false, CreatedAt: time.Now().Add(-1 * time.Hour)},
			{Title: "Deploy to production", Completed: false, CreatedAt: time.Now()},
		},
		stats: map[string]int{
			"Total Tasks":   3,
			"Completed":     1,
			"In Progress":   2,
			"Created Today": 3,
		},
		settings: map[string]string{
			"Theme":    "Dark",
			"Language": "English",
			"Timezone": "UTC",
			"AutoSave": "Enabled",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.MouseMsg:
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			// Tabs are rendered in a single row, starting at y=2 (after title)
			tabRowY := 2
			if msg.Y == tabRowY {
				tabNames := []string{"Tasks", "Stats", "Settings"}
				x := 0
				for i, name := range tabNames {
					tabWidth := len(name) + 4 // Approx: 2 border, 2 padding
					if msg.X >= x && msg.X < x+tabWidth {
						m.activeTab = tab(i)
						break
					}
					x += tabWidth
				}
			}
		}

	case tea.KeyMsg:
		if m.inputMode && m.activeTab == tabTasks {
			return m.updateTaskInput(msg)
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.activeTab = (m.activeTab + 1) % 3
		case "shift+tab":
			m.activeTab = (m.activeTab + 2) % 3
		}

		switch m.activeTab {
		case tabTasks:
			return m.updateTasks(msg)
		case tabStats:
			return m.updateStats(msg)
		case tabSettings:
			return m.updateSettings(msg)
		}

	case tickMsg:
		return m, tickCmd()
	}

	return m, nil
}

func (m model) updateTaskInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if strings.TrimSpace(m.newTaskText) != "" {
			m.tasks = append(m.tasks, Task{
				Title:     m.newTaskText,
				Completed: false,
				CreatedAt: time.Now(),
			})
			m.newTaskText = ""
		}
		m.inputMode = false
	case "esc":
		m.newTaskText = ""
		m.inputMode = false
	case "backspace":
		if len(m.newTaskText) > 0 {
			m.newTaskText = m.newTaskText[:len(m.newTaskText)-1]
		}
	default:
		m.newTaskText += msg.String()
	}
	return m, nil
}

func (m model) updateTasks(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.taskCursor > 0 {
			m.taskCursor--
		}
	case "down", "j":
		if m.taskCursor < len(m.tasks)-1 {
			m.taskCursor++
		}
	case "enter", " ":
		if len(m.tasks) > 0 {
			m.tasks[m.taskCursor].Completed = !m.tasks[m.taskCursor].Completed
		}
	case "n":
		m.inputMode = true
		m.newTaskText = ""
	case "d":
		if len(m.tasks) > 0 {
			m.tasks = append(m.tasks[:m.taskCursor], m.tasks[m.taskCursor+1:]...)
			if m.taskCursor >= len(m.tasks) && len(m.tasks) > 0 {
				m.taskCursor = len(m.tasks) - 1
			}
		}
	}
	return m, nil
}

func (m model) updateStats(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Stats are read-only, just update the values
	completed := 0
	for _, task := range m.tasks {
		if task.Completed {
			completed++
		}
	}

	m.stats["Total Tasks"] = len(m.tasks)
	m.stats["Completed"] = completed
	m.stats["In Progress"] = len(m.tasks) - completed

	// Simulate some changing stats
	m.stats["Memory Usage"] = 45 + rand.Intn(20)
	m.stats["CPU Usage"] = 10 + rand.Intn(30)

	return m, nil
}

func (m model) updateSettings(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Settings are read-only in this demo
	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Header
	title := titleStyle.Render("📋 Task Manager TUI")

	// Tabs
	var tabs []string
	tabNames := []string{"Tasks", "Stats", "Settings"}
	for i, name := range tabNames {
		if tab(i) == m.activeTab {
			tabs = append(tabs, activeTabStyle.Render(name))
		} else {
			tabs = append(tabs, tabStyle.Render(name))
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	// Content based on active tab
	var content string
	switch m.activeTab {
	case tabTasks:
		content = m.renderTasks()
	case tabStats:
		content = m.renderStats()
	case tabSettings:
		content = m.renderSettings()
	}

	// Footer
	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Render("Tab/Shift+Tab: Switch tabs • q: Quit")

	// Combine everything
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		tabRow,
		windowStyle.Width(m.width-4).Height(m.height-8).Render(content),
		footer,
	)
}

func (m model) renderTasks() string {
	var items []string

	if m.inputMode {
		items = append(items, fmt.Sprintf("➤ New task: %s_", m.newTaskText))
		items = append(items, "")
	}

	for i, task := range m.tasks {
		var style lipgloss.Style
		if i == m.taskCursor && !m.inputMode {
			style = selectedItemStyle
		} else {
			style = listItemStyle
		}

		checkbox := "☐"
		if task.Completed {
			checkbox = "☑"
		}

		timeAgo := time.Since(task.CreatedAt).Truncate(time.Minute)
		item := fmt.Sprintf("%s %s (%v ago)", checkbox, task.Title, timeAgo)
		items = append(items, style.Render(item))
	}

	if len(items) == 0 {
		items = append(items, "No tasks yet. Press 'n' to create one!")
	}

	help := "\nControls:\n" +
		"↑/k: Move up • ↓/j: Move down • Space/Enter: Toggle completion\n" +
		"n: New task • d: Delete task"

	return strings.Join(items, "\n") + "\n" + help
}

func (m model) renderStats() string {
	var items []string

	for key, value := range m.stats {
		items = append(items, fmt.Sprintf("%-15s %s", key+":", strconv.Itoa(value)))
	}

	// Add some charts/graphs using ASCII
	items = append(items, "\nTask Completion Progress:")
	completed := m.stats["Completed"]
	total := m.stats["Total Tasks"]
	if total > 0 {
		percentage := float64(completed) / float64(total) * 100
		progressBar := strings.Repeat("█", int(percentage/10)) + strings.Repeat("░", 10-int(percentage/10))
		items = append(items, fmt.Sprintf("[%s] %.1f%%", progressBar, percentage))
	}

	// System stats visualization
	items = append(items, "\nSystem Resources:")
	memUsage := m.stats["Memory Usage"]
	cpuUsage := m.stats["CPU Usage"]

	memBar := strings.Repeat("█", memUsage/5) + strings.Repeat("░", 20-memUsage/5)
	cpuBar := strings.Repeat("█", cpuUsage/5) + strings.Repeat("░", 20-cpuUsage/5)

	items = append(items, fmt.Sprintf("Memory: [%s] %d%%", memBar, memUsage))
	items = append(items, fmt.Sprintf("CPU:    [%s] %d%%", cpuBar, cpuUsage))

	return strings.Join(items, "\n")
}

func (m model) renderSettings() string {
	var items []string

	items = append(items, "Application Settings:\n")

	for key, value := range m.settings {
		items = append(items, fmt.Sprintf("%-12s %s", key+":", value))
	}

	items = append(items, "\nShortcuts:")
	items = append(items, "Tab/Shift+Tab  Switch between tabs")
	items = append(items, "q              Quit application")
	items = append(items, "↑/↓ or k/j     Navigate lists")
	items = append(items, "Space/Enter    Select/toggle items")

	return strings.Join(items, "\n")
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
