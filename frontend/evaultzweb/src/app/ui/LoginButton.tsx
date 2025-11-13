"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
export default function LoginButton() {
  const router = useRouter();
  const handleClick = () => {
    router.push("/login");
  };
  return (
    <div>
      <Button
        type="submit"
        onClick={handleClick}
        className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 dark:bg-indigo-500 dark:shadow-none dark:hover:bg-indigo-400 dark:focus-visible:outline-indigo-500"
      >
        Login
      </Button>
    </div>
  );
}
