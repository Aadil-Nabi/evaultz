import type { Metadata } from "next";
import "../globals.css";
import FooterSection from "../sections/FooterSection";
import HeaderSectionPage from "../sections/HeaderSection";

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
    <div className="h-full bg-white dark:bg-gray-900">
      {/* <LandingNavigationPage /> */}
      <HeaderSectionPage />
      <div className="h-full">
        {children}

        <FooterSection />
      </div>
    </div>
  );
}
