// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"os"
// // 	"regexp"
// // 	"strings"

// // 	"github.com/ledongthuc/pdf"
// // )

// // type OuterThing struct {
// // 	Input  string   `json:"input"`
// // 	Output OutThing `json:"output"`
// // }

// // type OutThing struct {
// // 	MainTitle string  `json:"mainTitle"`
// // 	Slides    []Slide `json:"slides"`
// // }

// // type Slide struct {
// // 	SlideTitle string   `json:"slideTitle"`
// // 	Content    []string `json:"content"`
// // }

// // func readPDF(filename string) (string, error) {
// // 	f, r, err := pdf.Open(filename + ".pdf")
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	defer f.Close()

// // 	var text strings.Builder
// // 	for i := 1; i <= r.NumPage(); i++ {
// // 		p := r.Page(i)
// // 		if p.V.IsNull() {
// // 			continue
// // 		}
// // 		content, err := p.GetPlainText(nil)
// // 		if err != nil {
// // 			return "", err
// // 		}
// // 		text.WriteString(content)
// // 	}
// // 	return text.String(), nil
// // }

// // func parseMarkdown(filename string) (OutThing, error) {
// // 	data, err := os.ReadFile(filename + ".md")
// // 	if err != nil {
// // 		return OutThing{}, err
// // 	}

// // 	content := string(data)

// // 	titleRegex := regexp.MustCompile(`# Main Title: (.*?)(\n|$)`)
// // 	titleMatch := titleRegex.FindStringSubmatch(content)

// // 	mainTitle := "Untitled"
// // 	if len(titleMatch) > 1 {
// // 		mainTitle = strings.TrimSpace(titleMatch[1])
// // 	}

// // 	slideRegex := regexp.MustCompile(`### \*\*Slide \d+: (.*?)\*\*\nContent:\n((?:.|\n)*?)(?:---|$)`)
// // 	slideMatches := slideRegex.FindAllStringSubmatch(content, -1)

// // 	var slides []Slide

// // 	for _, slideMatch := range slideMatches {
// // 		if len(slideMatch) > 2 {
// // 			slideTitle := strings.TrimSpace(slideMatch[1])
// // 			slideContent := slideMatch[2]

// // 			var bulletPoints []string

// // 			bulletLines := strings.SplitSeq(slideContent, "\n")
// // 			for line := range bulletLines {
// // 				line = strings.TrimSpace(line)
// // 				if strings.HasPrefix(line, "•") {
// // 					bulletContent := strings.TrimSpace(strings.TrimPrefix(line, "•"))
// // 					if bulletContent != "" {
// // 						bulletPoints = append(bulletPoints, bulletContent)
// // 					}
// // 				}
// // 			}

// // 			slides = append(slides, Slide{
// // 				SlideTitle: slideTitle,
// // 				Content:    bulletPoints,
// // 			})
// // 		}
// // 	}

// // 	return OutThing{
// // 		MainTitle: mainTitle,
// // 		Slides:    slides,
// // 	}, nil
// // }

// // func appendToJSON(data OuterThing) error {
// // 	resFile := "res.json"

// // 	var existingData []OuterThing

// // 	if _, err := os.Stat(resFile); err == nil {
// // 		fileData, err := os.ReadFile(resFile)
// // 		if err != nil {
// // 			return err
// // 		}

// // 		if len(fileData) > 0 {
// // 			err = json.Unmarshal(fileData, &existingData)
// // 			if err != nil {
// // 				fmt.Println("Warning: res.json exists but is not valid JSON. Creating new file.")
// // 				existingData = []OuterThing{}
// // 			}
// // 		}
// // 	}

// // 	existingData = append(existingData, data)

// // 	jsonData, err := json.MarshalIndent(existingData, "", "  ")
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return os.WriteFile(resFile, jsonData, 0644)
// // }

// // func processFile(filename string) error {
// // 	pdfText, err := readPDF(filename)
// // 	if err != nil {
// // 		return fmt.Errorf("error reading PDF: %v", err)
// // 	}

// // 	outputData, err := parseMarkdown(filename)
// // 	if err != nil {
// // 		return fmt.Errorf("couldn't parse markdown: %v", err)
// // 	}

// // 	jsonObject := OuterThing{
// // 		Input:  pdfText,
// // 		Output: outputData,
// // 	}

// // 	if err := appendToJSON(jsonObject); err != nil {
// // 		return fmt.Errorf("couldn't append to JSON: %v", err)
// // 	}

// // 	fmt.Printf("We done processed %s\n", filename)
// // 	return nil
// // }

// // func main() {
// // 	if len(os.Args) < 2 {
// // 		fmt.Println("Usage: go run script.go <filename>")
// // 		os.Exit(1)
// // 	}

