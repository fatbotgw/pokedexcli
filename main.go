package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	    cleanedInput := cleanInput(scanner.Text())
	    fmt.Println("Your command was:", cleanedInput[0])
	    fmt.Print(prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
