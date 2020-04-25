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
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var repoPath = os.Getenv("REPO_PATH")
var templatePath = os.Getenv("TEMPLATE_PATH")
var inputDescription = os.Getenv("INPUT_DESCRIPTION")
var inputFooter = os.Getenv("INPUT_FOOTER")
var inputListMostRecent = os.Getenv("INPUT_LIST_MOST_RECENT")

var re = regexp.MustCompile(`^Date:\s*`)
var re2 = regexp.MustCompile(`^#\s*`)

// I have no idea what RFC standard this may be (if any) but it's the format git
// seems to use to record datetime
const timeFormat = "Mon Jan 02 15:04:05 2006 -0700"

type Til struct {
	Title     string
	Filename  string
	Category  string
	DateAdded string
}

// sort TILs by DateAdded (DESC) and return n most recent
func cmdTrimMostRecentTils(tils *[]Til, n int) {
	if n <= 0 {
		n = 0
	}
	if n > len(*tils) {
		n = len(*tils)
	}
	sort.Slice(*tils, func(i, j int) bool {
		iTime, _ := time.Parse(timeFormat, (*tils)[i].DateAdded)
		jTime, _ := time.Parse(timeFormat, (*tils)[j].DateAdded)
		return iTime.After(jTime)
	})
	*tils = (*tils)[0:n]
}

// run a git cli command, the capture and parse the output to extact the date
// a file was added to the repository
func cmdGetDate(file *string) string {
	c1 := exec.Command("git", "log", "--diff-filter=A", "--", *file)
	c1.Dir = path.Dir(templatePath)
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
	// list of all (non-grouped by category) TILs for use with `list_most_recent` feature
	var tilsSlice []Til
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
		tilStruct := Til{
			Title:     linkTitle,
			Filename:  file,
			Category:  category,
			DateAdded: cmdGetDate(&til),
		}

		if _, exists := tilsMap[category]; exists {
			tilsMap[category] = append(tilsMap[category], tilStruct)
		} else {
			tilsMap[category] = []Til{tilStruct}
		}

		tilsSlice = append(tilsSlice, tilStruct)
	}

	n, err := strconv.Atoi(inputListMostRecent)
	if err != nil {
		n = 0
	}

	cmdTrimMostRecentTils(&tilsSlice, n)

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
		MostRecentTils   []Til
	}{
		Tils:             tilsMap,
		AllTils:          tils,
		InputDescription: inputDescription,
		InputFooter:      inputFooter,
		MostRecentTils:   tilsSlice,
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
