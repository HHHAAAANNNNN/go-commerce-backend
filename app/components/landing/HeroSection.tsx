"use client";

export default function HeroSection() {
  return (
    <section id="hero" className="relative min-h-screen overflow-hidden bg-[#0A0A0F]">
      {/* Animated Background */}
      <div className="absolute inset-0 bg-gradient-to-br from-[#1a0f2e] via-[#0e1628] to-[#0A0A0F]">
        {/* Floating Blur Circles */}
        <div className="absolute top-10 -left-20 w-96 h-96 bg-primary-300 rounded-full mix-blend-screen filter blur-2xl opacity-40 animate-blob"></div>
        <div className="absolute top-20 -right-20 w-80 h-80 bg-secondary-300 rounded-full mix-blend-screen filter blur-2xl opacity-35 animate-blob animation-delay-2000"></div>
        <div className="absolute -bottom-20 left-40 w-96 h-96 bg-accent-300 rounded-full mix-blend-screen filter blur-2xl opacity-30 animate-blob animation-delay-4000"></div>
        <div className="absolute bottom-40 right-20 w-72 h-72 bg-primary-400 rounded-full mix-blend-screen filter blur-xl opacity-40 animate-blob animation-delay-6000"></div>
        
        {/* Mesh Overlay */}
        <div className="absolute inset-0 bg-grid-pattern opacity-20"></div>
        
        {/* Grain Texture */}
        <div className="absolute inset-0 opacity-20 mix-blend-soft-light">
          <div className="absolute inset-0 bg-noise"></div>
        </div>
        
        {/* Vignette Effect */}
        <div className="absolute inset-0 bg-gradient-radial from-transparent via-transparent to-black/60"></div>
      </div>

      {/* Content */}
      <div className="relative z-10 container mx-auto px-4 min-h-screen flex flex-col items-center justify-center text-center text-white pt-20">
        {/* Trust Badge */}
        <div className="mb-6 inline-flex items-center gap-2 px-4 py-2 bg-white/5 backdrop-blur-md rounded-full border border-primary-400/30 animate-fade-in-down">
          <span className="w-2 h-2 bg-accent-400 rounded-full animate-pulse shadow-lg shadow-accent-400/50"></span>
          <span className="text-sm font-medium text-slate-300">Trusted by 10,000+ Customers</span>
        </div>

        {/* Main Headline */}
        <h1 className="text-5xl md:text-7xl lg:text-8xl font-bold mb-6 max-w-5xl leading-tight animate-fade-in-up">
          <span className="text-slate-100">Discover Your</span>
          <span className="block mt-2 bg-gradient-to-r from-primary-400 via-accent-400 to-secondary-300 bg-clip-text text-transparent animate-shimmer bg-[length:200%_auto]">
            Perfect Product
          </span>
        </h1>

        {/* Subheadline */}
        <p className="text-xl md:text-2xl mb-12 max-w-2xl text-slate-400 leading-relaxed animate-fade-in-up animation-delay-200">
          Shop 1000+ premium products with exclusive member deals and free shipping
        </p>

        {/* CTA Buttons */}
        <div className="flex flex-col sm:flex-row gap-4 mb-16 animate-fade-in-up animation-delay-400">
          <a
            href="#products"
            className="group relative px-8 py-4 bg-gradient-to-r from-primary-400 to-primary-500 text-white rounded-full font-semibold text-lg overflow-hidden transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-primary-400/30"
          >
            <span className="relative z-10">Start Shopping</span>
            <span className="inline-block ml-2 group-hover:translate-x-1 transition-transform">→</span>
            <div className="absolute inset-0 bg-gradient-to-r from-primary-300 to-accent-300 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          </a>
          <a
            href="#products"
            className="px-8 py-4 bg-white/5 backdrop-blur-md text-slate-200 rounded-full font-semibold text-lg border-2 border-primary-400/30 hover:bg-white/10 hover:border-primary-400/60 transition-all duration-300 hover:scale-105"
          >
            View Deals
          </a>
        </div>

        {/* Trust Badges */}
        <div className="grid grid-cols-3 gap-8 max-w-2xl animate-fade-in-up animation-delay-600">
          <div className="flex flex-col items-center">
            <div className="w-12 h-12 mb-3 bg-primary-400/10 backdrop-blur-sm rounded-full flex items-center justify-center transition-all hover:scale-110 hover:bg-primary-400/20 duration-300 border border-primary-400/20">
              <svg className="w-6 h-6 text-primary-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z"/>
                <path d="M3 4a1 1 0 00-1 1v10a1 1 0 001 1h1.05a2.5 2.5 0 014.9 0H10a1 1 0 001-1V5a1 1 0 00-1-1H3zM14 7a1 1 0 00-1 1v6.05A2.5 2.5 0 0115.95 16H17a1 1 0 001-1v-5a1 1 0 00-.293-.707l-2-2A1 1 0 0015 7h-1z"/>
              </svg>
            </div>
            <p className="text-sm font-medium text-slate-300">Free Shipping</p>
          </div>
          <div className="flex flex-col items-center">
            <div className="w-12 h-12 mb-3 bg-accent-300/10 backdrop-blur-sm rounded-full flex items-center justify-center transition-all hover:scale-110 hover:bg-accent-300/20 duration-300 border border-accent-300/20">
              <svg className="w-6 h-6 text-accent-400" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"/>
              </svg>
            </div>
            <p className="text-sm font-medium text-slate-300">Secure Payment</p>
          </div>
          <div className="flex flex-col items-center">
            <div className="w-12 h-12 mb-3 bg-secondary-300/10 backdrop-blur-sm rounded-full flex items-center justify-center transition-all hover:scale-110 hover:bg-secondary-300/20 duration-300 border border-secondary-300/20">
              <svg className="w-6 h-6 text-secondary-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"/>
              </svg>
            </div>
            <p className="text-sm font-medium text-slate-300">24/7 Support</p>
          </div>
        </div>

        {/* Scroll Indicator */}
        <div className="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
          <svg className="w-6 h-6 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 14l-7 7m0 0l-7-7m7 7V3" />
          </svg>
        </div>
      </div>
    </section>
  );
}
