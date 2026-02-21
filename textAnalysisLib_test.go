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

func TestCalculateTextBowProbability(t *testing.T) {
	language := "pt"
	testStrings := []string{
		"Vamos para o carnaval",
		"Vamos para a festa junina",
		"essa festa no carnaval é ótima",
		"tem que acabar o carnaval",
		"festa em tal lugar no carnaval, bora?",
	}
	
	threshold := 1

	bow := GenerateBOW(testStrings, language, threshold)

	testText := "vamos ao carnaval"

	testText2 := "gostaria de ir para a biblioteca"
	

	result := CalculateTextBowProbability(testText, bow, language)

	result2 := CalculateTextBowProbability(testText2, bow, language)


	if result < result2 {
		t.Error("CalculateTextBowProbability, unexpected result.")
	}


}

func TestCalculateTextProbabilityForBOWs(t *testing.T) {
	language := "pt"
	threshold := 1
	testStrings := []string{
		"Vamos para o carnaval",
		"Vamos para a festa junina",
		"essa festa no carnaval é ótima",
		"tem que acabar o carnaval",
		"festa em tal lugar no carnaval, bora?",
	}
	testStrings2 := []string{
		"Vamos para a biblioteca",
		"Vamos para a aula de matemática",
		"essa biblioteca é ótima",
		"tem que acabar as bibliotecas",
		"aula de tal coisa na biblioteca, bora?",
	}

	bow1 := GenerateBOW(testStrings, language, threshold)

	bow2 := GenerateBOW(testStrings2, language, threshold)

	bows := []map[string]int {bow1, bow2}

	testText := "Essa é a biblioteca das bibliotecas"

	result := CalculateTextProbabilityForBOWs(testText, bows, language)

	// fmt.Println(result)

	if result[0] > result[1] {
		t.Error("CalculateTextProbabilityForBOWs, unexpected result")
	}

}

func TestReadCSV(t *testing.T) {
	result := ReadCSV("./data/csv/test.csv", ",")

	line := result[0]

	name := line["name"]

	if name != "joão" {
		t.Error("ReadCSV, unexpected result", name)
	}
}
