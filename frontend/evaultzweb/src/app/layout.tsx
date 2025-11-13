import type { Metadata } from "next";
import "./globals.css";
import NavigationPage from "./ui/navigation";

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
      <body className="h-full">
        <NavigationPage />
        <main>{children}</main>
      </body>
    </html>
  );
}
