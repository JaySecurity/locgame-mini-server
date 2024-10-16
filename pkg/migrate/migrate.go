package migrate

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Migrate is type for performing migrations in provided database.
// Database versioned using dedicated collection.
// Each migration applying ("up" and "down") adds new document to collection.
// This document consists migration version, migration description and timestamp.
// Current database version determined as version in latest added document (biggest "_id") from collection mentioned above.
type Migrate struct {
	db                   *mongo.Database
	migrations           []Migration
	migrationsCollection string
}

type versionRecord struct {
	Version     uint64    `bson:"version"`
	Description string    `bson:"description,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`
}

func NewMigrate(db *mongo.Database, migrations ...Migration) *Migrate {
	internalMigrations := make([]Migration, len(migrations))
	copy(internalMigrations, migrations)
	return &Migrate{
		db:                   db,
		migrations:           internalMigrations,
		migrationsCollection: "migrations",
	}
}

// Version returns current database version and comment.
func (m *Migrate) Version() (uint64, string, error) {
	filter := bson.D{{}}
	sort := bson.D{bson.E{Key: "_id", Value: -1}}
	opts := options.FindOne().SetSort(sort)

	// find record with the greatest id (assuming it`s latest also)
	result := m.db.Collection(m.migrationsCollection).FindOne(context.Background(), filter, opts)
	err := result.Err()
	switch {
	case err == mongo.ErrNoDocuments:
		return 0, "", nil
	case err != nil:
		return 0, "", err
	}

	var rec versionRecord
	if err := result.Decode(&rec); err != nil {
		return 0, "", err
	}

	return rec.Version, rec.Description, nil
}

func (m *Migrate) SetVersion(version uint64, description string) error {
	rec := versionRecord{
		Version:     version,
		Timestamp:   time.Now().UTC(),
		Description: description,
	}

	_, err := m.db.Collection(m.migrationsCollection).InsertOne(context.Background(), rec)
	if err != nil {
		return err
	}

	return nil
}

// Up performs "up" migrations to latest available version.
// If n<=0 all "up" migrations with newer versions will be performed.
// If n>0 only n migrations with newer version will be performed.
func (m *Migrate) Up(n int) error {
	currentVersion, _, err := m.Version()
	if err != nil {
		return err
	}
	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := 0, 0; i < len(m.migrations) && p < n; i++ {
		migration := m.migrations[i]
		if migration.Version <= currentVersion || migration.Up == nil {
			continue
		}
		p++
		if err := migration.Up(m.db); err != nil {
			return err
		}
		if err := m.SetVersion(migration.Version, migration.Description); err != nil {
			return err
		}
	}
	return nil
}

// Down performs "down" migration to oldest available version.
// If n<=0 all "down" migrations with older version will be performed.
// If n>0 only n migrations with older version will be performed.
func (m *Migrate) Down(n int) error {
	currentVersion, _, err := m.Version()
	if err != nil {
		return err
	}
	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := len(m.migrations)-1, 0; i >= 0 && p < n; i-- {
		migration := m.migrations[i]
		if migration.Version > currentVersion || migration.Down == nil {
			continue
		}
		p++
		if err := migration.Down(m.db); err != nil {
			return err
		}

		var prevMigration Migration
		if i == 0 {
			prevMigration = Migration{Version: 0}
		} else {
			prevMigration = m.migrations[i-1]
		}
		if err := m.SetVersion(prevMigration.Version, prevMigration.Description); err != nil {
			return err
		}
	}
	return nil
}
