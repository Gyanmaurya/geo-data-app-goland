// src/services/api.js
import axios from 'axios';

export const uploadGeoJSON = async (file) => {
  const formData = new FormData();
  formData.append('file', file);
  return axios.post('http://localhost:8080/api/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  });
};

export const saveShape = async (shape) => {
  return axios.post('http://localhost:8080/api/save-shape', shape);
};
