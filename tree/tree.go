package tree

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Name     string
	Children map[string]*Node
	IsFile   bool
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: make(map[string]*Node),
	}
}

func BuildTree(paths []string) *Node {
	root := NewNode(".")
	for _, path := range paths {
		if path == "" {
			continue
		}
		parts := strings.Split(path, "/")
		current := root
		for i, part := range parts {
			if _, ok := current.Children[part]; !ok {
				current.Children[part] = NewNode(part)
			}
			current = current.Children[part]
			if i == len(parts)-1 {
				current.IsFile = true
			}
		}
	}
	return root
}

func PrintTree(node *Node, indent string, isLast bool, isRoot bool) {
	if !isRoot {
		marker := "├── "
		if isLast {
			marker := "└── "
			fmt.Printf("%s%s%s\n", indent, marker, node.Name)
		} else {
			fmt.Printf("%s%s%s\n", indent, marker, node.Name)
		}
	} else {
		fmt.Println(node.Name)
	}

	newIndent := indent
	if !isRoot {
		if isLast {
			newIndent += "    "
		} else {
			newIndent += "━E  "
		}
	}

	keys := make([]string, 0, len(node.Children))
	for k := range node.Children {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		// チE��レクトリを優允E
		iDir := len(node.Children[keys[i]].Children) > 0
		jDir := len(node.Children[keys[j]].Children) > 0
		if iDir != jDir {
			return iDir
		}
		return keys[i] < keys[j]
	})

	for i, key := range keys {
		PrintTree(node.Children[key], newIndent, i == len(keys)-1, false)
	}
}
