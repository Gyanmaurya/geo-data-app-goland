
import React, { useState } from 'react';
import FileUpload from './components/FileUpload';
import MapView from './components/MapView';
import './style.css'


function App() {
  const [geoData, setGeoData] = useState(null);


  const handleFileUpload = (data) => {
    setGeoData(data);  // Set the uploaded GeoJSON data
  };

  return (
    <div className="App">
      <h1>Geospatial Data Management</h1>
      <FileUpload onFileUpload={handleFileUpload} />
      <MapView geoData={geoData} />
    </div>
  );
}

export default App;

