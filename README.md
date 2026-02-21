# package textAnalysisLib
  // import "github.com/mathealgou/text-analysis-lib"


## Functions

### func CalculateTextBowProbability(text string, bow map[string]int, language string) float64
    Calculates the probability of a given string being found in a bag of words,
    which can be particularly useful when using a naive Bayes classifier.

    - test => string (the string for which the probability will be calculated)

    - bow => map[string]int (the bag of words)

    - language => string; two letter abbreviation ("pt", "en", "es"), following
    the ISO 639 standard.

    returns float64 (The probability of the text in a given bow)

### func CalculateTextProbabilityForBOWs(text string, bows []map[string]int, language string) []float64
    Calculate probability for diferent BOWs

    Particularly useful for comparing the probabilities of a text being a part
    of separate categories of texts.

    - test => string (the string for which the probability will be calculated)

    - bows => []map[string]int (the bags of words)

    - language => string; two letter abbreviation ("pt", "en", "es"), following
    the ISO 639 standard.

    returns []float64 (The probabilities of the text for each bag of words
    given, in the same order.)

### func GenerateBOW(texts []string, language string, threshold int) map[string]int
    Transforms a string in a list of clean normalized tokens

    - texts => []string;

    - language => string; two letter abbreviation ("pt", "en", "es"), following
    the ISO 639 standard.

    - threshold => int; the minimum number of occurences of a given token in
    order to be considered as part of the result.

    returns map[string]int (Where the index is the token in question (see the
    Tokenize() function), and the int is the count of each token in the text.)

### func ReadCSV(filePath string, separator string) []map[string]string
    Reads and returns the contents of a CSV file.

    The file in question must contain headers in it's first line. This function
    WILL PANIC if for some reason it is unable to read the file.

    - filePath => string (In the same format as would be passed to os.ReadFile,
    `go doc os.ReadFile` for more information)

    - separator => string (The separator for the CSV file)

    returns []map[string]string (A list of the file's lines mapped by the names
    of their columns)

### func ReadListFromFile(filePath string) ([]string, error)
    Reads a text file and returns a list of its's lines.

    - filePath => string (In the same format as would be passed to os.ReadFile,
    `go doc os.ReadFile` for more information)

    returns ([]string, error) (A list of the file's lines)

### func RemovePunctuation(text string) string
    Does what it says on the tin. - text => string

    returns string (With punctuation characters removed)

### func RemoveStopWords(text string, language string) (string, error)
    Removes stop words from a given string.

    Stopwords are substrings with little semantic value, such as "a", "the",
    "of", "that", "what" in english, for example.

    - text => any string

    - language => two letter abbreviation ("pt", "en", "es"), following the ISO
    639 standard.

    returns (string, error) (A string is returned with the stop words removed,
    see /data/stopwords for information on which words are skipped in each
    language)

### func Tokenize(text string, language string) []string
    Transforms a string in a list of clean normalized tokens

    - text => any string

    - language => two letter abbreviation ("pt", "en", "es"), following the ISO
    639 standard.

    returns []string (a list of clean, usable tokens, with punctuation and
    stopwords removed.)

