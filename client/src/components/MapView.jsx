import React, { useState } from 'react';
import { MapContainer, TileLayer, FeatureGroup, GeoJSON } from 'react-leaflet';
import { EditControl } from 'react-leaflet-draw';
import './index.css'
import axios from 'axios';

const MapView = ({ geoData }) => {
  const [shapes, setShapes] = useState([]);

  const handleShapeCreated = async (e) => {
    const layer = e.layer;
    const shapeData = layer.toGeoJSON();
    setShapes([...shapes, shapeData]);
    
    try {
      const response = await axios.post('http://localhost:8080/api/save-shape', shapeData);
      console.log('Shape saved:', response.data);
    } catch (error) {
      console.error('Error saving shape:', error);
    }
  };

  return (
    <MapContainer center={[51.505, -0.09]} zoom={13} style={{ height: '600px', width: '100%' }}>
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      <FeatureGroup>
        <EditControl
          position="topright"
          onCreated={handleShapeCreated}
          draw={{
            rectangle: true,
            polygon: true,
            circle: false,
            marker: false,
          }}
        />
      </FeatureGroup>
      {geoData && <GeoJSON data={geoData} />}
    </MapContainer>
  );
};

export default MapView;
