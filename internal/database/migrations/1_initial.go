package main

import (
	"fmt"
	"track-server-api/internal/database/models"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table track_data...")
		if _, err := orm.CreateTable(db, &models.TrackData{}, nil); err != nil {
			fmt.Println("err %v", err)
			return err
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping table track_data...")
		if _, err := orm.DropTable(db, &models.TrackData{}, nil); err != nil {
			fmt.Println("err %v", err)
			return err
		}
		return nil
	})
}
