package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"math/rand"
	"os"
	"time"
)

func quiz(fds []*fInfo, term terminal, number int) {

	rand.Seed(time.Now().UTC().UnixNano())

	score := 0

	for i := 0; i < number; i++ {
		fmt.Printf("loop, id:%d\n", i)
		if len(fds) > 0 {
			id := rand.Intn(len(fds))
			fd := fds[id]

			term.Clear()

			fmt.Printf("score:%d/%d\n", score, i)

			fmt.Println("total:", len(fds))

			term.Printf(blue|bold, "%d/%d\n", i+1, number)

			fd.Display(term, pack|comments)

			// So that it can be processed as a valid go file.
			var buf = []byte("package quiz\n")

			fIndex := len(buf)

			if fd.recv.str == "" {
				buf = append(buf, "func "...)
			} else {
				buf = append(buf, "func "...)
				buf = append(buf, fd.recv.str...)
				buf = append(buf, " "...)
			}

			fmt.Printf(string(buf[fIndex:]))

			reader := bufio.NewReader(os.Stdin)
			brl, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println("Error reading line")
				continue
			}

			buf = append(buf, brl...)

			fset := token.NewFileSet() // positions are relative to fset
			node, err := parser.ParseFile(fset, "", buf, parser.ParseComments)
			if err != nil {
				fmt.Println("non valid function declaration")

				fd.Display(term, nocomments)

			} else {
				// Do not check if first char name is up (exported function)
				fi := inspect(node, []byte(buf), false)

				if fd.isEqual(fi[0]) {
					score++
				}

				fd.Display(term, nocomments)

				// TODO Windows ?
				reader.ReadString('\n')
			}
			term.Println(dft, "press enter")
			// TODO press any key to continue
			reader.ReadString('\n')

			// remove this element id for next step
			// TODO implement another learning algorithms to be chosen
			slice := make([]*fInfo, len(fds)-1)

			copy(slice, fds[0:id])
			copy(slice[id:], fds[id+1:])
			fds = slice
		}

	}

}
