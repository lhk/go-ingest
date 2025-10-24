package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	clipboardFlag := flag.Bool("clipboard", false, "Copy output to clipboard")
	flag.Parse()
	patterns := flag.Args()

	if len(patterns) == 0 {
		fmt.Println("Usage: git-ingest-local [--clipboard] <glob patterns>")
		os.Exit(1)
	}

	// Collect files
	files := collectFiles(patterns)
	if len(files) == 0 {
		fmt.Println("No files matched the provided patterns.")
		return
	}

	// Build folder structure
	tree := buildTree(files)
	var sb strings.Builder
	sb.WriteString("# Folder structure\n")
	for _, line := range renderTree(tree, "", true) {
		sb.WriteString(line + "\n")
	}
	sb.WriteString("\n")

	// Append file contents
	for _, f := range files {
		sb.WriteString("# " + f + "\n")
		data, err := os.ReadFile(f)
		if err != nil {
			sb.WriteString(fmt.Sprintf("# [Error reading file: %v]\n", err))
		} else {
			sb.WriteString(string(data) + "\n")
		}
		sb.WriteString("\n")
	}

	output := sb.String()
	fmt.Print(output)

	// Copy to clipboard if requested
	if *clipboardFlag {
		if err := clipboard.Init(); err == nil {
			clipboard.Write(clipboard.FmtText, []byte(output))
			fmt.Println("\nOutput copied to clipboard.")
		} else {
			fmt.Printf("\nCould not initialize clipboard: %v\n", err)
		}
	}
}

// collectFiles expands glob patterns into unique, sorted file paths.
func collectFiles(patterns []string) []string {
	seen := make(map[string]bool)
	for _, pat := range patterns {
		matches, _ := filepath.Glob(pat)
		for _, m := range matches {
			info, err := os.Stat(m)
			if err == nil && !info.IsDir() {
				seen[m] = true
			}
		}
	}
	files := make([]string, 0, len(seen))
	for f := range seen {
		files = append(files, f)
	}
	sort.Strings(files)
	return files
}

// node represents a tree node.
type node struct {
	children map[string]*node
	isFile   bool
}

// buildTree constructs a nested folder tree from file paths.
func buildTree(files []string) *node {
	root := &node{children: make(map[string]*node)}
	for _, f := range files {
		parts := strings.Split(filepath.ToSlash(f), "/")
		curr := root
		for i, p := range parts {
			if curr.children == nil {
				curr.children = make(map[string]*node)
			}
			if _, ok := curr.children[p]; !ok {
				curr.children[p] = &node{children: make(map[string]*node)}
			}
			curr = curr.children[p]
			if i == len(parts)-1 {
				curr.isFile = true
			}
		}
	}
	return root
}

// renderTree returns a slice of lines representing the folder tree.
func renderTree(n *node, prefix string, root bool) []string {
	var lines []string
	keys := make([]string, 0, len(n.children))
	for k := range n.children {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		child := n.children[k]
		connector := "├── "
		ext := "│   "
		if i == len(keys)-1 {
			connector = "└── "
			ext = "    "
		}
		if root {
			lines = append(lines, connector+k)
		} else {
			lines = append(lines, prefix+connector+k)
		}
		if len(child.children) > 0 {
			lines = append(lines, renderTree(child, prefix+ext, false)...)
		}
	}
	return lines
}
