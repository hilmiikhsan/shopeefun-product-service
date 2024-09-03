package seeds

import (
	"context"
	"os"

	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// Seed struct.
type Seed struct {
	db *sqlx.DB
}

// NewSeed returns a Seed with a pool of connection to a database.
func newSeed(db *sqlx.DB) Seed {
	return Seed{
		db: db,
	}
}

func Execute(db *sqlx.DB, table string, total int) {
	seed := newSeed(db)
	seed.run(table, total)
}

// Run seeds.
func (s *Seed) run(table string, total int) {
	switch table {
	case "products":
		s.productsSeed(total)
	case "delete-all":
		s.deleteAll()
	default:
		log.Warn().Msg("No seed to run")
	}

	if table != "" {
		log.Info().Msg("Seed ran successfully")
		log.Info().Msg("Exiting ...")
		if err := s.db.Close(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
		os.Exit(0)
	}
}

func (s *Seed) deleteAll() {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		} else {
			err = tx.Commit()
			if err != nil {
				log.Error().Err(err).Msg("Error committing transaction")
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM products`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting products")
		return
	}
	log.Info().Msg("products table deleted successfully")

	log.Info().Msg("=== All tables deleted successfully ===")
}

func (s *Seed) productsSeed(total int) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	type Shop struct {
		ID string `db:"id"`
	}

	var shops []Shop
	err = s.db.Select(&shops, `SELECT id FROM shops`)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting shops")
		return
	}

	if len(shops) == 0 {
		log.Warn().Msg("No shops found. Products seeding aborted.")
		return
	}

	productMaps := make([]map[string]any, 0, total)

	for i := 0; i < total; i++ {
		selectedShop := shops[rand.Intn(len(shops))]

		dataProductToInsert := map[string]any{
			"id":          uuid.New().String(),
			"shop_id":     selectedShop.ID,
			"name":        gofakeit.ProductName(),
			"description": gofakeit.Paragraph(1, 3, 10, " "),
			"category":    gofakeit.Word(),
			"price":       gofakeit.Price(1, 1000),
			"stock":       gofakeit.Number(0, 100),
			"rating":      gofakeit.Number(1, 5),
			"brand":       gofakeit.Company(),
			"created_at":  gofakeit.Date(),
			"updated_at":  gofakeit.Date(),
		}

		productMaps = append(productMaps, dataProductToInsert)
	}

	_, err = tx.NamedExec(`
		INSERT INTO products (id, shop_id, name, description, category, price, stock, rating, brand, created_at, updated_at)
		VALUES (:id, :shop_id, :name, :description, :category, :price, :stock, :rating, :brand, :created_at, :updated_at)
	`, productMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating products")
		return
	}

	log.Info().Msg("products table seeded successfully")
}
