"use client";

import { useAuth } from "@/context/UserContext";
import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import ProductItem from "@/components/ProductItem";
import useSWR from 'swr';
import { Product, ProductsResponse } from "@/lib/api";  
import { useProductStore } from "@/app/store/productStore";

export default function ProductsPage() {

  const { selectedProduct } = useProductStore();
  if (!selectedProduct) {
    toast.error("No product selected");
    return; 
  }

  const product = selectedProduct;

  return (
    <>
      {/* Main Content */}
      <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8 min-h-screen">
  
        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">
          Product {product.id}
        </h2>
  
        {/* Product management content will go here */}
        <div className="w-full py-8">
          {/* Using full width for the product item */}
          <div className="w-full">
            <ProductItem key={product.id} product={product} full={true} checkout={true} />
          </div>
        </div>
  
      </div>
    </>
  );
}