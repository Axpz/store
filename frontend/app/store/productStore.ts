import { create } from 'zustand';
import { Product } from '@/lib/api'; // 确保导入你的 Product 类型

interface ProductState {
  selectedProduct: Product | null;
  setSelectedProduct: (product: Product | null) => void;
}

export const useProductStore = create<ProductState>()((set) => ({
  selectedProduct: null,
  setSelectedProduct: (product) => set({ selectedProduct: product }),
}));