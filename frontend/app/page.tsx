'use client';

import Image from "next/image";
import Header from "@/components/Header";

export default function Home() {
  const templates = [
    {
      title: "Dashboard Template",
      description: "A modern dashboard layout with charts and data visualization",
      image: "/dashboard-template.png"
    },
    {
      title: "Authentication Forms",
      description: "Beautiful sign-in and sign-up forms with validation",
      image: "/auth-template.png"
    },
    {
      title: "E-commerce Store",
      description: "Complete e-commerce interface with product listings and cart",
      image: "/ecommerce-template.png"
    },
    {
      title: "Blog Layout",
      description: "Clean and minimal blog design with rich text editing",
      image: "/blog-template.png"
    }
  ];

  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-center text-center max-w-7xl mx-auto">
        <main className="flex flex-col items-center gap-12 w-full">
          <div className="text-center max-w-4xl">
            <h1 className="text-[32px] font-bold leading-tight tracking-tight">
              Radix UI templates and examples
            </h1>
            <p className="text-xl text-gray-600 dark:text-gray-400 mt-6">
              Browse examples and templates of application builds using Radix UI.
            </p>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full">
            {templates.map((template, index) => (
              <div key={index} className="group relative overflow-hidden rounded-lg border border-gray-200 dark:border-gray-800 bg-white dark:bg-black hover:border-gray-300 dark:hover:border-gray-700 transition-all">
                <div className="aspect-[16/9] relative bg-gray-100 dark:bg-gray-900">
                  <Image
                    src={template.image}
                    alt={template.title}
                    fill
                    className="object-cover"
                  />
                </div>
                <div className="p-6">
                  <h3 className="text-xl font-semibold">{template.title}</h3>
                  <p className="mt-2 text-lg text-gray-600 dark:text-gray-400">
                    {template.description}
                  </p>
                </div>
                <div className="absolute inset-0 pointer-events-none border-2 border-transparent group-hover:border-black dark:group-hover:border-white transition-all rounded-lg"></div>
              </div>
            ))}
          </div>
        </main>
      </div>
    </>
  );
}
