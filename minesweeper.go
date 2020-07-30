package main

import (
	"math/rand"
	"time"

	"github.com/Lukaesebrot/dgc"
)

func minesweeper(ctx *dgc.Ctx) {
	var response string
	rand.Seed(time.Now().UnixNano()) // Initalize Randomness

	argumentSize := ctx.Arguments.Get(0)
	fieldsize, err := argumentSize.AsInt()
	if err != nil {
		ctx.RespondText("Incorrect Fieldsize Argument")
		return
	}

	argumentBombs := ctx.Arguments.Get(1)
	bombs, err := argumentBombs.AsInt()
	if err != nil {
		ctx.RespondText("Incorrect Bomb Amount Argument")
		return
	}

	if fieldsize > 12 {
		ctx.RespondText("The field is too big! [Max: 12]")
		return
	} else if fieldsize < 2 {
		ctx.RespondText("The field is too small [Min: 2]")
		return
	} else if bombs*bombs >= fieldsize {
		ctx.RespondText("Too many bombs. Max amount for your field: " + string(fieldsize^2-1))
		return
	}

	fieldmap := make([][]byte, fieldsize) // Initalize Minesweeper Playingfield
	for i := range fieldmap {
		fieldmap[i] = make([]byte, fieldsize)
	}

	for i := 0; i < fieldsize; i++ {
		// Bomb planting
		x := rand.Intn(fieldsize)
		y := rand.Intn(fieldsize)

		if fieldmap[x][y] != 9 { // 9 = Bombvalue
			fieldmap[x][y] = 9
		}
	}

	for x := 0; x < fieldsize; x++ {
		for y := 0; y < fieldsize; y++ {
			var bombcount byte

			if fieldmap[x][y] == 9 { // Detect if Bomb
				response += "|| :boom: "
			}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if x-1+i < 0 || y-1+j < 0 {
						continue
					} else if x-1+i > fieldsize || y-1+j > fieldsize {
						continue
					}

					if fieldmap[x-1+i][y-1+j] == 9 { // Detect Bomb
						bombcount++
					}
				}
			}

			response += "|| "
			switch bombcount {
			case 0:
				response += ":zero: "
			case 1:
				response += ":one: "
			case 2:
				response += ":two: "
			case 3:
				response += ":three: "
			case 4:
				response += ":four: "
			case 5:
				response += ":five: "
			case 6:
				response += ":six: "
			case 7:
				response += ":seven: "
			case 8:
				response += ":eight: "
			}
		}
		response += "||\n"
	}

	ctx.RespondText(response)
}
