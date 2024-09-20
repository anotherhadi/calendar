package main

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/x/ansi"
)

var ansiStyleRegexp = regexp.MustCompile(`\x1b[[\d;]*m`)

func (m model) drawHelp(src, sub string) string {
	wrappedBG := src
	overlay := sub
	if m.Width == 0 {
		return ""
	}
	row := len(strings.Split(wrappedBG, "\n")) - len(strings.Split(overlay, "\n"))
	col := 0

	bgLines := strings.Split(wrappedBG, "\n")
	overlayLines := strings.Split(overlay, "\n")

	for i, overlayLine := range overlayLines {
		bgLine := bgLines[i+row]
		if len(bgLine) < col {
			bgLine += strings.Repeat(" ", col-len(bgLine)) // add padding
		}

		bgLeft := ansi.Truncate(bgLine, col, "")
		bgRight := truncateLeft(bgLine, col+ansi.StringWidth(overlayLine))

		bgLines[i+row] = bgLeft + overlayLine + bgRight
	}

	result := strings.Join(bgLines, "\n")
	return result
}

func truncateLeft(line string, padding int) string {
	if strings.Contains(line, "\n") {
		panic("line must not contain newline")
	}

	wrapped := strings.Split(ansi.Hardwrap(line, padding, true), "\n")
	if len(wrapped) == 1 {
		return ""
	}

	var ansiStyle string
	ansiStyles := ansiStyleRegexp.FindAllString(wrapped[0], -1)
	if l := len(ansiStyles); l > 0 {
		ansiStyle = ansiStyles[l-1]
	}

	return ansiStyle + strings.Join(wrapped[1:], "")
}
