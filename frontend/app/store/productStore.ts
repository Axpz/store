import { create } from 'zustand';
import { Order, Product } from '@/lib/api'; // 确保导入你的 Product 类型

interface ProductState {
  productsArray: Product[];
  selectedProduct: Product | null;
  setProducts: (products: Product[]) => void;
  setSelectedProduct: (product: Product | null) => void;
}

export const useProductStore = create<ProductState>((set) => ({
  productsArray: [],
  selectedProduct: null,
  setProducts: (products) => set({ productsArray: products }),
  setSelectedProduct: (product) => set({ selectedProduct: product }),
}));


interface OrderState {
  ordersArray: Order[];
  ordersMap: Map<string, Order>;
  setOrders: (orders: Order[]) => void;
  getOrderById: (id: string) => Order | undefined;
}

export const useOrderStore = create<OrderState>()((set, get) => ({
  ordersArray: [],
  ordersMap: new Map(),
  setOrders: (orders) => {
    if (!Array.isArray(orders)) {
      // if null、undefined, set to Map
      set({
        ordersArray: [],
        ordersMap: new Map(),
      });
      return;
    }
  
    const newOrdersMap = new Map(orders.map(order => [order.id, order]));
  
    set({
      ordersArray: orders,
      ordersMap: newOrdersMap,
    });
  },
  getOrderById: (id) => get().ordersMap.get(id),
}));