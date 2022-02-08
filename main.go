package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Printing game state
//	* Print word you're-guessing
//	* Print hangman state
// Derive a word we have to guess
// Read user input
//	* validate it (e.g. only letters)
// Determine if the letter is a correct guess or not
//	* if correct, update the guessed letters
//	* if incorrect, update the hangman state
// If word is guessed -> game over, you win
// If hangman is complete -› game over, you lose

var guessWords = []string{"apple", "United States of America", "bitcoin", "disney", "book"}
var isLetter = regexp.MustCompile(`^[a-zA-Z]$`).MatchString
var reader = bufio.NewReader(os.Stdin)

func main() {
	guessWordSplit := generateGuessWord()

	wordState := getInitialState(guessWordSplit)
	printInitialState(wordState)

	mistakeCounter := 0

	for !isGameOver(mistakeCounter, getCurrentWordState(wordState)) {
		guess := readInput()

		if !isLetter(guess) {
			fmt.Println("Invalid input!")
			continue
		}

		correctGuess := isCorrectGuess(guessWordSplit, guess, wordState)

		printWordState(wordState)

		if !correctGuess {
			mistakeCounter++
			printHangmanState(mistakeCounter)
		}
	}
}

func printHangmanState(mistakeCounter int) {
	hangmanState, err := getHangmanState(mistakeCounter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hangmanState))
}

func printWordState(wordState []string) {
	fmt.Println(getCurrentWordState(wordState))
}

func isGameOver(mistakeCounter int, wordState string) bool {
	if isHangmanComplete(mistakeCounter) && !isWordComplete(wordState) {
		fmt.Println("GAME OVER, YOU LOST!")
		return true
	} else if isWordComplete(wordState) && !isHangmanComplete(mistakeCounter) {
		fmt.Println("GAME OVER, YOU WON!")
		return true
	}

	return false
}

func readInput() string {
	fmt.Println("Enter your guess.")
	fmt.Print("> ")
	guess, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	guess = strings.TrimSpace(guess)
	return guess
}

func isHangmanComplete(mistakeCounter int) bool {
	return mistakeCounter == 9
}

func getHangmanState(mistakeCounter int) ([]byte, error) {
	fileName := fmt.Sprintf("hangmanState/state%d.txt", mistakeCounter)
	hangmanState, err := os.ReadFile(fileName)

	return hangmanState, err
}

func isCorrectGuess(guessWordSplit []string, guess string, wordState []string) bool {
	correctGuess := false

	for i, char := range guessWordSplit {
		guess = strings.ToLower(guess)
		char = strings.ToLower(char)

		if isLetterAlreadyUsed(guess, wordState, i) {
			return true
		}

		if guess == char {
			wordState[i] = char
			correctGuess = true
		}
	}

	return correctGuess
}

func isLetterAlreadyUsed(guess string, wordState []string, i int) bool {
	if guess == wordState[i] {
		fmt.Printf("You already used the letter %s\n", guess)
		return true
	}
	return false
}

func getCurrentWordState(wordState []string) string {
	return strings.Join(wordState, "")
}

func isWordComplete(wordState string) bool {
	return !strings.Contains(wordState, "_")
}

func printInitialState(wordState []string) {
	str := strings.Join(wordState, "")
	fmt.Println(str)
}

func getInitialState(guessWord []string) []string {
	var res []string

	for _, char := range guessWord {
		if char == " " {
			res = append(res, " ")
		} else {
			res = append(res, "_")
		}
	}

	return res
}

func generateGuessWord() []string {
	rand.Seed(time.Now().UnixNano())
	word := guessWords[rand.Intn(len(guessWords))]

	return strings.Split(word, "")
}
