import React from "react";

export default function DashboardLayout({
  children,
  files,
  uploadfiles,
  notifications,
}: {
  children: React.ReactNode;
  files: React.ReactNode;
  uploadfiles: React.ReactNode;
  notifications: React.ReactNode;
}) {
  return (
    <>
      <div className="container-flex bg-gray-300 ">
        <div className="flex min-h-full flex-col">
          {/* 3 column wrapper */}
          <div className="mx-auto w-full max-w-7xl grow lg:flex xl:px-2 bg-amber-200">
            {children}
            <div className="flex-1 xl:flex">
              <div className="border-b bg-red-300 border-gray-200 px-4 py-6 sm:px-6 lg:pl-8 xl:w-64 xl:shrink-0 xl:border-r xl:border-b-0 xl:pl-6 dark:border-white/10">
                {uploadfiles}
              </div>

              <div className="px-4 py-6 sm:px-6 lg:pl-8 xl:flex-1 xl:pl-6 bg-red-400">
                {files}
              </div>
            </div>
            <div className="shrink-0 border-t bg-red-500 border-gray-200 px-4 py-6 sm:px-6 lg:w-96 lg:border-t-0 lg:border-l lg:pr-8 xl:pr-6 dark:border-white/10">
              {notifications}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
