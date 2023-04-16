package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"naval/internal/bot"
	"naval/internal/parse"
	"naval/internal/repository"
	"naval/internal/sources"
	"sync"
)

type Config struct {
	ConnectDb string `yaml:"connectdb"`
	DSToken   string `yaml:"dstoken"`
}

func LoadConfigFromYaml() (*Config, error) {
	var conf Config
	f, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "Can't read configs file")
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		fmt.Println("Error read file")
	}
	return &conf, nil
}

func main() {
	config, err := LoadConfigFromYaml()
	if err != nil {
		fmt.Print("Error load configs")
	}
	db, err := sql.Open("postgres", config.ConnectDb)
	if err != nil {
		log.Fatalln("Error connecting to dao", err)
	}
	defer db.Close()

	repo := repository.New(db)

	err = repo.ClearDB()
	if err != nil {
		fmt.Println("Error ClearDB", err)
	}
	parser := parse.New(repo)
	parser.ParceCity()
	parser.Items()
	parser.Port()

	err = repo.CreatTableInfo_city()
	if err != nil {
		log.Fatalln("Error CreatTableInfoCity", err)
	}

	ctx := context.Background()
	mybot := bot.NewBot(sources.NewDs("MTA5NjIxNDcwNTkzNDE3MjMwMQ.Ghkbx4.XIL_jBeuNYsUaMYWOilCJ-v7sdex8lGBEQB7mU"))

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go mybot.RunBot(ctx, wg, repo)

	wg.Wait()
}
