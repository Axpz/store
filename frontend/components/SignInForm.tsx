'use client';

import { useState } from 'react';
import * as Form from '@radix-ui/react-form';
import { useRouter, useSearchParams } from 'next/navigation';
import { useUser } from '../context/UserContext';
import { apiFetch } from '@/lib/apifetch';

export default function SignInForm() {
  const router = useRouter();
  const { setUser } = useUser();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const searchParams = useSearchParams();
  const callbackUrl = searchParams.get('callbackUrl') || '/dashboard/orders'; // 获取 callbackUrl，默认为首页
  const verified = searchParams.get("verify") === "success";

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);
    setError('');

    const formData = new FormData(event.currentTarget);
    const data = {
      email: formData.get('email'),
      password: formData.get('password'),
      remember: formData.get('remember-me') === 'on',
    };

    try {
      const response = await apiFetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        data: JSON.stringify(data),
      });

      const result = await response;
      
      // 保存用户信息
      const userInfo = {
        username: result.user.username,
        id: result.user.id,
        plan: result.user.plan,
        email: result.user.email,
      };
      setUser(userInfo);
      
      // 如果选择了记住我，在前端也设置一个标记
      data.remember = true; // TODO fix me
      if (data.remember) {
        localStorage.setItem('rememberMe', 'true');
        localStorage.setItem('user', JSON.stringify(userInfo));
      } else {
        localStorage.removeItem('rememberMe');
      }

      // Redirect to home page after successful signin
      router.push(callbackUrl);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Login failed, please try again');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
    {verified ? (
      <div>
        <h2 className="text-green-600 text-center text-2xl font-medium">
          ✅ Email verified, please login
        </h2>
      </div>
    ) : (
      <div>
        <h2 className="mt-6 text-center text-4xl font-extrabold text-gray-900 dark:text-white">
          Login to your account
        </h2>
        <p className="mt-3 text-center text-lg text-gray-600 dark:text-gray-400">
          Don't have an account?{" "}
          <a
            href="/signup"
            className="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400"
          >
            Sign up
          </a>
        </p>
      </div>
    )}

    <Form.Root className="space-y-8" onSubmit={handleSubmit}>
      {error && (
        <div className="p-4 text-base text-red-700 bg-red-100 rounded-lg dark:bg-red-900 dark:text-red-300" role="alert">
          {error}
        </div>
      )}

      <div>
        <Form.Field name="email">
          <div className="flex items-center justify-between">
            <Form.Label className="block text-lg font-medium text-gray-700 dark:text-gray-300">
              Email Address
            </Form.Label>
            <Form.Message className="text-sm text-red-600 dark:text-red-400" match="valueMissing">
              Please enter your email
            </Form.Message>
            <Form.Message className="text-sm text-red-600 dark:text-red-400" match="typeMismatch">
              Please enter a valid email address
            </Form.Message>
          </div>
          <Form.Control asChild>
            <input
              type="email"
              required
              className="mt-2 block w-full px-2 py-2 text-base bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="you@example.com"
            />
          </Form.Control>
        </Form.Field>
      </div>

      <div>
        <Form.Field name="password">
          <div className="flex items-center justify-between">
            <Form.Label className="block text-lg font-medium text-gray-700 dark:text-gray-300">
              Password
            </Form.Label>
            <Form.Message className="text-sm text-red-600 dark:text-red-400" match="valueMissing">
              Please enter your password
            </Form.Message>
          </div>
          <Form.Control asChild>
            <input
              type="password"
              required
              className="mt-2 block w-full px-2 py-2 text-base bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="••••••••"
            />
          </Form.Control>
        </Form.Field>
      </div>

      <div className="flex items-center justify-between">
        <div className="flex items-center">
          <input
            id="remember-me"
            name="remember-me"
            type="checkbox"
            className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            defaultChecked={typeof window !== 'undefined' && localStorage.getItem('rememberMe') === 'true'}
          />
          <label htmlFor="remember-me" className="ml-2 block text-sm text-gray-900 dark:text-gray-300">
            Remember me
          </label>
        </div>

        <div className="text-sm">
          <a href="#" className="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400">
            Forgot your password?
          </a>
        </div>
      </div>

      <Form.Submit asChild>
        <button
          type="submit"
          disabled={isLoading}
          className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-lg font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {isLoading ? 'Logging in...' : 'Login'}
        </button>
      </Form.Submit>
    </Form.Root>
    </>
  );
} 