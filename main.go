package main

//go:generate gopherjs build --output ./public/js/parseMarkdown.js --minify

import (
	"github.com/gomarkdown/markdown"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// to avoid name collision
	js.Global.Set("__custom_pkg__", map[string]interface{}{
		"parseMarkdown": ParseMarkdown,
	})
}

func ParseMarkdown(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	md := []byte(input)
	output := markdown.ToHTML(md, nil, nil)

	return string(output), nil
}
