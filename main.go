package main

import (
	"fmt"
	"strings"
)

const (
	// Text Styles
	Reset         = "\033[0m"
	Bold          = "\033[1m"
	Dim           = "\033[2m"
	Italic        = "\033[3m"
	Underline     = "\033[4m"
	Blink         = "\033[5m"
	Inverse       = "\033[7m"
	Hidden        = "\033[8m"
	Strikethrough = "\033[9m"

	// Foreground Colors
	FgBlack         = "\033[30m"
	FgRed           = "\033[31m"
	FgGreen         = "\033[32m"
	FgYellow        = "\033[33m"
	FgBlue          = "\033[34m"
	FgMagenta       = "\033[35m"
	FgCyan          = "\033[36m"
	FgWhite         = "\033[37m"
	FgBrightBlack   = "\033[90m"
	FgBrightRed     = "\033[91m"
	FgBrightGreen   = "\033[92m"
	FgBrightYellow  = "\033[93m"
	FgBrightBlue    = "\033[94m"
	FgBrightMagenta = "\033[95m"
	FgBrightCyan    = "\033[96m"
	FgBrightWhite   = "\033[97m"

	// Background Colors
	BgBlack         = "\033[40m"
	BgRed           = "\033[41m"
	BgGreen         = "\033[42m"
	BgYellow        = "\033[43m"
	BgBlue          = "\033[44m"
	BgMagenta       = "\033[45m"
	BgCyan          = "\033[46m"
	BgWhite         = "\033[47m"
	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"
)

type Peacock struct {
	color     string
	bgColor   string
	styles    []string
	text      string
	lastColor string
}

func (p *Peacock) Red(str ...string) *Peacock {
	p.color = FgRed

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "") + p.lastColor
	}

	return p
}

func (p *Peacock) BgRed(str ...string) *Peacock {
	p.bgColor = BgRed

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "")
	}

	return p
}

func (p *Peacock) Blue(str ...string) *Peacock {
	p.color = FgBlue

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "") + p.lastColor
	}

	return p
}

func (p *Peacock) Green(str ...string) *Peacock {
	p.color = FgGreen

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "") + p.lastColor
	}

	return p
}

func (p *Peacock) Bold(str ...string) *Peacock {
	p.styles = append(p.styles, Bold)

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "")
	}

	return p
}

func (p *Peacock) Dim(str ...string) *Peacock {
	p.styles = append(p.styles, Dim)

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "")
	}

	return p
}

func (p *Peacock) Italic(str ...string) *Peacock {
	p.styles = append(p.styles, Italic)

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "")
	}

	return p
}

func (p *Peacock) Underline(str ...string) *Peacock {
	p.styles = append(p.styles, Underline)

	trimmed := strings.Trim(strings.Join(str, ""), " ")

	if len(trimmed) > 0 {
		p.text = strings.Join(str, "")
	}

	return p
}

func (p *Peacock) Done() string {
	value := make([]string, 0)
	p.lastColor = p.color

	if len(p.styles) > 0 {
		value = append(value, p.styles...)
	}

	if len(p.color) > 0 {
		value = append(value, p.color)
	}

	if len(p.bgColor) > 0 {
		value = append(value, p.bgColor)
	}

	if len(p.text) > 0 {
		value = append(value, p.text)
	}

	return strings.Join(value, "")
}

func main() {
	p := &Peacock{}

	fmt.Println(p.Red("The quick brown fox jumps over the lazy dog").Done())
	fmt.Println(p.Green("The quick brown fox jumps over the lazy dog").Done())
	fmt.Println(p.Blue("The quick brown fox jumps over the lazy dog").Done())

	// === TEST ===
	// fmt.Println(p.Blue("testing").BgRed().Bold().Italic().Dim().Underline().Done())
	// fmt.Println(p.Green("I am a green line " + p.Red("and part of me is red").Done() + " `but i should be green` " + p.Blue("with a blue substring").Done() + " that becomes green again!").Done())
	// fmt.Println(FgGreen + "I am a green line " + FgBlue + "with a blue substring" + FgGreen + " that becomes green again!" + Reset);
}
