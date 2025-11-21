import AboutFounder from "../sections/AboutFounder";
import CTASection from "../sections/CTASection";
import FAQSection from "../sections/FAQSection";
import FeatureSection from "../sections/FeatureSection";
import HeroSection from "../sections/HeroSection";
import TrustStats from "../sections/TrustStats";

export default function LandingHomePage() {
  return (
    <div>
      <HeroSection />
      <FeatureSection />
      <TrustStats />
      <CTASection />
      {/* <AboutFounder /> */}
      <FAQSection />
    </div>
  );
}
