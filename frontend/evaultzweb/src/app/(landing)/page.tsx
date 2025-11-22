import { Testimonials } from "../sections/AnimatedTestimonials";
import CTASection from "../sections/CTASection";
import FAQSection from "../sections/FAQSection";
import FeatureSection from "../sections/FeatureSection";
import HeroSection from "../sections/HeroSection";
import PricingSection from "../sections/PricingSection";
import TrustStats from "../sections/TrustStats";

export default function LandingHomePage() {
  return (
    <div>
      <HeroSection />
      <FeatureSection />
      <TrustStats />
      <CTASection />
      <PricingSection />
      <Testimonials />
      <FAQSection />
    </div>
  );
}
