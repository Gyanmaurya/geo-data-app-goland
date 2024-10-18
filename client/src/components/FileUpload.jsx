// src/components/FileUpload.js
import React from 'react';
import axios from 'axios';

const FileUpload = ({ onFileUpload }) => {
  const handleFileChange = async (event) => {
    const file = event.target.files[0];
    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post('http://localhost:8080/api/upload', formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      });
      onFileUpload(response.data); // Pass the response (GeoJSON) to the parent component
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  };

  return (
    <div className="file-upload">
      <input type="file" onChange={handleFileChange} />
    </div>
  );
};

export default FileUpload;
