import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Evaultz",
  description: "A File Uploader App",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html className="h-full bg-white dark:bg-gray-900">
      <body className="h-full">{children}</body>
    </html>
  );
}
