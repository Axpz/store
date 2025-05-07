import { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'standalone', // generate minimum running files for container deployment
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'images.credly.com',
        pathname: '/**',
      },
    ],
  },
  reactStrictMode: true,
};

export default nextConfig;
