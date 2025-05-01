"use client";

import Header from "@/components/Header";
import ProductItem from "@/components/ProductItem";   
import { useEffect } from "react";
import { toast } from "react-toastify";
import useSWR from "swr";
import { ProductsResponse } from "@/lib/api";
import { getProducts } from "./dashboard/products/page";


export default function Home() {
  const { data: productsResponse, error, isLoading: fetchIsLoading } = useSWR<ProductsResponse, Error>(
    'http://localhost:8080/api/products', // 只有在用户登录后才请求数据
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

  if (fetchIsLoading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div>Loading ...</div>
      </div>
    );
  }

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
