package main

import (
	"flag"
	"fmt"
	"os"

	"git-changed-tree/git"
	"git-changed-tree/tree"
)

func main() {
	base := flag.String("base", "origin/main...HEAD", "Git base reference for diff")
	status := flag.Bool("status", true, "Include uncommitted changes (git status)")
	flag.Parse()

	files, err := git.GetChangedFiles(*base, *status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("No changes detected.")
		return
	}

	root := tree.BuildTree(files)
	tree.PrintTree(root, "", false, true)
}
