import AboutFounder from "../sections/AboutFounder";
import CTASection from "../sections/CTASection";
import FAQSection from "../sections/FAQSection";
import FeatureSection from "../sections/FeatureSection";
import HeroSection from "../sections/HeroSection";

export default function LandingHomePage() {
  return (
    <div>
      <HeroSection />
      <FeatureSection />
      <CTASection />
      <AboutFounder />
      <FAQSection />
    </div>
  );
}
