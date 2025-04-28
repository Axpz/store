'use client';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { usePathname, useRouter } from 'next/navigation';
import { useState, useEffect } from 'react';
import { Users, ShoppingCart, Menu, X } from 'lucide-react';
import { Button } from './ui/button';

interface SidebarProps {
  className?: string;
}

interface NavItem {
  label: string;
  value: string;
  icon: React.ReactNode;
}

const navItems: NavItem[] = [
  {
    label: 'User Management',
    value: 'users',
    icon: <Users className="w-5 h-5" />,
  },
  {
    label: 'Order Management',
    value: 'orders',
    icon: <ShoppingCart className="w-5 h-5" />,
  },
];

export default function Sidebar({ className = '' }: SidebarProps) {
  const router = useRouter();
  const pathname = usePathname();
  const [activeTab, setActiveTab] = useState<string>('');
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  useEffect(() => {
    const currentTab = pathname.includes('users') ? 'users' : 'orders';
    setActiveTab(currentTab);
  }, [pathname]);

  const handleNavigation = (value: string) => {
    setActiveTab(value);
    router.push(`/dashboard/${value}`);
    setIsSidebarOpen(false);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLAnchorElement>) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      const target = e.currentTarget as HTMLAnchorElement;
      target.click();
    }
  };

  return (
    <>
      {/* Mobile menu button */}
      <Button
        variant="ghost"
        size="icon"
        className="lg:hidden fixed top-4 left-4 z-50"
        onClick={() => setIsSidebarOpen(!isSidebarOpen)}
      >
        {isSidebarOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
      </Button>

      {/* Sidebar */}
      <NavigationMenu.Root 
        className={`
          fixed lg:relative inset-y-0 left-0 z-40
          w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700
          transform transition-transform duration-300 ease-in-out
          ${isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'}
          ${className}
        `}
        orientation="vertical"
      >
        <NavigationMenu.List className="flex flex-col h-full">
          <div className="p-4 border-b border-gray-200 dark:border-gray-700">
            <h1 className="text-2xl font-bold text-gray-800 dark:text-white">Dashboard</h1>
          </div>
          
          <div className="flex-1 overflow-y-auto p-4">
            {navItems.map((item) => (
              <NavigationMenu.Item key={item.value}>
                <NavigationMenu.Link
                  className={`
                    flex items-center gap-3 w-full px-4 py-2 text-left rounded-md transition-colors
                    focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2
                    ${activeTab === item.value
                      ? 'bg-blue-100 text-blue-600 dark:bg-blue-900 dark:text-blue-300'
                      : 'text-gray-600 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700'
                    }
                  `}
                  onClick={() => handleNavigation(item.value)}
                  onKeyDown={handleKeyDown}
                  tabIndex={0}
                  aria-current={activeTab === item.value ? 'page' : undefined}
                >
                  {item.icon}
                  <span>{item.label}</span>
                </NavigationMenu.Link>
              </NavigationMenu.Item>
            ))}
          </div>
        </NavigationMenu.List>
      </NavigationMenu.Root>

      {/* Overlay for mobile */}
      {isSidebarOpen && (
        <div 
          className="fixed inset-0 bg-black bg-opacity-50 z-30 lg:hidden"
          onClick={() => setIsSidebarOpen(false)}
        />
      )}
    </>
  );
} 