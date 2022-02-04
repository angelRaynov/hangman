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

func main() {
	guessWord := generateGuessWord()
	guessWordSplit := strings.Split(guessWord, "")

	initialWordState := getInitialState(guessWordSplit)
	printInitialState(initialWordState)

	reader := bufio.NewReader(os.Stdin)
	mistakeCounter := 0
	for {
		fmt.Println("Enter your guess.")

		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)

		if !isLetter(guess) {
			fmt.Println("Invalid input!")
		}

		correctGuess := false
		for i, char := range guessWordSplit {
			if strings.ToLower(guess) == strings.ToLower(char) {
				initialWordState[i] = char
				correctGuess = true
			}
		}

		printCurrentWordState(initialWordState)

		if !correctGuess {
			mistakeCounter++
			fileName := fmt.Sprintf("hangmanState/state%d.txt", mistakeCounter)
			hangmanState, err := os.ReadFile(fileName)
			if err != nil {
				log.Fatal(err)
			}

			if mistakeCounter == 9 {
				fmt.Println("GAME OVER!")
				break
			}

			fmt.Println(string(hangmanState))

		}
	}
}

func printCurrentWordState(initialWordState []string) {
	wordState := strings.Join(initialWordState, "")

	fmt.Println(wordState)

	if !strings.Contains(wordState, "_") {
		fmt.Println("You won!")
		os.Exit(0)
	}

}

func printInitialState(initialWordState []string) {
	str := strings.Join(initialWordState, "")
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

func generateGuessWord() string {
	rand.Seed(time.Now().UnixNano())
	return guessWords[rand.Intn(len(guessWords))]

}
