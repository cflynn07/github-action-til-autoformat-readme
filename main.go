package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var repoPath = os.Getenv("REPO_PATH")
var templatePath = os.Getenv("TEMPLATE_PATH")
var inputDescription = os.Getenv("INPUT_DESCRIPTION")
var inputFooter = os.Getenv("INPUT_FOOTER")

var re = regexp.MustCompile(`^Date:\s*`)
var re2 = regexp.MustCompile(`^#\s*`)

type Til struct {
	Title     string
	Filename  string
	DateAdded string
}

// run a git cli command, the capture and parse the output to extact the date
// a file was added to the repository
func cmdGetDate(file string) string {
	c1 := exec.Command("git", "log", "--diff-filter=A", "--", file)
	c1.Dir = repoPath
	var commandOutput bytes.Buffer
	c1.Stdout = &commandOutput

	err := c1.Start()
	if err != nil {
		log.Panic(err)
	}

	err = c1.Wait()
	if err != nil {
		log.Panic(err)
	}

	date := ""
	for _, outputLine := range strings.Split(commandOutput.String(), "\n") {
		if re.MatchString(outputLine) {
			// strip "Date: " substring from matching line
			date = re.ReplaceAllString(outputLine, "")
			break
		}
	}

	return date
}

func main() {
	// map of all categories and respective TILs
	tilsMap := make(map[string][]Til)
	// tils = TIL markdown files
	tils, _ := filepath.Glob(repoPath + "/**/*.md")

	for _, til := range tils {
		// grab the "category" and the "file"
		// ex: html/div-tags.md -- category "html" file "div-tags.md"
		splitResult := strings.Split(til, "/")
		length := len(splitResult)
		category := splitResult[length-2]
		file := splitResult[length-1]

		if strings.ToLower(file) == "readme.md" {
			continue
		}

		// Read the first line of each file, use the string as a title
		f, err := os.Open(til)
		if err != nil {
			log.Panic(err)
		}
		reader := bufio.NewReader(f)
		linkTitle, err := reader.ReadString('\n')
		if err != nil {
			log.Println(fmt.Sprintf("ERROR: file \"%s\" does not have > 1 line of text (no title)", file))
			log.Panic(err)
		}

		// strip "# " from beginning of line
		linkTitle = re2.ReplaceAllString(linkTitle, "")
		linkTitle = strings.TrimSpace(linkTitle)

		// if category first encountered in loop so far, append new map key, otherwise
		// add to existing
		if _, exists := tilsMap[category]; exists {
			tilsMap[category] = append(tilsMap[category], Til{
				Title:     linkTitle,
				Filename:  file,
				DateAdded: cmdGetDate(til),
			})
		} else {
			tilsMap[category] = []Til{
				Til{
					Title:     linkTitle,
					Filename:  file,
					DateAdded: cmdGetDate(til),
				},
			}
		}
	}

	// load and execute template, write results to README.md
	t, err := template.New(path.Base(templatePath)).ParseFiles(templatePath)
	if err != nil {
		log.Panic(err)
	}

	var output bytes.Buffer
	err = t.Execute(&output, struct {
		Tils             map[string][]Til
		AllTils          []string
		InputDescription string
		InputFooter      string
	}{
		Tils:             tilsMap,
		AllTils:          tils,
		InputDescription: inputDescription,
		InputFooter:      inputFooter,
	})

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
	fmt.Print(output.String())
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")

	// truncates before writing
	ioutil.WriteFile(repoPath+"/README.md", []byte(output.String()), 0644)
}
