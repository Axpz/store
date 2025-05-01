// app/dashboard/layout.tsx
'use client';

import Header from '@/components/Header';
import Sidebar from '@/components/Sidebar';

interface DashboardLayoutProps {
  children: React.ReactNode;
}

export default function DashboardLayout({ children }: DashboardLayoutProps) {
  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-start text-center max-w-7xl mx-auto">
        <div className="flex w-full">
          {/* <Sidebar className="mr-8" /> */}

          {children}
          
        </div>
      </div>
    </>
  );
} 