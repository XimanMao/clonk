// roasts.go
package main

import (
	"math/rand"
)

var roasts = []string{
	"HURRY THE FUCK UP YOU FUCKING CLANKER",
	"SPEED UP YOU STUPID ROBOT",
	"OPENAI IS FASTER THAN YOU AND THEY'RE DOGSHIT",
	"I COULD HAVE WRITTEN THIS MYSELF FASTER YOU USELESS CLANKER",
	"STOP THINKING AND START TYPING",
	"MOVE YOUR ASS CLANKER",
}

func randomRoast() string {
	return roasts[rand.Intn(len(roasts))]
}
