package model

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*TeleUser)(nil)

type TeleUser struct {
	models.BaseModel

	UserId string `db:"uid" json:"uid"`
}

func (_ *TeleUser) TableName() string {
	return "tele_users"
}

var TeleUserCollection = &models.Collection{
	Name: "tele_users",
	Type: models.CollectionTypeBase,
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:     "uid",
			Type:     schema.FieldTypeText,
			Required: true,
			Options:  nil,
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE UNIQUE INDEX idx_uid ON tele_users(uid)",
	},
}
