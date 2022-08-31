/*
	Program for

- retrieving all the topics from the tutorial gobyexample,
- retrieving their html via http request,
- editing and checking it line by line and depending on whether it is code or comment
- writing the content to a go-file.
In the end all 80 topics are crawled and the comments and lines of code are written to an executable go file.
Program was written to make use of the knowledge gained whithin the tutorial especially
- the basics
- methods
- string formatting
- regular expressions
- errors
- goroutines
- waitgroups
Cool thing:
- by using goroutines the execution time is reduced from ~30 seconds (not using goroutines in the main loop) to ~1.3 seconds
What I would like to have used but didn't
- interfaces!
- channels!
- case and select!
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var BASEPATH = "https://gobyexample.com"
var TOPICS = make(map[string]string)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	// for later getting the execution time of the program
	start := time.Now()

	// get the scanner that is able to read the document line by line (in this case the main page of gobyexample that lists all of the topics)
	scanner, resp := getScanner(BASEPATH)
	// body will be closed after the scanner is done reading
	defer resp.Body.Close()

	counter := 1
	// iterate through all lines of the scanned text
	for i := 0; scanner.Scan() && scanner.Text() != "</html>"; i++ {
		lineContent := scanner.Text()

		// <li> signals that a list entry and in the case of gobyexample a topic entry is following
		if strings.Contains(lineContent, "<li>") {

			// to extract the topic
			topic := strings.SplitN(lineContent, "\"", 3)[1]

			// to make the prefix with leading zeros so it will be in order when the subdirectories are created
			prefix := fmt.Sprintf("%03d", counter)

			// the global map is updated with the new found topic as value and the prefix as key
			TOPICS[prefix] = topic

			counter++
		}
	}

	// iterate through all the topics with their keys in the map TOPICS
	for k, v := range TOPICS {
		// add a worker to the workgroup
		wg.Add(1)

		fmt.Println(k, v)

		// create the subdirectory for the topic
		createTopic(v, k)

		// here the magic happens: using goroutines (that run in parallel) and multiple workers the execution time of the program is reduced from ~30 seconds to ~1.4 seconds
		// the anonymous function that will be run by individual goroutines takes k and v as arguments (otherwise they couldn't be used in there) and performs the heavy-lifting of the program
		// by calling the creation of the file for the topic and the retrieving the topic content and writing it to the respective file.
		go func(k string, v string) {
			fmt.Printf("Worker %s starting\n", k)
			defer wg.Done()
			f := createFile(k)
			defer f.Close()
			writeTopicContentToFile(f, v)
			fmt.Printf("Worker %s done\n", k)
		}(k, v) // this needs to be added (return value?!)
	}

	// waiting for the workgroups / goroutines to finish
	wg.Wait()

	// not sure if i need this here..
	check(scanner.Err())

	// print the execution time of the program to the terminal
	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}

// check checks whether the given error should cause panic.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// replaceHTMLCode takes the line of code crawled from the gobyexample.com subdomain and replaces the HTML codes extracted from the HTML into proper symbols.
func replaceHTMLCode(s string) string {
	s = strings.ReplaceAll(s, "&#34;", "\"")
	s = strings.ReplaceAll(s, "&lt;-", "<-")
	s = strings.ReplaceAll(s, "&#39;", "'")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "& lt;", "<")
	s = strings.ReplaceAll(s, "&amp;", "&")

	return s
}

// getScanner takes a topic (string) and return the *bufio.Scanner and the http.Response object for the given topic (on gobyexample.com).
func getScanner(t string) (*bufio.Scanner, *http.Response) {
	resp, err := http.Get(t)
	check(err)

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	return scanner, resp
}

// createTopic takes the given subtopic (t) from gobyexample.com and a prefix which is a number with leading 0's and creates a subdirectory for the topic in the current working directory.
func createTopic(t string, pre string) {
	dirpath := pre + "-" + t
	os.Mkdir(dirpath, 0755)
}

// createFile takes the prefix-number (k) and creates a go-script-file in the topic-subdirectory with the name of the topic and the correct extension. For subsequent writing to this file it is returned as f by this function and later closed.
func createFile(k string) *os.File {
	fullPath := filepath.Join(k+"-"+TOPICS[k], TOPICS[k]+".go")
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	return f
}

// checkPassageEnd return whether a block of important content (code or documentation) ends in the given line of code.
func checkPassageEnd(l string) bool {
	return strings.Contains(l, "</td>")
}

// checkDocumentEnd returns whether the relevant part of the document ends in the given line of code.
func checkDocumentEnd(l string) bool {
	return strings.Contains(l, "</table>")
}

// checkCodeCommentStart return the booleans code and doc which signify whether the given passage/line contains code or comment.
func checkCodeCommentStart(l string, code bool, doc bool) (bool, bool) {
	if strings.Contains(l, "<td class=\"docs\">") && !doc {
		doc = true
	} else if strings.Contains(l, "<td class=\"code\">") && !code {
		code = true
	} else if strings.Contains(l, "<td class=\"code leading\">") && !code {
		code = true
	}
	return code, doc
}

// writeTopicContentToFile takes the opened file and the topic as arguments and then creates a scanner that for the response of an http request made for the given topic.
// The scanner is then iterates through all the lines of the document.
func writeTopicContentToFile(f *os.File, t string) {
	dw := bufio.NewWriter(f)
	defer dw.Flush()

	scanner, resp := getScanner(BASEPATH + "/" + t)
	defer resp.Body.Close()
	docActive := false
	codeActive := false
	for i := 0; scanner.Scan(); i++ {
		lineContent := scanner.Text()

		if checkDocumentEnd(lineContent) {
			break
		}

		if checkPassageEnd(lineContent) {
			if codeActive {
				codeActive = !codeActive
			} else if docActive {
				docActive = !docActive
			}
		}

		if codeActive {
			writeCode(dw, lineContent)
		} else if docActive {
			writeComment(dw, lineContent)
		}

		codeActive, docActive = checkCodeCommentStart(lineContent, codeActive, docActive)
	}
}

// writeComment takes the datawriter (*bufio.Writer object) and the comment-line extracted from the gobyexample topic page and writes it to the file that the writer is hooked up to.
func writeComment(dw *bufio.Writer, comment string) {
	dw.WriteString("//" + comment + "\n")
}

// writeCode takes the datawriter (*bufio.Writer object) and the code-line extracted from the gobyexample topic page and writes it to the file that the writer is hooked up to.
// Before writing to the file writeCode does a few checks and alterations to the code so that it the whole go-file results in an executable program.
func writeCode(dw *bufio.Writer, code string) {
	rStart, _ := regexp.Compile("span class=")
	rEnd, _ := regexp.Compile("</span>")
	startIdxs := rStart.FindAllStringIndex(code, -1)
	endIdxs := rEnd.FindAllStringIndex(code, -1)

	var slices []string

	if len(endIdxs) > len(startIdxs) {
		endIdxs = endIdxs[len(endIdxs)-len(startIdxs):]
	} else if len(endIdxs) < len(startIdxs) {
		endIdxs = append(endIdxs, []int{len(code), len(code)})
	}

	if len(startIdxs) != 0 && endIdxs[0][0] < startIdxs[0][1] {
		endIdxs[0][0] = len(code)
	}

	for i := range startIdxs {
		slices = append(slices, strings.Split(string(code[startIdxs[i][1]:endIdxs[i][0]]), "\">")[1])
	}

	line := strings.Join(slices, " ")

	line = replaceHTMLCode(line)

	dw.WriteString(line + "\n")
}
