"use client";

import React from "react";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import {
  ArrowRightIcon,
  DotFilledIcon,
  ArrowDownIcon,
} from "@radix-ui/react-icons";
import { Order } from "@/lib/api";
import { formatPrice, formatDate } from "@/lib/utils";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";

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
    <Card className="group relative overflow-hidden rounded-lg justify-between border border-gray-200 dark:border-gray-800 bg-white dark:bg-black hover:border-gray-300 dark:hover:border-gray-700 transition-all shadow-md">
      <CardHeader className="flex justify-between items-center p-4">
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
      <CardContent className="p-4 grid gap-2">
        <div className="flex flex-col items-center"> {/* Center商品 */}
          <Label className="text-sm text-gray-500">商品:</Label>
          <p className="text-sm text-center"> {/* Center商品文本 */}
            {order.products.map((p) => p.name).join(", ")}
            {order.products.length > 2 ? "..." : ""}
          </p>
        </div>
        <div className="flex flex-col items-center"> {/* Center总金额 */}
          <Label className="text-sm text-gray-500">总金额:</Label>
          <p className="text-sm font-medium">
            {formatPrice(order.total_amount, order.currency)}
          </p>
        </div>

        <Button
          variant="ghost"
          onClick={toggleExpand}
          className="inline-flex items-center justify-center text-blue-500 hover:text-blue-700 hover:bg-blue-50 dark:hover:bg-blue-900/20 w-48 mt-2"
        >
          {isExpanded ? (
            <ArrowDownIcon className="mr-2 h-4 w-4" aria-hidden="true" />
          ) : (
            <ArrowRightIcon className="mr-2 h-4 w-4" aria-hidden="true" />
          )}
        </Button>

        {isExpanded && (
          <div className="mt-2 grid gap-4">
            {order.products.map((p) => (
              <div key={p.id} className="space-y-2">
                <p className="font-semibold">{p.name}</p>
                {order.status === "completed" && p.content && p.content.length > 0 && (
                  <div className="mt-1 flex flex-col gap-1">
                    {p.content.map((c, i) => (
                      <Link
                        href={`/products/${p.id}/${i}`}
                        key={c}
                        className="text-sm text-blue-600 hover:text-blue-800 hover:underline transition-colors"
                      >
                        {c}
                      </Link>
                    ))}
                  </div>
                )}
              </div>
            ))}

            <div className="space-y-2 flex flex-col items-center">
              <div>
                <Label className="text-sm text-gray-500">创建时间 {formatDate(order.created)}</Label>
              </div>
              <div>
                <Label className="text-sm text-gray-500">更新时间 {formatDate(order.updated)}</Label>
              </div>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default OrderItem;