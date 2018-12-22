package main

import "fmt"

type strv struct {
	str string
	// invalid receiver, default initializes to false: valid
	invalid bool
}

type fInfo struct {
	comments string
	pckg     string
	recv     strv
	name     strv

	// Parameters
	pNames [][]strv
	pTypes []strv
	// invalid parameters number
	ivpNb bool

	// Return values
	rNames [][]strv
	rTypes []strv
}

type display uint

const (
	comments display = 1 << iota
	pack
	recv
	name
	pNames
	pTypes
	rNames
	rTypes

	all        display = pack | comments | recv | name | pNames | pTypes | rNames | rTypes
	nocomments display = recv | name | pNames | pTypes | rNames | rTypes
)

func getColor(check bool, trueColor color, falseColor color) color {
	if check {
		return trueColor
	}
	return falseColor
}

func (f *fInfo) Display(term terminal, disp display) {

	if disp&pack != 0 {
		term.Printf(blue|bold, "package: %s\n", f.pckg)
	} else {
		fmt.Println()
	}
	if disp&comments != 0 {
		// TODO check comments and display "Deprecated" line in red bold
		term.Printf(dft, "%s\n", f.comments)
	}

	if disp&name != 0 {
		term.Printf(green, "func ")
	}

	if disp&recv != 0 {
		if f.recv.str != "" {
			term.Printf(green, "")
			term.Printf(getColor(f.recv.invalid, red, green), "%s ", f.recv.str)
		}
	}

	if disp&name != 0 {
		term.Printf(getColor(f.name.invalid, red|bold, green), "%s(", f.name.str)
	}

	for i, param := range f.pTypes {
		if disp&pNames != 0 {
			for n, name := range f.pNames[i] {
				term.Printf(getColor(name.invalid, red|green, green), "%s", name.str)

				if n != len(f.pNames[i])-1 {
					fmt.Printf(", ")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		if disp&pTypes != 0 {
			term.Printf(getColor(f.pTypes[i].invalid, red|bold, green), "%s", param.str)

			if i != len(f.pTypes)-1 {
				fmt.Printf(", ")
			}
		}
	}

	if disp&name != 0 {
		term.Printf(getColor(f.ivpNb, red|bold, green), ") ")
	}

	if len(f.rTypes) > 0 {
		if disp&rTypes != 0 {
			if len(f.rTypes) > 1 {
				fmt.Printf("(")
			}
		}
		for i, result := range f.rTypes {
			if disp&rNames != 0 {
				for n, name := range f.rNames[i] {
					term.Printf(getColor(name.invalid, red|green, green), "%s ", name.str)
					if n != len(f.rNames[i])-1 {
						fmt.Printf(", ")
					} else {
						fmt.Printf(" ")
					}

				}
			}
			if disp&rTypes != 0 {
				term.Printf(getColor(f.rTypes[i].invalid, red|bold, green), "%s", result.str)

				if i != len(f.rTypes)-1 {
					fmt.Printf(", ")
				}
			}
		}
		if disp&rTypes != 0 {
			if len(f.rTypes) > 1 {
				fmt.Printf(")")
			}
		}
	}

	fmt.Println()
	fmt.Println()

}

func (f *fInfo) isEqual(src *fInfo) bool {

	isEqual := true

	if src != nil {
		if f.recv != src.recv {
			f.recv.invalid = true
			isEqual = false
		}

		if f.name != src.name {
			f.name.invalid = true
			isEqual = false
		}

		for i, param := range f.pTypes {
			for n, name := range f.pNames[i] {
				if len(src.pNames) > i && len(src.pNames[i]) > n && src.pNames[i][n] != name {
					f.pNames[i][n].invalid = true
				} else if len(src.pNames)-1 < i {
					f.pNames[i][n].invalid = true
				} else if len(src.pNames) > i && len(src.pNames[i])-1 < n {
					f.pNames[i][n].invalid = true
				}
			}

			if (len(src.pTypes) > i && param != src.pTypes[i]) || len(src.pTypes)-1 < i {
				f.pTypes[i].invalid = true
				isEqual = false
			}
		}

		if len(f.pTypes) > 0 && len(src.pTypes) == 0 {
			f.ivpNb = true
			isEqual = false
		}

		if len(f.rTypes) > 0 {
			for i, result := range f.rTypes {
				for n, name := range f.rNames[i] {
					if len(src.rNames) > i && len(src.rNames[i]) > n && src.rNames[i][n] != name {
						f.rNames[i][n].invalid = true
					} else if len(src.rNames)-1 < i {
						f.rNames[i][n].invalid = true
					} else if len(src.rNames) > i && len(src.rNames[i])-1 < n {
						f.rNames[i][n].invalid = true
					}
				}

				if (len(src.rTypes) > i && result != src.rTypes[i]) || len(src.rTypes)-1 < i {
					f.rTypes[i].invalid = true
					isEqual = false
				}

			}
		}
	}

	return isEqual
}
