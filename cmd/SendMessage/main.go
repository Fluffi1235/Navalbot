package main

import (
	"context"
	"naval/internal/bot"
	"naval/internal/parse"
	"naval/internal/service"
	"naval/internal/sources"
	"sync"
)

func main() {
	service.ClearDb()
	parse.ParceCity()
	parse.Items()
	parse.Port()
	ctx := context.Background()
	mybot := bot.NewBot(sources.NewTG())

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go mybot.RunBot(ctx, wg)

	wg.Wait()
}
