package main

import (
	"io/ioutil"
	"log"

	//"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	discordToken, err := ioutil.ReadFile("token.txt")
	check(err)

	bot, err := discordgo.New("Bot " + string(discordToken))
	check(err)

	err = bot.Open()
	check(err)

	log.Print("Logged into Discord successfully")

	// Wait till we die, like in reality
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		bot.Close()
		log.Print("Closed connection to Discord. Goodbye.")
	}()

	// Create a dgc router
	router := dgc.Create(&dgc.Router{
		Prefixes:    []string{"."},
		BotsAllowed: false,
	},
	)

	router.RegisterDefaultHelpCommand(bot, nil) // Add help command

	router.RegisterCmd(&dgc.Command{
		Name:        "minesweeper",
		Description: "Play a round of minesweeper.",
		Usage:       ".minesweeper <Size [Max 12]> <Amount of Bombs [Max: Size*Size-1]>",
		Example:     ".minesweeper 23 5",
		IgnoreCase:  true,
		Handler:     minesweeper,
	})

	router.Initialize(bot)
}
