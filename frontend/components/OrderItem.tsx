'use client';

import React from "react";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Label } from "@radix-ui/react-label";
import { Separator } from "@radix-ui/react-separator";
import { ArrowRightIcon, DotFilledIcon, ArrowDownIcon } from "@radix-ui/react-icons";
import { Order } from "@/lib/api";
import { formatPrice, formatDate } from "@/lib/utils";
import Link from "next/link";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from "@radix-ui/react-dialog";
import { ScrollArea } from '@radix-ui/react-scroll-area';

interface OrderItemProps {
  order: Order;
}

const statusColors: { [key: string]: string } = {
  completed:
    "bg-green-200 text-green-800 dark:bg-green-700 dark:text-green-300 font-semibold",
  paid: "bg-green-100 text-green-700 dark:bg-green-800 dark:text-green-300",
  pending:
    "bg-yellow-100 text-yellow-700 dark:bg-yellow-800 dark:text-yellow-300",
  shipped: "bg-blue-100 text-blue-700 dark:bg-blue-800 dark:text-blue-300",
  cancelled: "bg-red-100 text-red-700 dark:bg-red-800 dark:text-red-300",
};

const OrderItem: React.FC<OrderItemProps> = ({ order }) => {
  const [isExpanded, setIsExpanded] = React.useState(false);

  const toggleExpand = () => {
    setIsExpanded(!isExpanded);
  };

  return (
    <Card className="shadow-md rounded-md p-4 flex flex-col justify-between">
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
          <Label className="text-sm">商品: {order.products.map((p) => p.name).join(", ")}{" "}
          {order.products.length > 2 ? "..." : ""}</Label>
        </div>
        <div>
          <Label className="text-sm">总金额: {formatPrice(order.total_amount, order.currency)}</Label>
        </div>

        <button
          onClick={toggleExpand}
          className="inline-flex items-center justify-center bg-white hover:bg-gray-100 text-blue-500 font-bold py-2 px-4 rounded text-sm focus:outline-none focus:ring-1 focus:ring-blue-300 active:outline-none cursor-pointer w-48 mt-2"
        >
          {isExpanded ? (
            <ArrowDownIcon className="mr-2 h-4 w-4" aria-hidden="true" />
          ) : (
            <ArrowRightIcon className="mr-2 h-4 w-4" aria-hidden="true" />
          )}
        </button>

        {isExpanded && (
          <div className="mt-2 grid gap-2 text-sm">
            {order.products.map((p) => (
              <div key={p.id}>
                <p>{p.name}</p>
                <p>{p.content.map((c) => (
                  <p key={c}>{c}</p>
                ))}</p>
              </div>
            ))}
            
            <div>
              <Label>创建时间</Label>
              <p>{formatDate(order.created)}</p>
            </div>
            <div>
              <Label>更新时间</Label>
              <p>{formatDate(order.updated)}</p>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default OrderItem;