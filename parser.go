package main

import (
	"regexp"
	"strings"
)

var interfaceRe = regexp.MustCompile(`(?m)interface (.+) {\n((?:.+?\n)+)}`)
var sigRe = regexp.MustCompile(`(?:(?:(.+) extends (?:(.+)(?:, )?)+)|(.+))`)
var propRe = regexp.MustCompile(`    (?:(.+?)(?:<.+>)?\((.+)?\): (.+);)?(?:(?:(readonly) )?(.+?): (.+);)?`)
var funcRe = regexp.MustCompile(`(.+?): ((\((.+?)\) => (.+?))|(?:.+?)), `)

func parseCode(code string) map[string]Interface {
	out := make(map[string]Interface)
	interfaces := interfaceRe.FindAllStringSubmatch(code, -1)
	for _, intf := range interfaces {
		in := Interface{
			Methods:    make([]Method, 0),
			Properties: make([]Property, 0),
		}

		// Get signature
		sig := sigRe.FindAllStringSubmatch(intf[1], -1)
		if len(sig[0][1]) > 0 {
			in.Name = sig[0][1]
			in.Implements = strings.Split(sig[0][2], ", ")
		} else {
			in.Name = sig[0][3]
			in.Implements = make([]string, 0)
		}

		// Get properties
		props := propRe.FindAllStringSubmatch(intf[2], -1)
		fns := make(map[string]Method)

		for _, match := range props {
			// Is property
			if match[5] != "" {
				in.Properties = append(in.Properties, Property{
					Name:       match[5],
					Type:       match[6],
					IsReadonly: match[4] != "",
				})
			} else if match[1] != "" {
				// Is function, get params
				params := funcRe.FindAllStringSubmatch(match[2], -1)
				pars := make([]Parameter, len(params))
				i := 0
				for _, par := range params {
					typ := par[2]
					if par[5] != "" {
						typ = par[5]
					}
					pars[i] = Parameter{
						Name: par[1],
						Type: typ,
					}
					i++
				}

				fns[match[1]] = Method{
					Name:       match[1],
					Parameters: pars,
					ReturnType: match[3],
				}
			}
		}

		for _, fn := range fns {
			in.Methods = append(in.Methods, fn)
		}

		out[in.Name] = in
	}
	return out
}
