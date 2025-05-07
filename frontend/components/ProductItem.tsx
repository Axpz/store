"use client";

import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Label } from "@radix-ui/react-label";
import { Separator } from "@radix-ui/react-separator";
import { Product } from "@/lib/api";
import React, { useState } from "react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useProductStore } from "@/app/store/productStore";
import PayPalButton from "./paypal-button";
import { formatPrice, formatDate } from "@/lib/utils";
import {
  ArrowRightIcon,
  ArrowDownIcon,
} from "@radix-ui/react-icons";
import { Button } from "@/components/ui/button";

export interface ProductItemProps {
  product: Product;
  checkout?: boolean;
}

const ProductItem: React.FC<ProductItemProps> = ({
  product,
  checkout = false,
}) => {
  const router = useRouter();
  const { setSelectedProduct } = useProductStore();
  const [isExpanded, setIsExpanded] = useState(false);

  const handleCardClick = () => {
    if (!checkout) {
      setSelectedProduct(product);
      router.push(`/products/${product.id}`);
    }
  };

  const handleCheckoutClick = () => {
    router.push(`/products/${product.id}/checkout`);
  };

  const toggleExpand = (e: React.MouseEvent) => {
    e.stopPropagation();
    setIsExpanded(!isExpanded);
  };

  return (
    <Card
      className={`group relative overflow-hidden rounded-lg justify-between border border-gray-200 dark:border-gray-800 bg-white dark:bg-black hover:border-gray-300 dark:hover:border-gray-700 transition-all shadow-md ${
        checkout ? "" : "cursor-pointer"
      }`}
      onClick={handleCardClick}
    >
      <CardHeader className="p-0">
        <div className="aspect-[16/9] relative bg-gray-100 dark:bg-gray-900">
          <Image
            src={`/${product.image}`}
            alt={product.name}
            fill
            className="object-contain"
          />
        </div>
        <Label className="text-lg font-semibold">{product.name}</Label>
        <div>
          <Label className="text-sm text-gray-500">{product.type}</Label>
          <p className="text-sm font-medium">
            {formatPrice(product.price, product.currency)}
          </p>
        </div>
      </CardHeader>
      <CardContent className="p-6 grid gap-2 mt-2">
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
          <div className="mt-2 grid gap-2">
            <div>
              <Label className="text-sm text-gray-500">描述:</Label>
              <p className="text-sm">{product.description}</p>
            </div>
            <Separator className="my-2 bg-gray-200" />
            <div>
              <Label className="text-sm text-gray-500">创建时间:</Label>
              <p className="text-sm">{formatDate(product.created)}</p>
            </div>
            <div>
              <Label className="text-sm text-gray-500">更新时间:</Label>
              <p className="text-sm">{formatDate(product.updated)}</p>
            </div>
          </div>
        )}
      </CardContent>
      <CardFooter className="flex justify-center items-center">
        {checkout && 
          <div className="w-full">
            <button
              onClick={handleCheckoutClick}
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded text-sm w-full text-center"
            >
              checkout {formatPrice(product.price, product.currency)}
            </button>
          </div>
        }
      </CardFooter>
    </Card>
  );
};

export default ProductItem;