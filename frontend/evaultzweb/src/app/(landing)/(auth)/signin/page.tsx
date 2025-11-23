"use client";

import { signInUser } from "@/app/api/auth/signIn";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function SignInPage() {
  const router = useRouter();

  const [formData, setFormData] = useState({
    email: "",
    password: "",
    companyname: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.id]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const res = await signInUser(formData);
      if (res.status === 200) {
        router.refresh();
        router.push("/dashboard");
      }
    } catch (error) {
      console.error("signin failed", error);
      router.push("/signin");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50">
      <Card className="w-full max-w-sm shadow-lg rounded-2xl">
        <form onSubmit={handleSubmit}>
          <CardHeader>
            <CardTitle>Login to your account</CardTitle>
            <CardDescription>
              Enter your email below to login to your account
            </CardDescription>
            <CardAction>
              <Link href="/signup">
                <Button variant="link" type="button">
                  Sign Up
                </Button>
              </Link>
            </CardAction>
          </CardHeader>

          <CardContent>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="user@evaultz.cloud"
                  required
                  onChange={handleChange}
                  value={formData.email}
                />
              </div>

              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="/forgotpassword"
                    className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <Input
                  id="password"
                  type="password"
                  required
                  placeholder="••••••••"
                  onChange={handleChange}
                  value={formData.password}
                />
              </div>

              <div className="grid gap-2">
                <Label htmlFor="companyname">Company Name</Label>
                <Input
                  id="companyname"
                  type="text"
                  placeholder="your company name"
                  required
                  onChange={handleChange}
                  value={formData.companyname}
                />
              </div>
            </div>
          </CardContent>

          <CardFooter className="flex-col gap-2 pt-4">
            <Button type="submit" className="w-full">
              Login
            </Button>
            <Link href={"/"} className="w-full">
              <Button variant="outline" className="w-full" type="button">
                Cancel
              </Button>
            </Link>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
