package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("What is your name?")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Wow,", input, "is a cool name! Let's break it down by letter.")
	// downcase the input so we don't have to worry about capitalization throwing a wrench in things.
	var sterilizedInput = strings.ToLower(input)

	printChars(sterilizedInput)

	fmt.Println("I hope you learned some cool things about your name!")
}

func printChars(sterilizedInput string) {
	outputWords := findRandomPositiveWords(sterilizedInput)
	for _, char := range sterilizedInput {
		fmt.Println(strings.ToUpper(string(char)), "stands for", outputWords[char])
	}
}

func findRandomPositiveWords(input string) map[rune]string {
	// Returns a map where the keys are the letters of the input name
	// and the values are words from the positive words list.
	// Note: this map must be 'made' here, not just initialized.
	// Otherwise there will be a runtime error when trying to append to a nil map.
	value := make(map[rune]string)
	positiveWords := mapOfPositiveWords()
	inputRunes := []rune(input)

	// iterate through the input name and find a random word for each letter.
	for _, char := range inputRunes {
		slice := positiveWords[char]
		word := slice[rand.Intn(len(slice))]
		value[char] = word
	}
	return value
}

func mapOfPositiveWords() map[rune][]string {
	// Returns a map where the keys are the letters of the alphabet and the values
	// are slices that contain positive words whose first letter matches the key.
	// Note: again, this map must be 'made' here, not just initialized.
	value := make(map[rune][]string)
	text := readFileOfPositiveWords()

	positiveWordsSlice := strings.Split(text, "\r\n")
	for _, word := range positiveWordsSlice {
		runes := []rune(word)
		slice := value[runes[0]]
		value[runes[0]] = append(slice, word)
	}
	return value
}

func readFileOfPositiveWords() string {
	// Returns a string of text of the contents from the positive words file.
	file, err := os.Open("positive_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}