// // 	fname := os.Args[1]
// // 	if strings.Contains(fname, ".") {
// // 		fname = strings.Split(fname, ".")[0]
// // 	}

// // 	if err := processFile(fname); err != nil {
// // 		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
// // 		os.Exit(1)
// // 	}
// // }


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strings"

// 	"github.com/ledongthuc/pdf"
// )

// type OuterThing struct {
// 	Input  string   `json:"input"`
// 	Output OutThing `json:"output"`
// }

// type OutThing struct {
// 	MainTitle string  `json:"mainTitle"`
// 	Slides    []Slide `json:"slides"`
// }

// type Slide struct {
// 	SlideTitle string   `json:"slideTitle"`
// 	Content    []string `json:"content"`
// }

// func readPDF(filename string) (string, error) {
// 	f, r, err := pdf.Open(filename + ".pdf")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()

// 	var text strings.Builder
// 	for i := 1; i <= r.NumPage(); i++ {
// 		p := r.Page(i)
// 		if p.V.IsNull() {
// 			continue
// 		}
// 		content, err := p.GetPlainText(nil)
// 		if err != nil {
// 			return "", err
// 		}
// 		text.WriteString(content)
// 	}
// 	return text.String(), nil
// }

// func parseMarkdown(filename string) (OutThing, error) {
// 	data, err := os.ReadFile(filename + ".md")
// 	if err != nil {
// 		return OutThing{}, err
// 	}

// 	content := string(data)
	
// 	// Debug: Print the file content
// 	fmt.Println("File content:", content)

// 	titleRegex := regexp.MustCompile(`# Main Title: (.*?)(\n|$)`)
// 	titleMatch := titleRegex.FindStringSubmatch(content)

// 	mainTitle := "Untitled"
// 	if len(titleMatch) > 1 {
// 		mainTitle = strings.TrimSpace(titleMatch[1])
// 		fmt.Println("Found main title:", mainTitle)
// 	} else {
// 		fmt.Println("No main title match found")
// 	}

// 	// Split content by slide sections
// 	sections := strings.Split(content, "---")
	
// 	var slides []Slide
	
// 	// Process each section
// 	for _, section := range sections {
// 		section = strings.TrimSpace(section)
// 		if section == "" {
// 			continue
// 		}
		
// 		// Extract slide title
// 		slideTitleRegex := regexp.MustCompile(`### \*\*Slide \d+: (.*?)\*\*`)
// 		titleMatch := slideTitleRegex.FindStringSubmatch(section)
		
// 		if len(titleMatch) < 2 {
// 			fmt.Println("No title match found in section:", section)
// 			continue
// 		}
		
// 		slideTitle := strings.TrimSpace(titleMatch[1])
// 		fmt.Println("Found slide title:", slideTitle)
		
// 		// Extract bullet points
// 		var bulletPoints []string
		
// 		// Split by lines
// 		lines := strings.Split(section, "\n")
// 		contentStarted := false
		
// 		for _, line := range lines {
// 			line = strings.TrimSpace(line)
			
// 			// Check if we've reached the content section
// 			if line == "Content:" {
// 				contentStarted = true
// 				continue
// 			}
			
// 			// Process bullet points only after "Content:" line
// 			if contentStarted && strings.HasPrefix(line, "-") {
// 				bulletContent := strings.TrimSpace(strings.TrimPrefix(line, "-"))
// 				if bulletContent != "" {
// 					bulletPoints = append(bulletPoints, bulletContent)
// 					fmt.Println("Added bullet point:", bulletContent)
// 				}
// 			}
// 		}
		
// 		if len(bulletPoints) > 0 {
// 			slides = append(slides, Slide{
// 				SlideTitle: slideTitle,
// 				Content:    bulletPoints,
// 			})
// 			fmt.Println("Added slide:", slideTitle, "with", len(bulletPoints), "bullet points")
// 		} else {
// 			fmt.Println("No bullet points found for slide:", slideTitle)
// 		}
// 	}

// 	fmt.Println("Total slides found:", len(slides))
	
// 	return OutThing{
// 		MainTitle: mainTitle,
// 		Slides:    slides,
// 	}, nil
// }

// func appendToJSON(data OuterThing) error {
// 	resFile := "res.json"

// 	var existingData []OuterThing

// 	if _, err := os.Stat(resFile); err == nil {
// 		fileData, err := os.ReadFile(resFile)
// 		if err != nil {
// 			return err
// 		}

// 		if len(fileData) > 0 {
// 			err = json.Unmarshal(fileData, &existingData)
// 			if err != nil {
// 				fmt.Println("Warning: res.json exists but is not valid JSON. Creating new file.")
// 				existingData = []OuterThing{}
// 			}
// 		}
// 	}

