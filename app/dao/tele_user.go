package dao

import (
	"github.com/kmou424/runabot/app/model"
	"github.com/pocketbase/dbx"
)

func TeleUserQuery() *dbx.SelectQuery {
	return dao.ModelQuery(&model.TeleUser{})
}
