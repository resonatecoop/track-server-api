package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type TrackData struct {
	Id        uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time `sql:"default:now()"`
	UserId    uuid.UUID `sql:"type:uuid,notnull"`
	TrackId   uuid.UUID `sql:"type:uuid,notnull"`
	Type      int
	Credits   float32
}
