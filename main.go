package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("What is your name?")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Wow,", input, "is a cool name! Let's break it down by letter.")
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
	value := make(map[rune]string)
	positiveWords := mapOfPositiveWords()
	inputRunes := []rune(input)

	for _, char := range inputRunes {
		slice := positiveWords[char]
		word := slice[rand.Intn(len(slice))]
		value[char] = word
	}
	return value
}

func mapOfPositiveWords() map[rune][]string {
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
