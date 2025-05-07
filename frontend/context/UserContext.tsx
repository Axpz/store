'use client';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { redirect, usePathname, useRouter } from 'next/navigation';

interface User {
  id: string;
  username: string;
  plan: string;
  email: string;  
}

interface UserContextType {
  user: User | null;
  setUser: (user: User | null) => void;
  logout: () => Promise<void>;
  isLoading: boolean;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export function UserProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (typeof window !== 'undefined') {
      const stored = localStorage.getItem('user');
      if (stored) {
        setUser(JSON.parse(stored));
      }
      setIsLoading(false); // 加上 loading 状态管理
    }
  }, []);

  const logout = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/auth/logout', {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error('Logout failed');
      }

      setUser(null);
      localStorage.removeItem('user');
      localStorage.removeItem('rememberMe');
      redirect('/');
    } catch (error) {
      console.error('Logout error:', error);
      setUser(null);
      localStorage.removeItem('user');
      localStorage.removeItem('rememberMe');
      redirect('/');
    }
  };

  return (
    <UserContext.Provider value={{ user, setUser, logout, isLoading }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
}

export function useAuth({ redirectTo = '/login', shouldRedirect = true } = {}) {
  const { user, isLoading } = useUser();
  const router = useRouter();
  const pathname = usePathname();

  useEffect(() => {
    console.log("User state (from useAuth):", { user, isLoading });

    if (!isLoading && !user && shouldRedirect && pathname !== redirectTo) {
      console.log("User not logged in, redirecting to login with callback:", pathname);
      router.push(`${redirectTo}?callbackUrl=${encodeURIComponent(pathname)}`);
    } else if (!isLoading && user) {
      console.log("User details (from useAuth):", {
        id: user.id,
        username: user.username,
        plan: user.plan
      });
      // 这里可以放置一些登录后的通用逻辑，如果需要的话
    }
  }, [user, isLoading, router, pathname, redirectTo, shouldRedirect]);

  return { user, isLoading }; // Hook 可以返回一些有用的状态
}
