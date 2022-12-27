package display

import (
	"bytes"

	"github.com/fatih/color"
)

const lineChar = "-"

func Line(lineLen int) string {
	buf := bytes.NewBufferString("")
	i := 0
	blue := color.New(color.FgHiBlue)
	for i < lineLen {
		blue.Fprint(buf, lineChar)
		i++
	}
	blue.Fprint(buf, "\n")
	return buf.String()
}

func YearInReviewTitle(year int) string {
	buf := bytes.NewBufferString("")

	yearDisp := color.New(color.FgHiGreen).Add(color.Underline)
	yearDisp.Fprintf(buf, "       🥳 %v 🥳 : Year In Review       \n", year)

	defer color.Unset()
	return buf.String()
}
