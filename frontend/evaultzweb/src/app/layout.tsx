import type { Metadata } from "next";
import "./globals.css";
import LandingNavigationPage from "./sections/LandingNavigationPage";
import FooterSection from "./sections/FooterSection";

export const metadata: Metadata = {
  title: "Evaultz",
  description: "A File Uploader App",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    // <html lang="en" className="h-full">
    //   <body className="h-full bg-white dark:bg-gray-900">{children}</body>
    // </html>

    <html lang="en" className="h-full" suppressHydrationWarning>
      <body className="h-full">
        <div className="min-h-screen bg-white dark:bg-gray-900">{children}</div>
      </body>
    </html>
  );
}
