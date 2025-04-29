import { Card, CardHeader, CardContent } from "@/components/ui/card";
import { Label } from "@radix-ui/react-label";
import { Separator } from "@radix-ui/react-separator";
import { DotFilledIcon } from "@radix-ui/react-icons";
import { format } from "date-fns";
import { Product } from "@/app/dashboard/products/page"; // 假设你定义了 Product 类型
import React from "react";

interface ProductItemProps {
  product: Product;
}

const formatPrice = (price: number, currency: string): string => {
  return `${currency} ${price.toFixed(2)}`;
};

const formatDate = (timestamp: number): string => {
  const date = new Date(timestamp); // 假设 timestamp 已经是毫秒
  return format(date, "yyyy-MM-dd HH:mm");
};

const typeColors: { [key: string]: string } = {
  electronic: "bg-blue-100 text-blue-700 dark:bg-blue-800 dark:text-blue-300",
  clothing: "bg-green-100 text-green-700 dark:bg-green-800 dark:text-green-300",
  book: "bg-yellow-100 text-yellow-700 dark:bg-yellow-800 dark:text-yellow-300",
  // 添加更多商品类型和对应的颜色
};

const ProductItem: React.FC<ProductItemProps> = ({ product }) => {
  return (
    <Card className="shadow-md rounded-md p-4 flex flex-col justify-between">
      <CardHeader className="flex justify-between items-center">
        <Label className="text-lg font-semibold">{product.name}</Label>
        {product.type && typeColors[product.type.toLowerCase()] && (
          <div
            className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-semibold ${
              typeColors[product.type.toLowerCase()]
            }`}
          >
            <DotFilledIcon className="mr-1" />
            {product.type}
          </div>
        )}
      </CardHeader>
      <CardContent className="grid gap-2 mt-2">
        <div>
          <Label className="text-sm text-gray-500">描述:</Label>
          <p className="text-sm">
            {product.description.substring(0, 50)}...
          </p>{" "}
          {/* 截取部分描述 */}
        </div>
        <div>
          <Label className="text-sm text-gray-500">价格:</Label>
          <p className="text-sm font-medium">
            {formatPrice(product.price, product.currency)}
          </p>
        </div>
        <div>
          <Label className="text-sm text-gray-500">创建时间:</Label>
          <p className="text-sm">{formatDate(product.created)}</p>
        </div>
        <div>
          <Label className="text-sm text-gray-500">更新时间:</Label>
          <p className="text-sm">{formatDate(product.updated)}</p>
        </div>
        <Separator className="my-2 bg-gray-200" />
        {/* 可以添加查看详情或操作按钮 */}
      </CardContent>
      <div className="mt-auto">
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded text-sm w-full">
          查看详情
        </button>
      </div>
    </Card>
  );
};

export default ProductItem;
