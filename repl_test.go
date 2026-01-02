package main

import (
	"testing"
    "fmt"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        // add more cases here
        {
            // Charmander Bulbasaur PIKACHU -> ["charmander", "bulbasaur", "pikachu"]
            input:  "Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        // Check the length of the actual slice against the expected slice
        // if they don't match, use t.Errorf to print an error message
        // and fail the test
        fmt.Printf("Actual:   %v \n", actual)
        fmt.Printf("Expected: %v \n", c.expected)
        if len(actual) != len(c.expected) {
            t.Errorf("slices don't match")
            t.FailNow()
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            // Check each word in the slice
            // if they don't match, use t.Errorf to print an error message
            // and fail the test
            if word != expectedWord {
                t.Errorf("Word does not match Expected Word")
                t.FailNow()
            }
        }
    }
}