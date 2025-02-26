package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
)

type OuterThing struct {
	Input  string   `json:"input"`
	Output OutThing `json:"output"`
}

type OutThing struct {
	MainTitle string  `json:"mainTitle"`
	Slides    []Slide `json:"slides"`
}

type Slide struct {
	SlideTitle string   `json:"slideTitle"`
	Content    []string `json:"content"`
}

func readPDF(filename string) (string, error) {
	f, r, err := pdf.Open(filename + ".pdf")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var text strings.Builder
	for i := 1; i <= r.NumPage(); i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}
		content, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		text.WriteString(content)
	}
	return text.String(), nil
}

func parseMarkdown(filename string) (OutThing, error) {
	data, err := os.ReadFile(filename + ".md")
	if err != nil {
		return OutThing{}, err
	}

	content := string(data)

	titleRegex := regexp.MustCompile(`# Main Title: (.*?)(\n|$)`)
	titleMatch := titleRegex.FindStringSubmatch(content)

	mainTitle := "Untitled"
	if len(titleMatch) > 1 {
		mainTitle = strings.TrimSpace(titleMatch[1])
	}

	slideRegex := regexp.MustCompile(`### \*\*Slide \d+: (.*?)\*\*\nContent:\n((?:.|\n)*?)(?:---|$)`)
	slideMatches := slideRegex.FindAllStringSubmatch(content, -1)

	var slides []Slide

	for _, slideMatch := range slideMatches {
		if len(slideMatch) > 2 {
			slideTitle := strings.TrimSpace(slideMatch[1])
			slideContent := slideMatch[2]

			var bulletPoints []string

			bulletLines := strings.SplitSeq(slideContent, "\n")
			for line := range bulletLines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "•") {
					bulletContent := strings.TrimSpace(strings.TrimPrefix(line, "•"))
					if bulletContent != "" {
						bulletPoints = append(bulletPoints, bulletContent)
					}
				}
			}

			slides = append(slides, Slide{
				SlideTitle: slideTitle,
				Content:    bulletPoints,
			})
		}
	}

	return OutThing{
		MainTitle: mainTitle,
		Slides:    slides,
	}, nil
}

func appendToJSON(data OuterThing) error {
	resFile := "res.json"

	var existingData []OuterThing

	if _, err := os.Stat(resFile); err == nil {
		fileData, err := os.ReadFile(resFile)
		if err != nil {
			return err
		}

		if len(fileData) > 0 {
			err = json.Unmarshal(fileData, &existingData)
			if err != nil {
				fmt.Println("Warning: res.json exists but is not valid JSON. Creating new file.")
				existingData = []OuterThing{}
			}
		}
	}

	existingData = append(existingData, data)

	jsonData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(resFile, jsonData, 0644)
}

func processFile(filename string) error {
	pdfText, err := readPDF(filename)
	if err != nil {
		return fmt.Errorf("error reading PDF: %v", err)
	}

	outputData, err := parseMarkdown(filename)
	if err != nil {
		return fmt.Errorf("couldn't parse markdown: %v", err)
	}

	jsonObject := OuterThing{
		Input:  pdfText,
		Output: outputData,
	}

	if err := appendToJSON(jsonObject); err != nil {
		return fmt.Errorf("couldn't append to JSON: %v", err)
	}

	fmt.Printf("We done processed %s\n", filename)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <filename>")
		os.Exit(1)
	}

	fname := os.Args[1]
	if strings.Contains(fname, ".") {
		fname = strings.Split(fname, ".")[0]
	}

	if err := processFile(fname); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
