package textAnalysisLib

import (
	"strings"
	"os"
	"fmt"
)

// Does what it says on the tin.
func RemovePunctuation(text string) string{

	result := text

	punctuation, err := ReadListFromFile("./data/punctuation.txt")
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	for _, p := range punctuation {
		result = strings.ReplaceAll(result, p, "")
	}
	return result
}


// Reads a text file and returns a list of its's lines.
//
func ReadListFromFile(filePath string) ([]string, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	textContent := string(contents)
	result := strings.Split(textContent, "\n")
	return result, nil
}


// Removes stop words from a given string.
//
// Stopwords are substrings with little semantic value, such as "a", "the", "of", "that", 
// "what" in english, for example.
//
//- text => any string
//
//- language => two letter abbreviation ("pt", "en", "es"), following the ISO 639 standard.
//
func RemoveStopWords(text string, language string) (string, error) {
	stopwordsFilePath := fmt.Sprintf("./data/stopwords/%v.txt", language)
	stopwords, err := ReadListFromFile(stopwordsFilePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// this is needed in case the stopword is at the very start of the text
	result := fmt.Sprintf(" %v", text)

	for _, word := range stopwords {
		// need to add an extra space before the word, for some of them will
		// be substrings of other words
		wordWithSpace := fmt.Sprintf(" %v ", word)
		result = strings.ReplaceAll(result, wordWithSpace, " ")
	}
	// also remove double spaces
	result = strings.ReplaceAll(result, "  ", " ")
	// also, trim
	result = strings.TrimSpace(result)
	return result, nil
}


// Transforms a string in a list of clean normalized tokens
//
//- text => any string
//
//- language => two letter abbreviation ("pt", "en", "es"), following the ISO 639 standard.
//
func Tokenize(text string, language string) []string {
	lowercaseText := strings.ToLower(text)

	cleanText, err := RemoveStopWords(lowercaseText, language)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	textWithoutPunctuation := RemovePunctuation(cleanText)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	tokens := strings.Fields(textWithoutPunctuation)

	return tokens
}
