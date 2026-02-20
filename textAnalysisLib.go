package textAnalysisLib
import (
	"strings"
	"os"
	"fmt"
	"math"
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


// Transforms a string in a list of clean normalized tokens
//
//- texts => []string;
//
//- language => string; two letter abbreviation ("pt", "en", "es"), following the ISO 639 standard.
//
//- threshold => int; the minimum number of occurences of a given token in order to be considered as part of the result.
func GenerateBOW(texts []string, language string, threshold int) map[string]int {

	bow := make(map[string]int)

	for _, text := range texts {
		tokens := Tokenize(text, language)
		for _, token := range tokens {
			val, ok := bow[token]
			if !ok {
				bow[token] = 1
			} else {
				bow[token] = val + 1
			}
		}
	}

	// apply the threshold
	result := make(map[string]int)

	for token, count := range bow {
		if count > threshold {
			result[token] = count
		}
	}

	return result
}


// Calculates the probability of a given string being found in a bag of words, which can be particularly useful when using a naive Bayes 
// classifier.
// 
//- test => string (the string for which the probability will be calculated)
// 
//- bow => map[string]int (the bag of words)
// 
//- language => string; two letter abbreviation ("pt", "en", "es"), following the ISO 639 standard.
// 
func CalculateTextBowProbability(text string, bow map[string]int, language string) float64 {
	tokens := Tokenize(text, language)
	probability := 0.0
	totalCounts := 0

	for _, count := range bow {
		totalCounts += count
	}

	
	denominator := float64(totalCounts + len(bow))

	for _, token := range tokens {
		count := bow[token]
		// Laplace smoothing for unseen words.
		sum := math.Log2(float64(count+1)) - math.Log2(denominator)
		probability += sum
	}
	return probability
}


// Calculate probability for diferent BOWs
func CalculateTextProbabilityForBOWs(text string, bows []map[string]int, language string) []float64 {
	probabilities := []float64{} 
	for _, bow := range bows {
		probability := CalculateTextBowProbability(text, bow, language)
		probabilities =	append(probabilities, probability)
	}
	return probabilities
}
