package main

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/coreyvan/go-address/internal/app"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := run(log.Sugar()); err != nil {
		log.Fatal(err.Error())
	}
}

func run(log *zap.SugaredLogger) error {
	cfg := app.Config{}

	help, err := conf.Parse("ADDRESS", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	log.Infow("using config", "cfg", cfg)

	return app.Run(cfg, log)
}
