"use client";

import { useAuth } from "@/context/UserContext";
import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";

import OrderItem from "@/components/OrderItem";
import { use, useEffect, useState } from "react";
import { log } from "console";
import { toast } from "react-toastify";
import { Order, OrdersResponse } from "@/lib/api";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Label } from "@radix-ui/react-label";
import { Separator } from "@radix-ui/react-separator";
import { ArrowRightIcon, DotFilledIcon } from "@radix-ui/react-icons";
import Link from "next/link";
import { formatPrice, formatDate } from "@/lib/utils";
import { useOrderStore } from "@/app/store/productStore";

const statusColors: { [key: string]: string } = {
    completed:
      "bg-green-200 text-green-800 dark:bg-green-700 dark:text-green-300 font-semibold", // 更强调完成，略微加深背景和文本，加粗
    paid: "bg-green-100 text-green-700 dark:bg-green-800 dark:text-green-300", // 基础的成功颜色
    pending:
      "bg-yellow-100 text-yellow-700 dark:bg-yellow-800 dark:text-yellow-300",
    shipped: "bg-blue-100 text-blue-700 dark:bg-blue-800 dark:text-blue-300",
    cancelled: "bg-red-100 text-red-700 dark:bg-red-800 dark:text-red-300",
  };

interface OrderPageProps {
  params: { id: string };
}

export default function OrderPage({ params }: OrderPageProps) {
  const { user, isLoading: isAuthLoading } = useAuth();
  const { getOrderById } = useOrderStore();
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [order, setOrder] = useState<Order | null>(null);
  const id = use(Promise.resolve(params?.id));

  useEffect(() => {
    if (!isAuthLoading) {
      if (!user) {
        setError("Please login to view this order");
        setIsLoading(false);
        return;
      }

      try {
        const orderData = getOrderById(id);
        if (!orderData) {
          setError("Order not found");
        } else {
          setOrder(orderData as Order);
        }
      } catch (err) {
        setError("Failed to load order");
      } finally {
        setIsLoading(false);
      }
    }
  }, [isAuthLoading, user, id, getOrderById]);

  if (isAuthLoading || isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-red-500">{error}</div>
      </div>
    );
  }

  if (!order) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-gray-500">Order not found</div>
      </div>
    );
  }

  return (
    <Card className="shadow-md rounded-md p-4 flex flex-col justify-between">
      {/* 使用 flex-col 和 justify-between */}
      <CardHeader className="flex justify-between items-center">
        <Label className="text-lg font-semibold">订单 ID: {order.id}</Label>
        <div
          className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-semibold ${
            statusColors[order.status.toLowerCase()]
          }`}
        >
          <DotFilledIcon className="mr-1" />
          {order.status}
        </div>
      </CardHeader>
      <CardContent className="grid gap-2 mt-2">
        <div>
          <Label className="text-sm text-gray-500">商品:</Label>
          <p className="text-sm">
            {order.products.map((p) => p.name).join(", ")}{" "}
            {order.products.length > 2 ? "..." : ""}
          </p>
        </div>
        <Separator className="my-2 bg-gray-200" />
        <div>
          <Label className="text-sm text-gray-500">总金额:</Label>
          <p className="text-sm font-medium">
            {formatPrice(order.total_amount, order.currency)}
          </p>
        </div>
        <Separator className="my-2 bg-gray-200" />
        <div>
          <Label className="text-sm text-gray-500">创建时间:</Label>
          <p className="text-sm">{formatDate(order.created)}</p>
        </div>
        <div>
          <Label className="text-sm text-gray-500">更新时间:</Label>
          <p className="text-sm">{formatDate(order.updated)}</p>
        </div>
        <Separator className="my-2 bg-gray-200" />
        {/* 查看详情按钮保持在 CardContent 中 */}
      </CardContent>
      {/* 将按钮放在 CardContent 外部，并利用 justify-between 推到底部 */}
      {/* <div className="mt-auto">
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded text-sm w-full">
          查看详情
        </button>
      </div> */}
      {/* <CardFooter>
        <div className="flex justify-center">
          <Link href={`/dashboard/orders/${order.id}`} passHref>
            <button className="inline-flex items-center justify-center bg-white hover:bg-gray-100 text-blue-500 font-bold py-2 px-4 rounded text-sm focus:outline-none focus:ring-1 focus:ring-blue-300 active:outline-none cursor-pointer w-48">
              <ArrowRightIcon className="mr-2 h-4 w-4" aria-hidden="true" />
            </button>
          </Link>
        </div>{" "}
      </CardFooter> */}
    </Card>
  );
}
