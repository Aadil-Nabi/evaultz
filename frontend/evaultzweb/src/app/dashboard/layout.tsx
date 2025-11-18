"use client";
import React from "react";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="h-full bg-white dark:bg-gray-900">
      <div className="h-full">{children}</div>
    </div>
    // <div className="w-screen h-screen flex flex-col bg-background">
    //   {children}
    // </div>
  );
}
