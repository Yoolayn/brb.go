package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func newStyle(color string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Background(lipgloss.Color(color))
}

func main() {
	b := newStyle("#5BCEFA").Render("h")
	p := newStyle("#F5A9B8").Render("r")
	w := newStyle("#FFFFFF").Render("t")

	text := `
 %s%s%s%s%s%s%s%s   %s%s%s%s%s%s%s%s   %s%s%s%s%s%s%s%s
 %s%s     %s%s  %s%s     %s%s  %s%s     %s%s
 %s%s     %s%s  %s%s     %s%s  %s%s     %s%s
 %s%s     %s%s  %s%s     %s%s  %s%s     %s%s
 %s%s%s%s%s%s%s%s   %s%s%s%s%s%s%s%s   %s%s%s%s%s%s%s%s
 %s%s     %s%s  %s%s    %s%s%s  %s%s     %s%s
 %s%s     %s%s  %s%s     %s%s  %s%s     %s%s
 %s%s     %s%s  %s%s     %s%s  %s%s     %s%s
 %s%s%s%s%s%s%s%s   %s%s     %s%s  %s%s%s%s%s%s%s%s
`

	rendered := fmt.Sprintf(text,
		b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b,
		b, b, b, b, b, b, b, b, b, b, b, b,
		p, p, p, p, p, p, p, p, p, p, p, p,
		p, p, p, p, p, p, p, p, p, p, p, p,
		w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w,
		p, p, p, p, p, p, p, p, p, p, p, p, p,
		p, p, p, p, p, p, p, p, p, p, p, p,
		b, b, b, b, b, b, b, b, b, b, b, b,
		b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b,
	)

	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	fmt.Print(lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, rendered))
}
