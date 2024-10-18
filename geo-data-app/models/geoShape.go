// models/geoShape.go
package models

// GeoShape represents a geospatial shape in the system.
type GeoShape struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Geometry string `json:"geometry"` // Stores the GeoJSON or KML data as a string
}
