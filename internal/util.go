package internal

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"
)

func CheckError(err error, table string) twirp.Error {
	if err != nil {
		if err == pg.ErrNoRows {
			return twirp.NotFoundError(fmt.Sprintf("%s does not exist", table))
		}
		twerr, ok := err.(twirp.Error)
		if ok && twerr.Meta("argument") == "id" {
			argument := "id"
			if table != "" {
				argument = table + " id"
			}
			return twirp.InvalidArgumentError(argument, "must be a valid uuid")
		}
		pgerr, ok := err.(pg.Error)
		if ok {
			code := pgerr.Field('C')
			name := pgerr.Field('n')
			var message string
			if code == "23505" { // unique_violation
				message = strings.TrimPrefix(strings.TrimSuffix(name, "_key"), fmt.Sprintf("%ss_", table))
				return twirp.NewError("already_exists", message)
			} else if code == "23503" { // foreign_key_violation
				message = pgerr.Field('M')
				return twirp.NotFoundError(message)
			} else {
				message = pgerr.Field('M')
				fmt.Println(twirp.NewError("unknown", message))
				return twirp.NewError("unknown", message)
			}
		}
		return twirp.NewError("unknown", err.Error())
	}
	return nil
}

func GetUuidFromString(id string) (uuid.UUID, twirp.Error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return uuid.UUID{}, twirp.InvalidArgumentError("id", "must be a valid uuid")
	}
	return uid, nil
}

// Compare two uuid slices
func Equal(a, b []uuid.UUID) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
