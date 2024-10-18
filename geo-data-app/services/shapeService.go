
package services

import (
	"geo-data-app/models"
	"geo-data-app/utils"
)

func SaveShape(shape models.GeoShape) error {
	_, err := utils.DB.Exec("INSERT INTO shapes (name, geometry) VALUES ($1, $2)", shape.Name, shape.Geometry)
	return err
}

func GetAllShapes() ([]models.GeoShape, error) {
	rows, err := utils.DB.Query("SELECT id, name, geometry FROM shapes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shapes []models.GeoShape
	for rows.Next() {
		var shape models.GeoShape
		err = rows.Scan(&shape.ID, &shape.Name, &shape.Geometry)
		if err != nil {
			return nil, err
		}
		shapes = append(shapes, shape)
	}
	return shapes, nil
}
