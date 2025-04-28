'use client';

import { redirect } from 'next/navigation';
import { useUser } from '@/context/UserContext';
import Header from '@/components/Header';
import Sidebar from '@/components/Sidebar';

import OrderItem from '@/components/OrderItem';
import { useEffect } from 'react';

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

// 模拟订单数据
const ordersData: Order[] = [
  {
    id: 'ORDER123450',
    user_id: 'tkK0IXs0sejTvZFfxlxEUg',
    status: 'paid',
    currency: 'USD',
    products: [{ id: 'gcp_1001', name: 'Google Cloud Solution Architect' }],
    total_amount: 1,
    paid_amount: 1,
    description: 'Google Cloud Solution Architect',
    created: 1682678400,
    updated: 1682678460,
  },  
  {
    id: 'ORDER123451',
    user_id: 'tkK0IXs0sejTvZFfxlxEUg',
    status: 'paid',
    currency: 'USD',
    products: [{ id: 'gcp_1001', name: 'Google Cloud Solution Architect' }, { id: 'gcp_1002', name: 'Google AI Solution Architect' }],
    total_amount: 2,
    paid_amount: 2,
    description: 'Google Cloud Solution Architect and Google AI Solution Architect',
    created: 1682678400,
    updated: 1682678460,
  },  
  {
    id: 'ORDER12345',
    user_id: 'tkK0IXs0sejTvZFfxlxEUg',
    status: 'paid',
    currency: 'CNY',
    products: [{ id: 'P001', name: '商品 A' }, { id: 'P002', name: '商品 B' }],
    total_amount: 15000,
    paid_amount: 15000,
    description: '测试订单 1',
    created: 1682678400,
    updated: 1682678460,
  },
  {
    id: 'ORDER12346',
    user_id: 'tkK0IXs0sejTvZFfxlxEUg',
    status: 'pending',
    currency: 'CNY',
    products: [{ id: 'P003', name: '商品 C' }],
    total_amount: 5000,
    paid_amount: 0,
    description: '测试订单 2',
    created: 1682592000,
    updated: 1682592060,
  },
  {
    id: 'ORDER11223',
    user_id: 'USER002',
    status: 'shipped',
    currency: 'USD',
    products: [{ id: 'P004', name: '商品 D' }, { id: 'P005', name: '商品 E' }, { id: 'P006', name: '商品 F' }],
    total_amount: 20000,
    paid_amount: 20000,
    description: '测试订单 3',
    created: 1682505600,
    updated: 1682505660,
  },
];

export default function OrdersPage() {
  const { user, isLoading } = useUser();

  useEffect(() => {
    if (!isLoading && !user) {
      redirect('/login');
    }
  }, [user]);

  if (!user) {
    return null;
  }


  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-center text-center max-w-7xl mx-auto">
        <div className="flex w-full">
          <Sidebar className="mr-8" />
          
          {/* Main Content */}
          <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
            <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">Order Management</h2>
            {/* Order management content will go here */}
            <div className="container mx-auto py-8">
              <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                {ordersData.map((order) => (
                  <OrderItem key={order.id} order={order} />
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
} 