package main

import (
	"flag"
	"fmt"

	"dinodb/pkg/config"
	"dinodb/pkg/list"
	"dinodb/pkg/repl"

	"github.com/google/uuid"
)

// Default port 8335 (BEES).
const DEFAULT_PORT int = 8335

const LOG_FILE_NAME = "data/dinodb.log"

// Start the database.
func main() {
	// Set up flags.
	promptFlag := flag.Bool("c", true, "use prompt?")
	projectFlag := flag.String("project", "", "choose project: [go,pager,hash,b+tree,join,concurrency,recovery] (required)")

	flag.Parse()

	// Set up REPL resources.
	prompt := config.GetPrompt(*promptFlag)
	repls := make([]*repl.REPL, 0)

	// Get the right REPLs.
	switch *projectFlag {
	case "go":
		l := list.NewList[list.Str]()
		repls = append(repls, list.ListRepl(l))

	default:
		fmt.Println("must specify -project [go,pager,hash,b+tree,join,concurrency,recovery]")
		return
	}

	// Combine the REPLs.
	r, err := repl.CombineRepls(repls)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run the REPL.
	r.Run(uuid.New(), prompt, nil, nil)
}
