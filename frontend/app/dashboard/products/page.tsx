"use client";

import { useAuth } from "@/context/UserContext";
import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import ProductItem from "@/components/ProductItem";
import { fetchAndStoreProducts, Product, ProductsResponse } from "@/lib/api";  
import { useProductStore } from "@/app/store/productStore";

export default function ProductsPage() {
  const { user, isLoading: authIsLoading } = useAuth();
  const { productsArray, setProducts } = useProductStore();
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    if (user && productsArray.length === 0) {
      fetchAndStoreProducts("http://localhost:8080/api/products").then((products) => {
        setProducts(products.data);
      });
    }
  }, [user]);

  if (authIsLoading || (!user && !error)) {
    return toast.info("Checking authentication...");
  }

  if (!user) {
    return toast.info("Please log in to view products.");
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
          <div className="grid gap-4 grid-cols-auto-fit">
            {productsArray.map((product: Product) => (
              <ProductItem key={product.id} product={product} />
            ))}
          </div>
        </div>
      </div>
    </>
  );
}