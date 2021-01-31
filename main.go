package main

import (
	"fmt"
	"github.com/alameddinc/passwordBook-golang/cryptor"
	"github.com/alameddinc/passwordBook-golang/record"
	"github.com/manifoldco/promptui"
)

func main() {
	cry := cryptor.Cryptor{}
	cry.Load()

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Record", "Find", "Delete"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Record":
		cry.Encoded()
	case "Find":
		r := record.Record{}
		r.PromptURL()
		cry.Unload(r.URL)
	case "Delete":
		fmt.Println("Will add")
	}
}
