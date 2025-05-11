import Header from '@/components/Header';
import SignUpForm from '@/components/SignUpForm';

export default function SignUpPage() {
  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <SignUpForm />
        </div>
      </div>
    </>
  );
} 