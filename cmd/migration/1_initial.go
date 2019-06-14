package main

import (
	"fmt"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"

	"track-server-api/internal/model"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table track_data...")
		options := &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		}
		if err := orm.CreateTable(db.(orm.DB), &model.TrackData{}, options); err != nil {
			fmt.Println("err %v", err)
			return err
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping table track_data...")
		options := &orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		}
		if err := orm.DropTable(db.(orm.DB), &model.TrackData{}, options); err != nil {
			fmt.Println("err %v", err)
			return err
		}
		return nil
	})
}
