import SignInForm from '@/components/SignInForm';

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-4xl font-extrabold text-gray-900 dark:text-white">
            Login to your account
          </h2>
          <p className="mt-3 text-center text-lg text-gray-600 dark:text-gray-400">
            Don't have an account?{' '}
            <a href="/signup" className="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400">
              Sign up
            </a>
          </p>
        </div>
        <SignInForm />
      </div>
    </div>
  );
} 