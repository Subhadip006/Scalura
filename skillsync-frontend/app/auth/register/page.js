"use client";
import React, { useState } from 'react';
import { Eye, EyeOff, go } from 'lucide-react';
import Image from 'next/image';
import { motion } from 'framer-motion';
import { FcGoogle } from "react-icons/fc";
import { FaGithub } from "react-icons/fa";

export default function LoginPage() {
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Login attempt:', { email, password, rememberMe });
  };

  return (
    <div className="min-h-screen bg-gray-100 flex">

      {/* Right side */}
      <motion.div
        initial={{ opacity: 0, x: 100 }}
        animate={{ opacity: 1, x: 0 }}
        transition={{ duration: 0.6, ease: 'anticipate' }}
        className="flex-1 bg-white flex items-center justify-center px-8"
      >
        <div className="w-full max-w-lg">
          <div className="flex justify-end mb-2">
            <Image src="/SkillSync_logo.png" alt="Star Icon" width={100} height={100} className="mx-auto" />
          </div>

          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.3, duration: 0.6 }}
            className="text-center mb-8"
          >
            <h1 className="text-4xl font-bold text-gray-900 mb-2">Hey Welcome!</h1>
            <p className="text-gray-600">Please enter your details</p>
          </motion.div>

          {/* Form */}
          <form className="space-y-6" onSubmit={handleSubmit}>
            <div>
              <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full px-0 py-3 border-0 border-b border-gray-900 bg-transparent placeholder-gray-500 focus:border-gray-900 focus:outline-none focus:ring-0 text-gray-900"
                required
              />
            </div>

            <div className="relative">
              <input
                type={showPassword ? 'text' : 'password'}
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-0 py-3 pr-12 border-0 border-b border-gray-900 bg-transparent placeholder-gray-500 focus:border-gray-900 focus:outline-none focus:ring-0 text-gray-900"
                required
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-0 top-1/2 transform -translate-y-1/2 text-gray-500 hover:text-gray-700"
              >
                {showPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>

            <div className="flex items-center justify-between">
              <label className="flex items-center">
                <input
                  type="checkbox"
                  checked={rememberMe}
                  onChange={(e) => setRememberMe(e.target.checked)}
                  className="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span className="ml-2 text-sm text-gray-600">Remember for 30 days</span>
              </label>
              <a href="#" className="text-sm text-gray-600 hover:text-gray-900">
                Forgot password?
              </a>
            </div>

            <button
              type="submit"
              className="w-full bg-black text-white py-3 px-4 rounded-full font-medium hover:bg-gray-800 transition duration-200"
            >
              Sign Up
            </button>

            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.98 }}
              type="button"
              className="w-full gap-2 bg-gray-200 text-black py-3 px-4 rounded-full font-medium border border-gray-300 hover:bg-gray-100 transition duration-200 flex items-center justify-center"
            >
              <FcGoogle size={30}/>
              Sign Up with Google
            </motion.button>
            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.98 }}
              type="button"
              className="w-full gap-2 bg-gray-200 text-black py-3 px-4 rounded-full font-medium border border-gray-300 hover:bg-gray-100 transition duration-200 flex items-center justify-center"
            >
              <FaGithub size={30}/>
              Sign Up with Github
            </motion.button>
          </form>

          <div className="mt-8 text-center">
            <p className="text-gray-600">
              Already have an account?{' '}
              <a href="/auth/login" className="text-black font-medium hover:underline">
                Login
              </a>
            </p>
          </div>
        </div>
      </motion.div>

      <div className="flex-1 bg-gray-200 flex items-center justify-center relative overflow-hidden">
        <motion.div
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.8, ease: 'easeOut' }}
          className="relative flex flex-col items-center"
        >
          <Image src="/SkillSync.png" alt="SkillSync" width={400} height={400} />
        </motion.div>
      </div>
    </div>
  );
}
