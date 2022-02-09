# hangman

A simple CLI Hangman game. The game uses some predefined words that the user have to guess. To play the game you have to run:

- `go build`
- `./hangman`

Or

- `go run main.go`

After the start you will be prompted to enter a letter. If it's correct it will reveal it's position in the word, if not it will start drawing the hangman.
You have 9 chances to guess the word. Only single alphanumeric characters are accepted.
