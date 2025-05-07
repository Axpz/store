"use client";

import { PayPalScriptProvider } from "@paypal/react-paypal-js";
import { useRouter } from "next/navigation";
import { useEffect, useReducer, useState } from "react";
import { toast } from "react-toastify";
import PayPalButton from "@/components/paypal-button";
import { Product } from "@/lib/api";
import { useProductStore } from "@/app/store/productStore";
import { useAuth } from "@/context/UserContext";

export default function CheckoutPage() {
  const { user, isLoading } = useAuth();

  if (isLoading) {
    toast.error("Loading...");
    return <div>Loading...</div>;
  }

  if (!user) {
    // toast.error("Please login");
    return <div>Please login</div>;
  }

  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const { selectedProduct } = useProductStore();
  if (!selectedProduct) {
    toast.error("No product selected");
    return;
  }

  const product = selectedProduct;

  const handlePaymentSuccess = () => {
    toast.success("Payment successful!");
    router.push("/dashboard/orders");
  };

  const handlePaymentError = (error: any) => {
    console.error("Payment error:", error);
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!product) {
    return <div>Product not found</div>;
  }

  return (
    <div className="container mx-auto py-8">
      <h1 className="text-2xl font-bold mb-4">Confirm Order</h1>
      <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-2">{product.name}</h2>
        <p className="text-gray-600 dark:text-gray-300 mb-4">
          {product.description}
        </p>
        <p className="text-2xl font-bold mb-6">
          {product.currency} {(product.price / 100).toFixed(2)}
        </p>
        <div className="mt-6 flex justify-center">
          <PayPalScriptProvider
            options={{
              clientId: process.env.NEXT_PUBLIC_PAYPAL_CLIENT_ID as string,
              currency: product.currency, // Keep if your app supports multiple currencies
            }}
          >
            <PayPalButton
              product={product}
              onSuccess={handlePaymentSuccess}
              onError={handlePaymentError}
            />
          </PayPalScriptProvider>
        </div>
      </div>
    </div>
  );
}
