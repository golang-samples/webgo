package main

import (
    "github.com/hoisie/web"
)

func main() {
	web.Get("/", func() string {
		return "Hello Go言語!"
	})
	web.Run(":8080")
}
