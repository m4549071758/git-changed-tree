package git

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func GetChangedFiles(base string, includeUncommitted bool) ([]string, error) {
	fileMap := make(map[string]bool)

	// git diff --name-only
	diffCmd := exec.Command("git", "diff", base, "--name-only")
	diffOut, err := diffCmd.Output()
	if err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(diffOut))
		for scanner.Scan() {
			path := strings.TrimSpace(scanner.Text())
			if path != "" {
				fileMap[path] = true
			}
		}
	}

	// git status --short (未コミットの変更)
	if includeUncommitted {
		statusCmd := exec.Command("git", "status", "--short")
		statusOut, err := statusCmd.Output()
		if err == nil {
			scanner := bufio.NewScanner(bytes.NewReader(statusOut))
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) > 3 {
					path := strings.TrimSpace(line[3:])
					// 改名などの場合は "old -> new" となるため考慮
					if idx := strings.Index(path, " -> "); idx != -1 {
						path = path[idx+4:]
					}
					fileMap[path] = true
				}
			}
		}
	}

	var files []string
	for f := range fileMap {
		files = append(files, f)
	}
	return files, nil
}
