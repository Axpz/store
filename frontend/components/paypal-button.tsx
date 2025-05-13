'use client';

import React from 'react';
import { PayPalButtons, usePayPalScriptReducer } from '@paypal/react-paypal-js';
import { toast } from 'react-toastify';
import { OrderRequest, Product, OrdersResponse } from '@/lib/api';

import { Order } from '@/lib/api';
import { apiFetch } from '@/lib/apifetch';

interface PayPalButtonProps {
  product: Product; // 更精确的类型
  onSuccess?: () => void;
  onError?: (error: any) => void;
}

const PayPalButtonComponent: React.FC<PayPalButtonProps> = ({ product, onSuccess, onError }) => {
  const [{ isPending }] = usePayPalScriptReducer();
  
  const handleCreateOrder = (data: any, actions: any) => {
    console.log('handleCreateOrder - data:', data);
    console.log('handleCreateOrder - actions:', actions);

    const orderReq: OrderRequest = {
      currency: product.currency as 'CNY' | 'USD',
      products: [{ 
        id: product.id, 
        name: product.name, 
        quantity: 1, 
        price: product.price,
        content: product.content
      }],
      total_amount: product.price
    };
  
    return fetch('/api/orders', {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(orderReq),
    })
      .then((res) => {
        console.log('handleCreateOrder - fetch response (res):', res);
        return res.json(); // 解析响应体为 JSON
      })
      .then((orderResp: Order) => {
        console.log('handleCreateOrder - parsed order data (order):', orderResp);
        console.log('handleCreateOrder - orderID:', orderResp.id);
        return orderResp.id; // 返回 orderID 给 PayPal SDK
      });
  };

  const handleApprove = async (data: any, actions: any) => {
    try {
        const captureResponse = await apiFetch(`/api/orders/${data.orderID}/capture`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const captureData = await captureResponse;
      console.log('handleApprove - captureData:', captureData);

      if (captureData?.status === 'success') {
        toast.success('支付成功！');
        if (onSuccess) {
          onSuccess();
        }
      } else {
        const error = captureData?.error || '支付失败';
        toast.error(`支付失败：${error}`);
        if (onError) {
          onError(error);
        }
      }
    } catch (error: any) {
      toast.error(`支付失败：${error?.error || error || '未知错误'}`);
      if (onError) {
        onError(error);
      }
    }
  };

  const handleError = (err: any) => {
    console.error('PayPal 错误：', err);
    toast.error(`PayPal 错误：${err?.error || '未知错误'}`);
    if (onError) {
      onError(err);
    }
  };

  return (
    <div>
      {isPending && <p>Loading...</p>}
      <PayPalButtons
        createOrder={handleCreateOrder}
        onApprove={handleApprove}
        onError={handleError}
        style={{
          layout: 'vertical',
          color: 'blue',
          shape: 'rect',
          label: 'pay',
        }}
      />
    </div>
  );
};

export default PayPalButtonComponent;
