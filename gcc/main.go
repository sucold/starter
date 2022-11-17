package main

import "github.com/hinego/starter/gcc/ge"

func main() {
	var data = ge.APP{
		"test": {
			API: "v2",
			Model: map[string]*ge.Model{
				"User": {
					Tags: "用户",
					Actions: map[string]*ge.Action{
						"Fetch":  {},
						"Delete": {},
						"Create": {},
						"Update": {},
						"Get":    {},
						"Test":   {},
					},
				},
			},
			Logic: map[string]*ge.Model{
				"User": {},
			},
		},
	}
	data.Init()
}
