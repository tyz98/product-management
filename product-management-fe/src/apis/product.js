import axios from 'axios';
import { BACKEND_BASE_URL, API_TIMEOUT } from '@/config'

// Configure Axios instance with baseUrl and timeout
const apiInstance = axios.create({
  baseURL: BACKEND_BASE_URL,
  timeout: API_TIMEOUT,
});


const apiUrl = '/api/product';

export const getProducts = function(params) {
 return apiInstance.get(apiUrl, { params });
}
export const getProduct = function(id) {
    return apiInstance.get(`${apiUrl}/${id}`);
}
export const createProduct = function(product) {
    return apiInstance.post(apiUrl, product);
}
export const updateProduct = function(id, product) {
    return apiInstance.put(`${apiUrl}/${id}`, product);
}
export const deleteProduct = function(id) {
    return apiInstance.delete(`${apiUrl}/${id}`);
}
export const getProductTypes = () => {
    return apiInstance.get(`${apiUrl}/type`);
};