package db

import "sokol_proxy/models"

func InsertData(tableName string, data []models.SokolM) error {
	for _, s := range data {
		err := db.Table(tableName).Where("time = ?", s.Time).FirstOrCreate(&s).Error

		if err != nil {
			return err
		}
	}
	return nil
}
