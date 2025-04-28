'use client';

import { useState } from 'react';
import { useUser } from '@/context/UserContext';

export default function Header() {
  const { user, logout, isLoading } = useUser();
  const [isLoggingOut, setIsLoggingOut] = useState(false);
  const [logoutError, setLogoutError] = useState('');

  const handleLogout = async () => {
    try {
      setIsLoggingOut(true);
      setLogoutError('');
      await logout();
    } catch (error) {
      setLogoutError('Failed to logout. Please try again.');
    } finally {
      setIsLoggingOut(false);
    }
  };

  if (isLoading) {
    return null; // 正在加载时不渲染，避免 hydration 问题
  }

  return (
    <nav className="fixed top-0 left-0 right-0 h-16 border-b border-gray-200 dark:border-gray-800 bg-white dark:bg-black z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-full flex items-center justify-between">
        <div className="flex items-center">
          <a href="/" className="flex items-center">
            <span className="text-xl font-bold">Axpz</span>
          </a>
        </div>
        <div className="hidden md:flex items-center space-x-8">
          {user && (
            <a href="/dashboard/orders" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">
              Dashboard
            </a>
          )}
          <a href="#" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Solutions</a>
          <a href="#" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Docs</a>
          <a href="#" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Pricing</a>
        </div>
        <div className="flex items-center space-x-4">
          {user ? (
            <div className="flex items-center space-x-4">
              <span className="text-lg text-gray-600 dark:text-gray-400">{user.username}</span>
              <button
                onClick={handleLogout}
                disabled={isLoggingOut}
                className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isLoggingOut ? 'Logging out...' : 'Logout'}
              </button>
              {logoutError && (
                <span className="text-lg text-red-600 dark:text-red-400">{logoutError}</span>
              )}
            </div>
          ) : (
            <>
              <a href="/login" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Log In</a>
              <a href="#" className="text-lg text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Contact</a>
              <a href="/signup" className="bg-black text-white dark:bg-white dark:text-black px-4 py-2 rounded-md text-lg font-medium">Sign Up</a>
            </>
          )}
        </div>
      </div>
    </nav>
  );
}
