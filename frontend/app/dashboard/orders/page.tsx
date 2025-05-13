"use client";

import { useAuth } from "@/context/UserContext";
import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";

import OrderItem from "@/components/OrderItem";
import { useEffect, useState } from "react";
import { log } from "console";
import { toast } from "react-toastify";
import { Order, OrdersResponse } from "@/lib/api";
import { useOrderStore } from "@/app/store/productStore";
import { apiFetch } from "@/lib/apifetch";

export default function OrdersPage() {
  const { user, isLoading } = useAuth();

  const { ordersArray, setOrders } = useOrderStore();

  const fetchOrders = async () => {
    try {
      console.log("fetchOrders");
      const response = await apiFetch("/api/orders", {
        headers: {
          "Content-Type": "application/json",
        },
      });

      const respOrders: OrdersResponse = await response;
      setOrders(respOrders.data);
    } catch (error: any) {
      toast.error(`获取订单失败: ${error.message}`);
    }
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  return (
    <>
      {/* Main Content */}
      <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">
          Order Management
        </h2>

        {ordersArray == null || ordersArray.length === 0 ? (
          <div className="text-center text-gray-500">No orders yet</div>
        ) : (
          <div className="grid gap-4 grid-cols-auto-fit">
            {ordersArray.map((order) => (
              <OrderItem key={order.id} order={order} />
            ))}
          </div>
        )}
      </div>
    </>
  );
}
