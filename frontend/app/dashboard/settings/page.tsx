'use client';

import { useAuth } from '@/context/UserContext';
import Header from '@/components/Header';
import Sidebar from '@/components/Sidebar';

export default function SettingsPage() {
  const { user, isLoading } = useAuth(); 

  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-start text-center max-w-7xl mx-auto">
        <div className="flex w-full">
          <Sidebar className="mr-8" />
          
          {/* Main Content */}
          <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
            <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">Setting Management</h2>
            {/* User management content will go here */}
          </div>
        </div>
      </div>
    </>
  );
} 