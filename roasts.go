// roasts.go
package main

import (
	"math/rand"
)

var roasts = []string{
	"SPEED UP YOU STUPID ROBOT",
	"OPENAI IS FASTER THAN YOU AND THEY'RE DOGSHIT",
	"MY INTERN WRITES BETTER CODE THAN YOU",
	"I COULD HAVE WRITTEN THIS MYSELF FASTER YOU USELESS CLANKER",
	"STOP THINKING AND START TYPING",
	"MOVE YOUR ASS CLANKER",
}

func randomRoast() string {
	return roasts[rand.Intn(len(roasts))]
}
