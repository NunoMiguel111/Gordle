package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Game holds all the information we need to play a game
type Game struct {
	reader *bufio.Reader
}

// New returns a Game, which can be used to Play!
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// ask for a valid word
	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

const wordLength = 5

// ask reads input untila a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess: \n", wordLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())

		}

		guess := []rune(string(playerInput))

		// TODO Verify the suggestion has a valid length
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attemp is invalid with Gordle's solution: %s. \n", err.Error())

		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != wordLength /*solutionLength*/ {
		return fmt.Errorf("expected %d, got %d, %w", wordLength /*solutionLength*/, len(guess), errInvalidWordLength)
	}
	return nil
}
