package data

import (
	"context"

	"github.com/alichz2001/gombrella/internal/conf"
	"github.com/alichz2001/gombrella/services/todo/data/ent"

	"github.com/go-kratos/kratos/v3/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTodoRepo)

// Data holds the long-lived storage clients shared by repos.
type Data struct {
	db *ent.Client
}

// NewData opens the database client and returns it with a cleanup function.
func NewData(c *conf.Data) (*Data, func(), error) {
	dc := c.GetDatabase()
	// The generated client inherits the service namespace from its Ent schemas,
	// so repositories never need schema prefixes or session-level search_path.
	db, err := ent.Open(dc.GetDriver(), dc.GetSource())
	if err != nil {
		return nil, nil, err
	}
	if dc.GetDebug() {
		db = db.Debug()
	}
	// Auto migration is a convenience for local development. PostgreSQL enum
	// types must already exist; apply the versioned SQL migrations first.
	// In production, apply schema changes as a separate reviewed step instead.
	if dc.GetAutoMigrate() {
		if err := db.Schema.Create(context.Background()); err != nil {
			db.Close()
			return nil, nil, err
		}
	}
	cleanup := func() {
		log.Info("closing the data resources")
		if err := db.Close(); err != nil {
			log.Error("failed closing the database", "err", err)
		}
	}
	return &Data{db: db}, cleanup, nil
}
