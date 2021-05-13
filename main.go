/*
* PROJECT: go-project-13-05-2021
* DESCRIPTION: Given a string, find the word with the highest score by adding the letters of the word based on its position in the alphabet.
* AUTHOR: Vahin Sharma
*/

package main

import (
  "fmt"
  "strings"
)

func wordScore(s string) (score byte) {
  for i, _ := range s {
    score += s[i] - 96
  }
  return
}

func solution(s string) string {
  var score, newScore byte
  var word string
  for _, splitWord := range strings.Split(s, " ") {
    newScore = wordScore(splitWord)
    if newScore > score {
      score = newScore
      word = splitWord
    }
  }
  return word
}

func main() {
  tests := map[string]string {
    "man i need a taxi up to ubud": "taxi",
    "what time are we climbing up the volcano": "volcano",
    "take me to semynak": "semynak",
    "take two bintang and a dance please": "bintang",
  }
  score := 0
  
  for q, a := range tests {
    fmt.Println("Input:", q)
    if returnedA := solution(q); returnedA != a {
      fmt.Printf("Expected %d, instead got %d\n\n", a, returnedA)
    } else {
      score++
      fmt.Println("Passed\n")
    }
  }
  
  fmt.Printf("Score: %d out of %d\n", score, len(tests))
}
