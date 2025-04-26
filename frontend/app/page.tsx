import Image from "next/image";

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
      <nav className="fixed top-0 left-0 right-0 h-16 border-b border-gray-200 dark:border-gray-800 bg-white dark:bg-black z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-full flex items-center justify-between">
          <div className="flex items-center">
            <a href="/" className="flex items-center">
              <span className="text-xl font-bold">Axpz</span>
            </a>
          </div>
          <div className="hidden md:flex items-center space-x-8">
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Products</a>
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Solutions</a>
            {/* <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Resources</a>
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Enterprise</a> */}
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Docs</a>
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Pricing</a>
          </div>
          <div className="flex items-center space-x-4">
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Log In</a>
            <a href="#" className="text-sm text-gray-600 hover:text-black dark:text-gray-400 dark:hover:text-white">Contact</a>
            <a href="/signup" className="bg-black text-white dark:bg-white dark:text-black px-4 py-2 rounded-md text-sm font-medium">Sign Up</a>
          </div>
        </div>
      </nav>
      <div className="min-h-screen pt-24 p-8 flex flex-col items-center justify-center text-center max-w-7xl mx-auto">
        <main className="flex flex-col items-center gap-12 w-full">
          <div className="text-center max-w-4xl">
            <h1 className="text-[64px] font-bold leading-tight tracking-tight">
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
                  <h3 className="text-lg font-semibold">{template.title}</h3>
                  <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">
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
