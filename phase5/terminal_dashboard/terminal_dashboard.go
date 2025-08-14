package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

const barLength int = 30

type lastValues struct {
	netBytesSent uint64
	netBytesRecv uint64
}

type model struct {
	metrics map[string]float64
	history map[string][]float64
	last    lastValues
	width   int
	height  int
	status  string
}

func initialModel() model {
	return model{
		metrics: map[string]float64{
			"CPU":     0,
			"Memory":  0,
			"Disk":    0,
			"Network": 0,
		},
		history: map[string][]float64{
			"CPU":     make([]float64, barLength),
			"Memory":  make([]float64, barLength),
			"Disk":    make([]float64, barLength),
			"Network": make([]float64, barLength),
		},
		last: lastValues{
			netBytesSent: math.MaxUint64,
			netBytesRecv: math.MaxUint64,
		},
		status: "OK",
	}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return m, tea.Quit
		}

	case tickMsg:
		// Update metrics with simulated data
		cpuPercent := 0.0
		if v, e := cpu.Percent(0, false); e == nil {
			cpuPercent = v[0]
		}
		memPercent := 0.0
		if v, e := mem.VirtualMemory(); e == nil {
			memPercent = v.UsedPercent
		}
		diskPercent := 0.0
		if v, e := disk.Usage("/"); e == nil {
			diskPercent = v.UsedPercent
		}
		netPercentage := 0.0
		if v, e := net.IOCounters(false); e == nil && len(v) > 0 {
			if v[0].BytesSent > m.last.netBytesSent {
				netPercentage += float64(v[0].BytesSent-m.last.netBytesSent) / 100_000.0
			}
			if v[0].BytesRecv > m.last.netBytesRecv {
				netPercentage += float64(v[0].BytesRecv-m.last.netBytesRecv) / 100_000.0
			}
			if netPercentage < 0 {
				netPercentage = 0
			} else if netPercentage > 100 {
				netPercentage = 100
			}
			m.last.netBytesSent = v[0].BytesSent
			m.last.netBytesRecv = v[0].BytesRecv
		}

		m.metrics["CPU"] = cpuPercent
		m.metrics["Memory"] = memPercent
		m.metrics["Disk"] = diskPercent
		m.metrics["Network"] = netPercentage

		// Update history
		for key, value := range m.metrics {
			m.history[key] = append(m.history[key][1:], value)
		}

		t := time.Time(msg)
		m.status = t.Format("3:04:05 PM")

		return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading dashboard..."
	}

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		Width(m.width).
		Align(lipgloss.Center).
		Render("🖥️  System Monitor Dashboard")

	// Create metric displays
	var metricRows []string

	for _, metric := range []string{"CPU", "Memory", "Network", "Disk"} {
		value := m.metrics[metric]
		history := m.history[metric]

		// Create progress bar
		filled := int(math.Round(float64(barLength) * value / 100))
		bar := lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Render(
			fmt.Sprintf("%s%s", strings.Repeat("█", filled), strings.Repeat("░", barLength-filled)),
		)

		// Create mini sparkline
		sparkline := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7a7f55ff")).
			Render(m.createSparkline(history, barLength))

		row := fmt.Sprintf("%-8s %s %5.1f%%\n  %s", metric, bar, value, sparkline)
		metricRows = append(metricRows, row)
	}

	content := strings.Join(metricRows, "\n\n")

	// Add timestamp
	status := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Render(m.status)
	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Render("Press 'q' to quit")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		content,
		"",
		status,
		footer,
	)
}

func (m model) createSparkline(data []float64, width int) string {
	if len(data) == 0 {
		return strings.Repeat("_", width)
	}

	// Find min and max for scaling
	min, max := 33.0, 67.0
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if max == min {
		return strings.Repeat("▁", width)
	}

	// Create sparkline
	chars := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
	var result strings.Builder

	for i := 0; i < width && i < len(data); i++ {
		value := data[i]
		normalized := (value - min) / (max - min)
		charIndex := int(normalized * float64(len(chars)-1))
		if charIndex >= len(chars) {
			charIndex = len(chars) - 1
		}
		result.WriteString(chars[charIndex])
	}

	return result.String()
}

func main() {
	// if v, e := net.IOCounters(false); e == nil && len(v) > 0 {
	// 	fmt.Printf("BytesSent: %d\n", v[0].BytesSent)
	// 	fmt.Printf("BytesRecv: %d\n", v[0].BytesRecv)
	// }
	// os.Exit(0)

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
