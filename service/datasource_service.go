package service

import (
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
)

func FindDatasources(filter *model.Datasource) []model.Datasource {
	var datasources []model.Datasource
	dao.DB.Where(&filter).Find(&datasources)

	return datasources
}
