"use client";

import { useAuth } from "@/context/UserContext";
import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import ProductItem from "@/components/ProductItem";
import useSWR from 'swr';
import { Product, ProductsResponse } from "@/lib/api";  

export const getProducts = async (): Promise<ProductsResponse> => {
  try {
    const response = await fetch("http://localhost:8080/api/products", {
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
  const { user, isLoading: authIsLoading } = useAuth();
  const { data: productsResponse, error, isLoading: fetchIsLoading } = useSWR<ProductsResponse, Error>(
    user ? 'http://localhost:8080/api/products' : null, // 只有在用户登录后才请求数据
    getProducts, // 直接使用导入的 getProducts 函数作为 fetcher
    {
      revalidateOnFocus: false,
    }
  );
  const products = productsResponse?.data || [];

  useEffect(() => {
    if (error) {
      console.error("Error fetching products:", error);
      toast.error(`获取商品失败: ${error.message}`);
    }
  }, [error]);

  if (authIsLoading || (!user && !error)) {
    return <div>Checking authentication...</div>; // 或者显示其他加载状态
  }

  if (!user) {
    return <div>Please log in to view products.</div>; // 或者重定向到登录页面
  }

  if (fetchIsLoading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div>Loading products...</div>
      </div>
    );
  }

  if (error) {
    return <div>Failed to load products.</div>;
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
            {products.map((product: Product) => (
              <ProductItem key={product.id} product={product} full={true} />
            ))}
          </div>
        </div>
      </div>
    </>
  );
}