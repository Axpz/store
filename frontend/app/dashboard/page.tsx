'use client';

import { redirect } from 'next/navigation';
import { useUser } from '@/context/UserContext';
import Header from '@/components/Header';
import Sidebar from '@/components/Sidebar';
import { useEffect } from 'react';

export default function DashboardPage() {
  const { user, isLoading } = useUser();

  useEffect(() => {
    console.log("User state:", { user, isLoading });
    if (!isLoading && !user) {
      redirect('/login');
    } else if (!isLoading && user) {
      console.log("User details:", {
        id: user.id,
        username: user.username,
        plan: user.plan

      });
      // fetchOrders();
    }
  }, [user, isLoading]);

  if (!user) {
    return null;
  }

  return (
    <>
      <Header />
      <div className="min-h-screen flex flex-col items-center justify-start text-center max-w-7xl mx-auto">
        <div className="flex w-full">
          <Sidebar className="mr-4" />
          
          {/* Main Content */}
          <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-6">
            <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">Welcome to Dashboard</h2>
            <p className="text-gray-600 dark:text-gray-400">
              Select a section from the sidebar to get started.
            </p>
          </div>
        </div>
      </div>
    </>
  );
} 