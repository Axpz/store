import React from 'react';
import { Card, CardHeader, CardContent } from "@/components/ui/card";
import { Label } from '@radix-ui/react-label';
import { Separator } from '@radix-ui/react-separator';
import { DotFilledIcon } from '@radix-ui/react-icons';
import { Order } from '@/lib/api';
import { formatPrice, formatDate } from '@/lib/utils'; 

// const statusColors: { [key: string]: string } = {
//   pending: 'bg-yellow-100 text-yellow-700',
//   paid: 'bg-green-100 text-green-700',
//   shipped: 'bg-blue-100 text-blue-700',
//   completed: 'bg-gray-100 text-gray-700',
//   cancelled: 'bg-red-100 text-red-700',
// };

interface OrderItemProps {
  order: Order;
}

const statusColors: { [key: string]: string } = {
  completed: 'bg-green-200 text-green-800 dark:bg-green-700 dark:text-green-300 font-semibold', // 更强调完成，略微加深背景和文本，加粗
  paid: 'bg-green-100 text-green-700 dark:bg-green-800 dark:text-green-300', // 基础的成功颜色
  pending: 'bg-yellow-100 text-yellow-700 dark:bg-yellow-800 dark:text-yellow-300',
  shipped: 'bg-blue-100 text-blue-700 dark:bg-blue-800 dark:text-blue-300',
  cancelled: 'bg-red-100 text-red-700 dark:bg-red-800 dark:text-red-300',
};


const OrderItem: React.FC<OrderItemProps> = ({ order }) => {
  return (
    <Card className="shadow-md rounded-md p-4 flex flex-col justify-between"> {/* 使用 flex-col 和 justify-between */}
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
        <Separator className="my-2 bg-gray-200" />
        <div>
          <Label className="text-sm text-gray-500">总金额:</Label>
          <p className="text-sm font-medium">{formatPrice(order.total_amount, order.currency)}</p>
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
    </Card>
  );
};

export default OrderItem;