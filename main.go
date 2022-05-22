package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println((event.Timestamp))
		fmt.Println(event.Parameters)
		fmt.Println((event.Event))
		fmt.Println()

	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3560972372132-3571246630993-bW8wM77sSqHWfHhBH9STSZ0Z")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03GBLT0W9K-3544026715863-2e5cad99e589ce74d38262060834b780b831359762d387b69c54431e5168fc16")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>  ", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "My yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, respsonse slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			respsonse.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
