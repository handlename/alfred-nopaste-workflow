package main

import (
	"fmt"
	"os"

	"github.com/handlename/alfred-nopaste-workflow"
)

func main() {
	client := wf.Client{
		Endpoint: os.Getenv("NOPASTE_ENDPOINT"),
		User:     os.Getenv("NOPASTE_USER"),
		Password: os.Getenv("NOPASTE_PASSWORD"),
	}

	r, err := client.Post(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to post: %s", err)
		os.Exit(1)
	}

	fmt.Print(r.String())
}
