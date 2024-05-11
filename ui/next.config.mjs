/** @type {import('next').NextConfig} */
const nextConfig = {
    //useEffect'in iki defa çalışmasını engeller
    reactStrictMode: false,
    eslint: {
      // Warning: This allows production builds to successfully complete even if
      // your project has ESLint errors.
      ignoreDuringBuilds: true,
      enable: false,
    }
  };
  
  export default nextConfig;
  