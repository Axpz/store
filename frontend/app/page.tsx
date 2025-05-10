"use client";

import Header from "@/components/Header";
import ProductItem from "@/components/ProductItem";
import { useEffect } from "react";
import { toast } from "react-toastify";
import useSWR from "swr";
import { ProductsResponse } from "@/lib/api";
import { fetchAndStoreProducts } from "@/lib/api";
import { useProductStore } from "./store/productStore";

export default function Home() {
  const { productsArray, setProducts } = useProductStore();

  // 判断 productsArray 是否为空，且发送请求
  const {
    data: productsResponse,
    error: fetchError,
    isLoading: fetchIsLoading,
  } = useSWR<ProductsResponse, Error>(
    productsArray.length === 0 ? "http://localhost:8080/api/products" : null,
    fetchAndStoreProducts,
    { revalidateOnFocus: false }
  );

  // 处理请求错误
  useEffect(() => {
    if (fetchError) {
      console.error("Error fetching products:", fetchError);
      toast.error(`Error fetching products: ${fetchError.message}`);
    }
  }, [fetchError]);

  // 如果有新的产品数据，更新到状态
  useEffect(() => {
    if (productsResponse?.data) {
      setProducts(productsResponse.data);
    }
  }, [productsResponse, setProducts]);

  // 如果正在加载数据，显示 loading 页面
  if (fetchIsLoading || productsArray.length === 0) {
    return (
      <>
        <Header />
        <div className="flex justify-center items-center h-screen">
          <div>Loading...</div>
        </div>
      </>
    );
  }

  // 产品数据加载完成后渲染
  const products = productsArray;

  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-center text-center max-w-7xl mx-auto">
        <main className="flex flex-col items-center gap-12 w-full">
          <div className="text-center max-w-4xl">
            {/* <h1 className="text-[32px] font-bold leading-tight tracking-tight">
              Radix UI templates and examples
            </h1> */}
            {/* <p className="text-xl text-gray-600 dark:text-gray-400 mt-6">
              Browse examples and templates of application builds using Radix
              UI.
            </p> */}
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full">
            {products.map((product, index) => (
              <ProductItem key={index} product={product} />
            ))}
          </div>
        </main>
      </div>
    </>
  );
}
