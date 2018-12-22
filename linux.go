package main

import "fmt"

type linuxTerminal struct{}

func newLinuxTerminal() *linuxTerminal {
	return &linuxTerminal{}
}

func (l *linuxTerminal) Clear() {
	// Clear terminal screen
	fmt.Printf("\033[2J")
	// Move cursor to the top
	fmt.Printf("\033[1;1H")
}

func (l *linuxTerminal) Printf(c color, format string, a ...interface{}) (n int, err error) {
	s := fmt.Sprintf("%s%s", linuxColors[c], format)

	return fmt.Printf(s, a...)
}

func (l *linuxTerminal) Println(c color, a ...interface{}) (n int, err error) {
	fmt.Printf(linuxColors[c])
	return fmt.Println(a...)
}

var linuxColors = map[color]string{
	red:                       "\x1b[0;31m",
	red | bold:                "\x1b[1;31m",
	green:                     "\x1b[0;32m",
	green | bold:              "\x1b[1;32m",
	blue:                      "\x1b[0;34m",
	blue | bold:               "\x1b[1;34m",
	red | green:               "\x1b[0;33m",
	red | green | bold:        "\x1b[01;33m",
	blue | red:                "\x1b[0;35m",
	blue | red | bold:         "\x1b[1;35m",
	blue | green:              "\x1b[0;36m",
	blue | green | bold:       "\x1b[1;36m",
	red | blue | green:        "\x1b[0m",
	red | blue | green | bold: "\x1b[0m", //TODO
	0:    "\x1b[0m", //TODO
	bold: "\x1b[0m", //TODO
}
