"use client";

import { useAuth } from "@/context/UserContext";
import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import ProductItem from "@/components/ProductItem";
import { Order } from "../orders/page";

export interface Product {
  id: string;
  name: string;
  type: string;
  description: string;
  price: number; // 转换成 number 类型
  currency: string;
  created: number; // 转换成 number 类型
  updated: number; // 转换成 number 类型
}

export interface ProductsResponse {
  data: Product[];
  page: number;
  size: number;
  total: number;
}

// 模拟的 getProducts 函数，你需要替换成真实的 API 调用
const getProducts = async (): Promise<ProductsResponse> => {
  try {
    const response = await fetch("http://localhost:8080/api/products", {
      // 替换成你的 API 端点
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      let errorMessage = `HTTP error! status: ${response.status}`;
      try {
        const errorData = await response.json();
        if (errorData && errorData.message) {
          errorMessage += `: ${errorData.message}`;
        }
      } catch (parseError) {
        console.error("Failed to parse error JSON:", parseError);
      }
      throw new Error(errorMessage);
    }

    const products: ProductsResponse = await response.json();
    return products;
  } catch (error: any) {
    console.error("Error fetching products:", error);
    toast.error(`获取商品失败: ${error.message}`);
    return { data: [], page: 1, size: 10, total: 0 }; // 返回空数组作为错误处理
  }
};

export default function ProductsPage() {
  const { user, isLoading } = useAuth();
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    const fetchProductsData = async () => {
      if (!isLoading && user) {
        const fetchedProducts = await getProducts();
        setProducts(fetchedProducts.data);
      }
    };

    fetchProductsData();
  }, [isLoading, user]);

  if (!user) {
    return null; // 或者返回加载中的状态
  }

  return (
    <>
      {/* Main Content */}
      <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">
          Product Management
        </h2>

        {/* Product management content will go here */}
        <div className="container mx-auto py-8">
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            {products.map((product) => (
              <ProductItem key={product.id} product={product} />
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
