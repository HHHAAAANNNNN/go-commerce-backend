import Navbar from "./components/Navbar";
import HeroSection from "./components/landing/HeroSection";
import FeaturedCategories from "./components/landing/FeaturedCategories";
import BestSellers from "./components/landing/BestSellers";
import WhyUs from "./components/landing/WhyUs";
import SocialProof from "./components/landing/SocialProof";
import CTASection from "./components/landing/CTASection";

export default function Home() {
  return (
    <>
      <Navbar />
      <main>
        <HeroSection />
        <FeaturedCategories />
        <BestSellers />
        <WhyUs />
        <SocialProof />
        <CTASection />
      </main>
    </>
  );
}
