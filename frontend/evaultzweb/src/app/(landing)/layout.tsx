import type { Metadata } from "next";
import "../globals.css";
import LandingNavigationPage from "../sections/LandingNavigationPage";
import FooterSection from "../sections/FooterSection";

export const metadata: Metadata = {
  title: "Evaultz",
  description: "A File Uploader App",
};

export default function LandingLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html className="h-full bg-white dark:bg-gray-900">
      <body className="h-full">
        <LandingNavigationPage />
        {children}
        <FooterSection />
      </body>
    </html>
  );
}
