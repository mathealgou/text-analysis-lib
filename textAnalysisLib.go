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
