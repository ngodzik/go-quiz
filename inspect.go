package main

import (
	"go/ast"
	"strings"
)

func inspect(node *ast.File, buf []byte, upCheck bool) []*fInfo {
	var fis []*fInfo

	ast.Inspect(node, func(n ast.Node) bool {
		if fd, ok := n.(*ast.FuncDecl); ok {

			if upCheck && (fd.Doc.Text() == "" || !isUpper(fd.Name.Name[0])) {
				return true
			}
			fi := &fInfo{pckg: node.Name.Name, name: strv{str: fd.Name.Name}}

			fi.comments = fd.Doc.Text()

			// Do not include "Deprecated" functions in the quiz
			if strings.Contains(fi.comments, "Deprecated") {
				return true
			}

			if fd.Recv != nil {
				fi.recv.str = string(buf[fd.Recv.Pos()-1 : fd.Recv.End()-1])
			}

			fi.pNames = make([][]strv, len(fd.Type.Params.List))
			for i, param := range fd.Type.Params.List {
				if param.Names != nil {
					for _, name := range fd.Type.Params.List[i].Names {
						fi.pNames[i] = append(fi.pNames[i], strv{str: name.Name})
					}
				}

				fi.pTypes = append(fi.pTypes, strv{str: string(buf[param.Type.Pos()-1 : param.Type.End()-1])})
			}

			if fd.Type.Results != nil {
				fi.rNames = make([][]strv, len(fd.Type.Results.List))
				for i, result := range fd.Type.Results.List {
					if result.Names != nil {
						for _, name := range fd.Type.Results.List[i].Names {
							fi.rNames[i] = append(fi.rNames[i], strv{str: name.Name})
						}
					}
					fi.rTypes = append(fi.rTypes, strv{str: string(buf[result.Type.Pos()-1 : result.Type.End()-1])})
				}
			}

			fis = append(fis, fi)
		}
		return true
	})

	return fis
}
