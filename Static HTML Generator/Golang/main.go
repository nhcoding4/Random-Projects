package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
	TODO:
		List elements are borked when there is inline text
		Italic elements
		Need to add inline quotes (just add another delimiter)
*/

// --------------------------------------------------------------------------------------------------------------------

func main() {

	copyDestinationFolderName := "static"
	err := copyFiles(copyDestinationFolderName)
	if err != nil {
		log.Fatal(err)
	}

	const markdownLocation = "content/content.md"
	const templateLocation = "content/template.html"
	const writeIndexLocation = "static/index.html"
	err = generatePage(markdownLocation, templateLocation, writeIndexLocation)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// --------------------------------------------------------------------------------------------------------------------

// Takes a path of a text-file containing markdown, a template path and a destination path and creates a html file.

func generatePage(sourcePath, templatePath, destPath string) error {
	fmt.Println("Generating page from HTML")

	fileContents, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	html, err := textToHtml(string(fileContents))
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	title, err := extractTitleHtml(html)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	contentHtml, err := joinHtml(html)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	templateString := createHtmlString(templateBytes, title, contentHtml)

	htmlPage, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}
	defer htmlPage.Close()

	_, err = htmlPage.WriteString(templateString)
	if err != nil {
		return fmt.Errorf("generatePage: %v", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// Creates a html string.

func joinHtml(htmlStrings []string) (string, error) {
	if len(htmlStrings) == 0 {
		return "", errors.New("no html has been generated")
	}

	var htmlString string
	for _, html := range htmlStrings {
		htmlString += html
	}

	return htmlString, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Replaces template placeholders with HTML.

func createHtmlString(rawBytes []byte, title, contents string) string {
	templateString := string(rawBytes)
	titlePlaceHolder := "{{ Title }}"
	contentPlaceHolder := "{{ Content }}"
	templateString = strings.Replace(templateString, titlePlaceHolder, title, 1)
	templateString = strings.Replace(templateString, contentPlaceHolder, contents, 1)

	return templateString
}

// --------------------------------------------------------------------------------------------------------------------
