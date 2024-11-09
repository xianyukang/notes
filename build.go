package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// not (Unicode Letter && Unicode Number && Space)
var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func main() {
	Shell("cp -r '/mnt/d/notes/Golang' '/mnt/d/project/notes/'")
	Shell("rm /mnt/d/project/notes/Golang/_*.md")
	GenerateTOC("/mnt/d/project/notes/**/*.md")
	GenerateFileLinks("/mnt/d/project/notes/**/*.md")
}

func GenerateTOC(pattern string) {
	files, err := filepath.Glob(pattern)
	CheckError(err)

	for _, path := range files {
		content, err := os.ReadFile(path)
		CheckError(err)
		toc := TOC(content)

		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		CheckError(err)

		_, err = file.WriteString(toc)
		CheckError(err)

		replacer := strings.NewReplacer("\t", "    ", "[➤ ", "#### [➤ ", "➤ ", "#### ➤ ")
		_, err = file.WriteString(replacer.Replace(string(content)))
		CheckError(err)

		err = file.Close()
		CheckError(err)
	}
}

func TOC(content []byte) string {
	lines := strings.Split(string(content), "\r\n")
	var result strings.Builder
	result.WriteString("## Table of Contents\r\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			title := ExtractLinkText(line[3:])
			result.WriteString(fmt.Sprintf("  - [%s](#%s)\r\n", title, hashtag(title)))
		}
		if strings.HasPrefix(line, "### ") {
			title := ExtractLinkText(line[4:])
			result.WriteString(fmt.Sprintf("    - [%s](#%s)\r\n", title, hashtag(title)))
		}
	}
	result.WriteString("\r\n")

	return result.String()
}

// ExtractLinkText extracts "aaa" from "[aaa](bbb)"
func ExtractLinkText(text string) string {
	var markdownLink = regexp.MustCompile(`\[(.+)\]\(.+\)`)
	result := markdownLink.FindStringSubmatch(text)
	if result != nil {
		return result[1]
	} else {
		return text
	}
}

func hashtag(str string) string {
	str = strings.TrimRight(str, " ")
	str = nonAlphanumericRegex.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, " ", "-")
	return url.PathEscape(str)
}

func Shell(shellCommand string) {
	cmd := exec.Command("fish", "-c", shellCommand)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateFileLinks(pattern string) {
	files, err := filepath.Glob(pattern)
	CheckError(err)

	var res strings.Builder
	for _, path := range files {
		ext := filepath.Ext(path)
		name := strings.TrimSuffix(filepath.Base(path), ext)
		dir := filepath.Base(filepath.Dir(path))
		res.WriteString(fmt.Sprintf("- [%s](%s/%s%s#table-of-contents)\r\n", name, url.PathEscape(dir), url.PathEscape(name), ext))
	}

	err = os.WriteFile("file_links.txt", []byte(res.String()), 0644)
	CheckError(err)
}
