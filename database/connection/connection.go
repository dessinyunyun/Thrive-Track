package connection

import (
	"context"
	"fmt"
	"go-gin/database/ent"
	"go-gin/database/ent/migrate"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func ConnectionDB(ctx context.Context, log *logrus.Entry, dbuser string, dbpass string, dbhost string, dbport string, dbname string) *ent.Client {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname,
	)

	// Inisialisasi schema Ent
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Errorf("Failed to open connection Ent: %v", err)
		return nil
	}

	// Migrasi schema ke database
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				// Run custom code here.
				return next.Create(ctx, tables...)
			})
		}),
	)
	if err != nil {
		log.Errorf("Gagal melakukan migrasi schema: %v", err)
		return nil
	}

	return client
}
