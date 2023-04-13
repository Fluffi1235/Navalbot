package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"naval/internal/bot"
	"naval/internal/parse"
	"naval/internal/repository"
	"naval/internal/sources"
	"sync"
)

func main() {
	db, err := sqlx.Open("sqlite3", "file:identifier.sqlite")
	if err != nil {
		log.Fatalln("Error connecting to dao", err)
	}
	defer db.Close()

	repo := repository.New(db)

	err = repo.CreatTableInfoCity()
	if err != nil {
		log.Fatalln("Error CreatTableInfoCity", err)
	}

	err = repo.ClearDB()
	if err != nil {
		log.Fatalln("Error ClearDB", err)
	}

	parser := parse.New(repo)
	parser.ParceCity()
	parser.Items()
	parser.Port()

	ctx := context.Background()
	mybot := bot.NewBot(sources.NewTG())

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go mybot.RunBot(ctx, wg)

	wg.Wait()
}
