package main

import (
	"os"
	"strings"
	"testing"
)

func TestIsGameOver(t *testing.T) {
	mistakeCounter := 9
	wordIncomplete := "wor_"

	if !isGameOver(mistakeCounter, wordIncomplete) {
		t.Error("The game should be over, but continues")
	}

	mistakeCounter = 6
	wordIncomplete = "completed"

	if !isGameOver(mistakeCounter, wordIncomplete) {
		t.Error("The game should continue, bud ends")
	}

}

func TestIsHangmanComplete(t *testing.T) {
	if isHangmanComplete(6) {
		t.Error("Hangman should not be completed")
	}

	if !isHangmanComplete(9) {
		t.Error("Hangman should be completed")
	}
}

func TestGetHangmanState(t *testing.T) {
	state, _ := getHangmanState(1)

	expectedState, _ := os.ReadFile("hangmanState/state1.txt")

	if string(state) != string(expectedState) {
		t.Errorf("Expected state1.txt")
	}
}

func TestIsCorrectGuess(t *testing.T) {
	guessWordSplit := []string{"t", "e", "s", "t"}
	guess := "t"
	wordState := []string{"_", "_", "_", "_"}

	if !isCorrectGuess(guessWordSplit, guess, wordState) {
		t.Error("The guess must be correct")
	}

	guess = "a"

	if isCorrectGuess(guessWordSplit, guess, wordState) {
		t.Error("The guess must be incorrect")
	}
}

func TestIsLetterAlreadyUsed(t *testing.T) {
	guess := "t"
	wordState := []string{"t", "e", "s", "t"}

	if !isLetterAlreadyUsed(guess, wordState, 0) {
		t.Error("Letter t must be used, but it's not")
	}

	guess = "a"
	if isLetterAlreadyUsed(guess, wordState, 3) {
		t.Error("Letter a must not be used, but it is")
	}
}

func TestGetCurrentWordState(t *testing.T) {
	expected := "test"
	wordState := []string{"t", "e", "s", "t"}

	res := getCurrentWordState(wordState)

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}

	expected = "bitcoin"
	wordState = []string{"e", "t", "h"}

	res = getCurrentWordState(wordState)

	if res == expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestIsWordComplete(t *testing.T) {
	if !isWordComplete("test") {
		t.Error("The word must be complete")
	}

	if isWordComplete("t____t") {
		t.Error("The word must be incomplete")
	}
}

func TestGetInitialState(t *testing.T) {
	expected := []string{"_", "_", "_", "_", "_", "_"}
	res := getInitialState([]string{"z", "o", "m", "b", "i", "e"})
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %s, got %s", expected, res)
		}
	}
}

func TestGenerateGuessWord(t *testing.T) {
	guessWords := []string{"apple", "United States of America", "bitcoin", "disney", "book"}

	res := strings.Join(generateGuessWord(), "")

	wordExists := false
	for _, v := range guessWords {
		if res == v {
			wordExists = true
		}
	}

	if !wordExists {
		t.Errorf("%s does not exist in the slice", res)
	}
}
