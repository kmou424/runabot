package dao

import (
	G "github.com/kmou424/runabot/app/global"
	"github.com/kmou424/runabot/app/model"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var dao *daos.Dao

func Init() {
	dao = G.Server.Dao()

	initCollections(
		model.TeleUserCollection,
	)
}

func initCollections(collections ...*models.Collection) {
	initCollection := func(collection *models.Collection) {
		logger := G.Logger.With("name", collection.Name)

		if oldCollection, err := dao.FindCollectionByNameOrId(collection.Name); err != nil {
			logger.Info("creating collection")
		} else {
			logger.Info("updating collection")
			collection.BaseModel = oldCollection.BaseModel
		}

		err := dao.SaveCollection(collection)
		if err != nil {
			logger.Error("create/update collection failed", "err", err)
			G.Exit()
		}
	}
	for _, collection := range collections {
		initCollection(collection)
	}
}
