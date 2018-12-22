package main

type color uint

const (
	red color = 1 << iota
	green
	blue
	bold

	// default terminal color
	dft color = red | green | blue
)

type terminal interface {
	Clear()
	Printf(c color, format string, a ...interface{}) (n int, err error)
	Println(c color, a ...interface{}) (n int, err error)
}
