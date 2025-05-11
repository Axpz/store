'use client';

import { useState } from 'react';
import * as Form from '@radix-ui/react-form';
import { useRouter } from 'next/navigation';

export default function SignUpForm() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);
    setError('');

    const formData = new FormData(event.currentTarget);
    const data = {
      email: formData.get('email'),
      password: formData.get('password'),
      username: formData.get('username'),
    };

    try {
      const response = await fetch('http://localhost:8080/api/auth/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to sign up');
      }

      // Redirect to login page after successful signup
      router.push('/login');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Sign up failed, please try again');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
    <div>
      <h2 className="mt-6 text-center text-4xl font-extrabold text-gray-900 dark:text-white">
        Create your account
      </h2>
      <p className="mt-3 text-center text-lg text-gray-600 dark:text-gray-400">
        Already have an account?{' '}
        <a href="/login" className="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400">
          Sign in
        </a>
      </p>
    </div>
    <Form.Root className="space-y-8" onSubmit={handleSubmit}>
      {error && (
        <div className="p-4 text-lg text-red-700 bg-red-100 rounded-lg dark:bg-red-900 dark:text-red-300" role="alert">
          {error}
        </div>
      )}

      <div>
        <Form.Field name="username">
          <div className="flex items-center justify-between">
            <Form.Label className="block text-lg font-medium text-gray-700 dark:text-gray-300">
              Username
            </Form.Label>
            <Form.Message className="text-sm text-red-600 dark:text-red-400" match="valueMissing">
              Please enter your username
            </Form.Message>
          </div>
          <Form.Control asChild>
            <input
              type="text"
              required
              className="mt-2 block w-full px-2 py-2 text-base bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your name"
            />
          </Form.Control>
        </Form.Field>
      </div>

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
              Please enter a password
            </Form.Message>
          </div>
          <Form.Control asChild>
            <input
              type="password"
              required
              className="mt-2 block w-full px-2 py-2 text-base bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="••••••••"
              minLength={8}
            />
          </Form.Control>
          <p className="mt-2 text-sm text-gray-500">Password must be at least 8 characters</p>
        </Form.Field>
      </div>

      <Form.Submit asChild>
        <button
          type="submit"
          disabled={isLoading}
          className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-lg font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {isLoading ? 'Signing up...' : 'Sign up'}
        </button>
      </Form.Submit>
    </Form.Root>
    </>
  );
} 