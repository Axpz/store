"use client";

import { PayPalScriptProvider } from "@paypal/react-paypal-js";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import PayPalButton from "@/components/paypal-button";
import { Product } from "@/lib/api";
import { useProductStore } from "@/app/store/productStore";

export default function CheckoutPage({
  params,
}: {
  params: { productId: string };
}) {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const { selectedProduct } = useProductStore();
  if (!selectedProduct) {
    toast.error("No product selected");
    return;
  }

  const product = selectedProduct;

  // useEffect(() => {
  //   const fetchProduct = async () => {
  //     try {
  //       const response = await fetch(`/api/products/${params.productId}`);
  //       if (!response.ok) {
  //         throw new Error('Failed to fetch product');
  //       }
  //       const data = await response.json();
  //       setProduct(data);
  //     } catch (error) {
  //       toast.error('获取商品信息失败');
  //       router.push('/products');
  //     } finally {
  //       setLoading(false);
  //     }
  //   };

  //   fetchProduct();
  // }, [params.productId, router]);

  const handlePaymentSuccess = () => {
    toast.success("支付成功！");
    router.push("/dashboard/orders");
  };

  const handlePaymentError = (error: any) => {
    console.error("支付错误：", error);
  };

  if (loading) {
    return <div>加载中...</div>;
  }

  if (!product) {
    return <div>商品不存在</div>;
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
