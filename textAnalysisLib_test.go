package textAnalysisLib

import (
	"testing"
	"fmt"
)
func TestRemovePunctuation(t *testing.T){
	testString := "This, is. a string!!! wi;th pun:ctuation{}"
	result := RemovePunctuation(testString)
//	fmt.Println(result)
	expectedResult := "This is a string with punctuation"
	if result != expectedResult {
		t.Error("RemovePunctuation \n=> expected\n => got")
	}
}

func TestReadListFromFile(t *testing.T){
	filePath := "./data/stopwords/test.txt"
	result, err := ReadListFromFile(filePath)
	if err != nil {
		t.Error("ReadListFromFile")
	}
	var expectedResult = []string {"a", "b", "c"}
	if result[0] != expectedResult[0] || result[1] != expectedResult[1] || result[2] != 
	expectedResult[2]{
		t.Error("ReadListFromFile, unexpected result")
	}
}

func TestRemoveStopWords(t *testing.T) {
	language := "pt"
	testString := "é certamente o texto de marcador de posição mais famoso"
	result, err := RemoveStopWords(testString, language)
	expectedResult := "certamente texto marcador posição famoso"
	if err != nil {
		t.Error(err)
	}

	if result != expectedResult {
		fmt.Println(result)
		t.Error("RemoveStopWords, unexpected result")
	}
}

func TestTokenize(t *testing.T) {
	language := "pt"
	testString := "Vamos para o carnaval"

	result := Tokenize(testString, language)

	expectedResult := []string{"vamos", "carnaval"}

	if result[0] != expectedResult[0] || len(result) != len(expectedResult) {
		fmt.Println(result)
		t.Error("Tokenize, unexpected result")
	}

}


func TestGenerateBOW(t *testing.T) {
	language := "pt"
	testStrings := []string{"Vamos para o carnaval", "Vamos para a festa junina"}
	threshold := 2


	result := GenerateBOW(testStrings, language, threshold)

	expectedResult := map[string]int {
		"vamos": 2,
		"festa": 1,
		"carnaval": 1,
	}

	for token, count := range result{
		expectedValue, ok := expectedResult[token]
		if !ok || expectedValue != count {
			fmt.Println(result)
			t.Error("GenerateBOW, unexpected result")
		}
	}

}
