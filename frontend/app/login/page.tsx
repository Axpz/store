"use client";

import { Suspense } from "react";
import { useSearchParams } from "next/navigation";
import Header from "@/components/Header";
import SignInForm from "@/components/SignInForm";

export default function LoginPage() {
  const params = useSearchParams();
  const verified = params.get("verify") === "success";

  return (
    <>
      <Header />
      <div className="min-h-screen pt-24 flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
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

          <Suspense fallback={<div>Loading...</div>}>
            <SignInForm />
          </Suspense>
        </div>
      </div>
    </>
  );
}
