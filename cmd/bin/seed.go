package bin

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/hilmiikhsan/shopeefun-product-service/db/seeds"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/adapter"
)

func RunSeed(cmd *flag.FlagSet, args []string) {
	var (
		table = cmd.String("table", "", "seed to run")
		total = cmd.Int("total", 1, "total of records to seed")
	)

	if err := cmd.Parse(args); err != nil {
		log.Fatal().Err(err).Msg("Error while parsing flags")
	}

	adapter.Adapters.Sync(
		adapter.WithShopeefunPostgres(),
	)
	defer func() {
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
	}()

	seeds.Execute(adapter.Adapters.ShopeefunPostgres, *table, *total)
}
