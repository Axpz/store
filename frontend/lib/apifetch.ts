// lib/apiFetch.ts
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

const BASE_URL = process.env.NEXT_PUBLIC_SERVER_URL || 'http://localhost:8080';

const apiClient = axios.create({
  baseURL: BASE_URL,
  withCredentials: true, // 支持 Cookie（适配你设置的 CORS）
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求封装
export async function apiFetch<T = any>(path: string, config: AxiosRequestConfig = {}): Promise<T> {
  try {
    const response: AxiosResponse<T> = await apiClient({
      url: path,
      ...config,
    });
    return response.data;
  } catch (error: any) {
    // 错误处理
    if (error.response) {
      // 服务器响应了一个错误的状态码
      console.error(`API Error: ${error.response.status} ${error.response.statusText}`);
      throw new Error(`API Error: ${error.response.status} ${error.response.statusText}`);
    } else if (error.request) {
      // 请求已发送但没有响应
      console.error('No response received from server');
      throw new Error('No response received from server');
    } else {
      // 其他错误
      console.error(`Request failed: ${error.message}`);
      throw new Error(`Request failed: ${error.message}`);
    }
  }
}

