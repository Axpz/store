export interface Product {
    id: string;
    name: string;
    image: string;
    type: string;
    description: string;
    price: number; // 转换成 number 类型
    currency: string;
    created: number; // 转换成 number 类型
    updated: number; // 转换成 number 类型
    status: string; // 状态
  }
  
  export interface ProductsResponse {
    data: Product[];
    page: number;
    size: number;
    total: number;
  }

  export interface Order {
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
  
  export interface OrderProduct {
    id: string;
    name: string;
    quantity: number;
    price: number;
  }
  
  export interface OrderRequest {
    currency: 'CNY' | 'USD'; // 货币类型
    products: OrderProduct[]; // 订单商品
    total_amount?: number;    // 总金额，单位分 (可选，因为后端可能会计算)
    description?: string;     // 订单描述 (可选)
  }
  
  export interface OrdersResponse {
    data: Order[];
    page: number;
    size: number;
    total: number;
  }