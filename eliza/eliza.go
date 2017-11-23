package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//create struct of responses
type Response struct {
	Patterns *regexp.Regexp
	Answers  []string
}

//this function creates elizas response
func makeResponses(path string) []Response {
	fullFile, _ := ReadLines(path)
	responses := make([]Response, 0)
	for i := 0; i < len(fullFile); i += 2 {
		allPatterns := strings.Split(fullFile[i], ";")
		allResponses := strings.Split(fullFile[i+1], ";")
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern
			Patterns := regexp.MustCompile(pattern)
			responses = append(responses, Response{Patterns: Patterns, Answers: allResponses})
		}
	}
	return responses
}

//test if responses are being populated
func PrintResponses(path string) {
	response := makeResponses(path)
	fmt.Printf("%+v\n", response)
}

func subWords(inputStr string) string {
	// split inputStr into slice of strings
	splitStr := strings.Fields(inputStr)

	//map of reflected words
	words := map[string]string{
		"am":     "are",
		"was":    "were",
		"i":      "you",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "me",
		"me":     "you",
	}

	// swap words
	for index, word := range splitStr {
		if value, ok := words[strings.ToLower(word)]; ok {
			splitStr[index] = value
		}
	}

	return strings.Join(splitStr, " ")
}

//read lines from file
//adapted from https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readLine := scanner.Text()

		if skipComment(readLine) {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func skipComment(readLine string) bool {
	return strings.HasPrefix(readLine, "//") || len(strings.TrimSpace(readLine)) == 0
}

func replaceWords(pattern *regexp.Regexp, input string) string {
	match := pattern.FindStringSubmatch(input)
	if len(match) == 1 {
		return "" // no capture is needed
	}
	replaceWord := match[1]
	return replaceWord
}

func responseBuilder(response, replaceWord string) string {
	if strings.Contains(response, "%s") {
		return fmt.Sprintf(response, replaceWord)
	}
	return response
}

func AskEliza(input string) string {
	//create response[]response from file
	response := makeResponses("./data/responses.dat")
	rand.Seed(time.Now().Unix())

	for _, response := range response {
		if response.Patterns.MatchString(input) {
			replaceWord := replaceWords(response.Patterns, input)
			genResp := response.Answers[rand.Intn(len(response.Answers))]
			genResp = responseBuilder(genResp, replaceWord)
			return genResp
		}
	}

	return "Why do you say that?"
}
