package main

import (
	github.com/kirves/go-form-it
)

func main() {
	form := BaseForm(POST, "action.html")
	
	tmpl.Execute(buf, map[string]interface{}{"form": form})
}