import React from 'react';
import { Card, CardHeader, CardContent } from '@/components/ui/card';
import { Label } from '@radix-ui/react-label';
import { Separator } from '@radix-ui/react-separator';
import { DotFilledIcon } from '@radix-ui/react-icons';

interface OrderProduct {
  id: string;
  name: string;
}

interface Order {
  id: string;
  user_id: string;
  status: string;
  currency: string;
  products: OrderProduct[];
  total_amount: number;
  paid_amount: number;
  description: string;
  created: number;
  updated: number;
}

const statusColors: { [key: string]: string } = {
  pending: 'bg-yellow-100 text-yellow-700',
  paid: 'bg-green-100 text-green-700',
  shipped: 'bg-blue-100 text-blue-700',
  completed: 'bg-gray-100 text-gray-700',
  cancelled: 'bg-red-100 text-red-700',
};

interface OrderItemProps {
  order: Order;
}

const OrderItem: React.FC<OrderItemProps> = ({ order }) => {
  const formatDate = (timestamp: number): string => {
    const date = new Date(timestamp * 1000);
    return date.toLocaleDateString();
  };

  const formatAmount = (amount: number): string => {
    return (amount / 100).toFixed(2);
  };

  return (
    <Card className="shadow-md rounded-md p-4">
      <CardHeader className="flex justify-between items-center">
        <Label className="text-lg font-semibold">订单 ID: {order.id}</Label>
        <div className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-semibold ${statusColors[order.status.toLowerCase()]}`}>
          <DotFilledIcon className="mr-1" />
          {order.status}
        </div>
      </CardHeader>
      <CardContent className="grid gap-2 mt-2">
        <div>
          <Label className="text-sm text-gray-500">商品:</Label>
          <p className="text-sm">{order.products.map(p => p.name).join(', ')} {order.products.length > 2 ? '...' : ''}</p>
        </div>
        <div>
          <Label className="text-sm text-gray-500">总金额:</Label>
          <p className="text-sm font-medium">{order.currency} {formatAmount(order.total_amount)}</p>
        </div>
        <div>
          <Label className="text-sm text-gray-500">创建时间:</Label>
          <p className="text-sm">{formatDate(order.created)}</p>
        </div>
        <Separator className="my-2 bg-gray-200" />
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded text-sm">
          查看详情
        </button>
      </CardContent>
    </Card>
  );
};

export default OrderItem;