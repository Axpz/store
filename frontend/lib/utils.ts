import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { format } from 'date-fns';
 
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
} 

export const formatDate = (timestamp: number): string => {
  const date = new Date(timestamp * 1000); // 将秒转换为毫秒
  return format(date, 'yyyy-MM-dd HH:mm');
};

export const formatPrice = (price: number, currency: string): string => {
  const priceInDecimal = price / 100;
  return `${currency} ${priceInDecimal.toFixed(2)} ${
    currency === "CNY" ? "人民币" : "$"
  }`;
};