// 	existingData = append(existingData, data)

// 	jsonData, err := json.MarshalIndent(existingData, "", "  ")
// 	if err != nil {
// 		return err
// 	}

// 	return os.WriteFile(resFile, jsonData, 0644)
// }

// func processFile(filename string) error {
// 	pdfText, err := readPDF(filename)
// 	if err != nil {
// 		fmt.Printf("Warning: Could not read PDF: %v, continuing with markdown only\n", err)
// 		pdfText = "No PDF content available"
// 	}

// 	outputData, err := parseMarkdown(filename)
// 	if err != nil {
// 		return fmt.Errorf("couldn't parse markdown: %v", err)
// 	}

// 	jsonObject := OuterThing{
// 		Input:  pdfText,
// 		Output: outputData,
// 	}

// 	if err := appendToJSON(jsonObject); err != nil {
// 		return fmt.Errorf("couldn't append to JSON: %v", err)
// 	}

// 	fmt.Printf("We done processed %s\n", filename)
// 	return nil
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run script.go <filename>")
// 		os.Exit(1)
// 	}

// 	fname := os.Args[1]
// 	if strings.Contains(fname, ".") {
// 		fname = strings.Split(fname, ".")[0]
// 	}

// 	if err := processFile(fname); err != nil {
// 		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
// 		os.Exit(1)
// 	}
// }


package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

func readPDFWithPython(pdfPath string) (string, error) {
	// Create a temporary file name
	tempFile := "temp_pdf_text.txt"
	
	// Remove the temporary file if it exists
	os.Remove(tempFile)
	
	// Run the Python script to extract text to the temporary file
	cmd := exec.Command("python", "pdf_extractor.py", pdfPath+".pdf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run Python PDF extractor: %v\nOutput: %s", err, string(output))
	}
	
	// Check if the temporary file was created
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		return "", fmt.Errorf("Python script did not create the expected output file")
	}
	
	// Read the text from the temporary file
	content, err := os.ReadFile(tempFile)
	if err != nil {
		return "", fmt.Errorf("failed to read extracted text from temporary file: %v", err)
	}
	
	// Clean up the temporary file
	os.Remove(tempFile)
	
	return string(content), nil
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
		fmt.Println("Found main title:", mainTitle)
	} else {
		fmt.Println("No main title match found")
	}

	// Split content by slide sections
	sections := strings.Split(content, "---")
	
	var slides []Slide
	
	// Process each section
	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}
		
		// Extract slide title
		slideTitleRegex := regexp.MustCompile(`### \*\*Slide \d+: (.*?)\*\*`)
		titleMatch := slideTitleRegex.FindStringSubmatch(section)
		
		if len(titleMatch) < 2 {
			fmt.Println("No title match found in section:", section)
			continue
		}
		
		slideTitle := strings.TrimSpace(titleMatch[1])
		fmt.Println("Found slide title:", slideTitle)
		
		// Extract bullet points
		var bulletPoints []string
		
		// Split by lines
		lines := strings.Split(section, "\n")
		contentStarted := false
		
		for _, line := range lines {
			line = strings.TrimSpace(line)
			
			// Check if we've reached the content section
			if line == "Content:" {
				contentStarted = true
				continue
			}
			
			// Process bullet points only after "Content:" line
			if contentStarted && strings.HasPrefix(line, "-") {
				bulletContent := strings.TrimSpace(strings.TrimPrefix(line, "-"))
				if bulletContent != "" {
					bulletPoints = append(bulletPoints, bulletContent)
					fmt.Println("Added bullet point:", bulletContent)
				}
			}
		}
		
		if len(bulletPoints) > 0 {
			slides = append(slides, Slide{
				SlideTitle: slideTitle,
				Content:    bulletPoints,
			})
			fmt.Println("Added slide:", slideTitle, "with", len(bulletPoints), "bullet points")
		} else {
			fmt.Println("No bullet points found for slide:", slideTitle)
		}
	}

	fmt.Println("Total slides found:", len(slides))
	
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
	// First try with Go PDF library
	pdfText, err := readPDF(filename)
	if err != nil {
		fmt.Printf("Warning: Could not read PDF with Go library: %v\n", err)
		fmt.Printf("Trying Python extractor for %s.pdf...\n", filename)
		
		// Fall back to Python script
		pdfText, err = readPDFWithPython(filename)
		if err != nil {
			fmt.Printf("Warning: Python extractor also failed: %v\n", err)
			pdfText = "No PDF content available"
		} else {
			fmt.Printf("Successfully extracted PDF content using Python\n")
		}
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