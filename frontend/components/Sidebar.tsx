import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Home, Menu, ShoppingCart, Users, X, Package, BarChart, Settings } from 'lucide-react';
import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { usePathname, useRouter } from 'next/navigation';
import { cn } from '@/lib/utils'; // 假设你有这个 utility 函数

interface NavItem {
  label: string;
  value: string;
  icon: React.ReactNode;
  href?: string;
}

const navItems: NavItem[] = [
  {
    label: 'User Management',
    value: 'users',
    icon: <Users className="w-5 h-5" />,
    href: '/dashboard/users', // 示例链接
  },
  {
    label: 'Order Management',
    value: 'orders',
    icon: <ShoppingCart className="w-5 h-5" />,
    href: '/dashboard/orders', // 示例链接
  },
  {
    label: 'Products',
    value: 'products',
    icon: <Package className="w-5 h-5" />,
    href: '/dashboard/products', // 示例链接
  },
  {
    label: 'Analytics',
    value: 'analytics',
    icon: <BarChart className="w-5 h-5" />,
    href: '/dashboard/analytics', // 示例链接
  },
  {
    label: 'Settings',
    value: 'settings',
    icon: <Settings className="w-5 h-5" />,
    href: '/dashboard/settings', // 示例链接
  },
  // 你可以根据你的应用添加更多的导航项
];

interface SidebarProps {
  className?: string;
  // navItems: NavItem[];
}

const Sidebar: React.FC<SidebarProps> = ({ className }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const pathname = usePathname();
  const router = useRouter();
  const activeTab = pathname;

  const handleNavigation = (value: string) => {
    setIsSidebarOpen(false);
    // 如果有 href，使用 href 进行跳转，否则只更新 activeTab 状态
    const selectedItem = navItems.find(item => item.value === value);
    if (selectedItem?.href) {
      router.push(selectedItem.href);
    }
    // 如果你还需要在没有 href 时更新 activeTab，可以保留或修改下面的逻辑
    // setActiveTab(value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLAnchorElement>) => {
    if (event.key === 'Enter' || event.key === ' ') {
      const value = event.currentTarget.getAttribute('data-value');
      if (value) {
        handleNavigation(value);
      }
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
        className={cn(
          `
            fixed lg:relative left-0 z-40
            w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700
            transform transition-transform duration-300 ease-in-out
            ${isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'}
            top-4
            h-[calc(100vh - 4rem)]
            will-change-transform
          `,
          className
        )}
        orientation="vertical"
      >
        <NavigationMenu.List className="flex flex-col h-full">
          <div className="p-4 border-b border-gray-200 dark:border-gray-700">
            <h1 className="text-2xl font-bold text-gray-800 dark:text-white">Dashboard</h1>
          </div>

          <div className="flex-1 overflow-y-auto p-4 flex flex-col gap-y-4">
            {navItems.map((item) => (
              <NavigationMenu.Item key={item.value}>
                <NavigationMenu.Link
                  data-value={item.value} // 添加 data-value 方便键盘导航
                  className={cn(
                    `
                      flex items-center gap-3 w-full px-4 py-2 text-left rounded-md transition-colors
                      focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2
                    `,
                    activeTab === item.value
                      ? 'bg-blue-100 text-blue-600 dark:bg-blue-900 dark:text-blue-300'
                      : 'text-gray-600 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700'
                  )}
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
};

export default Sidebar;