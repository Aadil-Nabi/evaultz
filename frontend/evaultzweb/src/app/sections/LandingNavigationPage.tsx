"use client";

import { useState, useEffect, useRef } from "react";
import {
  Dialog,
  DialogPanel,
  Disclosure,
  DisclosureButton,
  DisclosurePanel,
} from "@headlessui/react";
import {
  ArrowPathIcon,
  Bars3Icon,
  ChartPieIcon,
  CursorArrowRaysIcon,
  FingerPrintIcon,
  SquaresPlusIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import {
  ChevronDownIcon,
  PhoneIcon,
  PlayCircleIcon,
} from "@heroicons/react/20/solid";
import Link from "next/link";
import { DatabaseZap } from "lucide-react";

/* ---------------------------
   Menu data
---------------------------- */
const products = [
  {
    name: "Analytics",
    description: "Get a better understanding of your traffic",
    href: "/",
    icon: ChartPieIcon,
  },
  {
    name: "Security",
    description: "Your customers’ data will be safe and secure",
    href: "/",
    icon: FingerPrintIcon,
  },
  {
    name: "Integrations",
    description: "Connect with third-party tools",
    href: "/",
    icon: SquaresPlusIcon,
  },
];

const callsToAction = [
  { name: "Watch demo", href: "/", icon: PlayCircleIcon },
  { name: "Contact sales", href: "/", icon: PhoneIcon },
];

const about = [
  { name: "About us", href: "/about" },
  { name: "Support Project", href: "/" },
  { name: "Help", href: "/" },
];

export default function LandingNavigationPage() {
  // mobile menu (Dialog)
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  // header hide/show on scroll
  const [showHeader, setShowHeader] = useState(true);

  // Controlled dropdown states for desktop
  const [productOpen, setProductOpen] = useState(false);
  const [companyOpen, setCompanyOpen] = useState(false);

  // ref wrapping the whole nav area - used by outside click handler
  const navRef = useRef<HTMLElement | null>(null);

  // ------------- Scroll hide/show logic -------------
  useEffect(() => {
    let lastScroll = window.scrollY;
    const handleScroll = () => {
      const current = window.scrollY;
      if (current > lastScroll && current > 80) {
        // scrolling down -> hide header
        setShowHeader(false);
      } else {
        // scrolling up -> show header
        setShowHeader(true);
      }
      lastScroll = current;
    };
    window.addEventListener("scroll", handleScroll, { passive: true });
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // ------------- Outside click handler (closes all menus) -------------
  useEffect(() => {
    function handleClickOutside(e: MouseEvent) {
      const target = e.target as Node;
      if (navRef.current && !navRef.current.contains(target)) {
        // If click is outside the nav wrapper, close everything:
        setProductOpen(false);
        setCompanyOpen(false);
        setMobileMenuOpen(false);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // Helper to toggle product dropdown but ensure company closes
  const toggleProduct = () => {
    setProductOpen((v) => {
      const next = !v;
      if (next) setCompanyOpen(false);
      return next;
    });
  };

  // Helper to toggle company dropdown but ensure product closes
  const toggleCompany = () => {
    setCompanyOpen((v) => {
      const next = !v;
      if (next) setProductOpen(false);
      return next;
    });
  };

  return (
    <header
      className={` dark:bg-gray-900 fixed top-0 left-0 right-0 z-50 transition-transform duration-300 ${
        showHeader ? "translate-y-0" : "-translate-y-full"
      }`}
    >
      {/* Wrap nav with ref so outside clicks can be detected */}
      <nav
        aria-label="Global"
        ref={navRef}
        className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8"
      >
        <div className="flex lg:flex-1">
          <Link href="/" className="-m-1.5 p-1.5">
            {/* <span className="sr-only">eVaultz</span> */}
            <div className="flex shrink-0 items-center">
              <DatabaseZap className=" w-auto dark:hidden" />
              <h1 className="font-bold  pl-1">eVaultz Inc.</h1>
            </div>
          </Link>
        </div>
        {/* Mobile menu button */}
        <div className="flex lg:hidden">
          <button
            type="button"
            onClick={() => setMobileMenuOpen(true)}
            className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700 dark:text-gray-400"
          >
            <span className="sr-only">Open main menu</span>
            <Bars3Icon aria-hidden="true" className="size-6" />
          </button>
        </div>
        {/* Desktop menu */}
        <div className="hidden lg:flex lg:gap-x-12 items-center">
          {/* ---------------- Product dropdown (controlled) ---------------- */}
          <div className="relative">
            <button
              type="button"
              aria-expanded={productOpen}
              onClick={toggleProduct}
              className="flex items-center gap-x-1 text-sm font-semibold text-gray-900 dark:text-white"
            >
              Product
              <ChevronDownIcon
                aria-hidden="true"
                className="size-5 flex-none text-gray-400 dark:text-gray-500"
              />
            </button>

            {/* Render panel when productOpen is true */}
            {productOpen && (
              <div className="absolute left-1/2 z-10 mt-3 w-screen max-w-md -translate-x-1/2 overflow-hidden rounded-3xl bg-white shadow-lg outline-1 outline-gray-900/5 transition dark:bg-gray-800 dark:shadow-none dark:outline-white/10">
                <div className="p-4">
                  {products.map((product) => (
                    <div
                      key={product.name}
                      className="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm hover:bg-gray-50 dark:hover:bg-white/5"
                    >
                      <div className="flex size-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white dark:bg-gray-700/50 dark:group-hover:bg-gray-700">
                        <product.icon
                          aria-hidden="true"
                          className="size-6 text-gray-600 group-hover:text-indigo-600 dark:text-gray-400 dark:group-hover:text-white"
                        />
                      </div>
                      <div className="flex-auto">
                        <Link
                          href={product.href}
                          onClick={() => {
                            // close menus when a link is clicked
                            setProductOpen(false);
                          }}
                          className="block font-semibold text-gray-900 dark:text-white"
                        >
                          {product.name}
                        </Link>
                        <p className="mt-1 text-gray-600 dark:text-gray-400">
                          {product.description}
                        </p>
                      </div>
                    </div>
                  ))}
                </div>

                <div className="grid grid-cols-2 divide-x divide-gray-900/5 bg-gray-50 dark:divide-white/10 dark:bg-gray-700/50">
                  {callsToAction.map((item) => (
                    <Link
                      key={item.name}
                      href={item.href}
                      onClick={() => setProductOpen(false)}
                      className="flex items-center justify-center gap-x-2.5 p-3 text-sm font-semibold text-gray-900 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700/50"
                    >
                      <item.icon aria-hidden="true" className="size-5" />
                      {item.name}
                    </Link>
                  ))}
                </div>
              </div>
            )}
          </div>

          {/* Regular links */}
          <Link
            href="/"
            className="text-sm font-semibold text-gray-900 dark:text-white"
          >
            Solution
          </Link>
          <Link
            href="/"
            className="text-sm font-semibold text-gray-900 dark:text-white"
          >
            Feature
          </Link>
          <Link
            href="/"
            className="text-sm font-semibold text-gray-900 dark:text-white"
          >
            Pricing
          </Link>

          {/* ---------------- Company dropdown (controlled) ---------------- */}
          <div className="relative">
            <button
              type="button"
              aria-expanded={companyOpen}
              onClick={toggleCompany}
              className="flex items-center gap-x-1 text-sm font-semibold text-gray-900 dark:text-white"
            >
              About eVaultz
              <ChevronDownIcon
                aria-hidden="true"
                className="size-5 flex-none text-gray-400 dark:text-gray-500"
              />
            </button>

            {companyOpen && (
              <div className="absolute left-1/2 z-10 mt-3 w-56 -translate-x-1/2 rounded-xl bg-white p-2 shadow-lg outline-1 outline-gray-900/5 transition dark:bg-gray-800 dark:shadow-none dark:outline-white/10">
                {about.map((about) => (
                  <Link
                    key={about.name}
                    href={about.href}
                    onClick={() => setCompanyOpen(false)}
                    className="block rounded-lg px-3 py-2 text-sm font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                  >
                    {about.name}
                  </Link>
                ))}
              </div>
            )}
          </div>
        </div>
        {/* Right-side login */}
        <div className="hidden lg:flex lg:flex-1 lg:justify-end">
          <Link
            href="/signin"
            className="text-sm font-semibold text-gray-900 dark:text-white"
          >
            Sign in <span aria-hidden="true">→</span>
          </Link>
        </div>
      </nav>

      {/* ---------------- Mobile Menu (Dialog) ---------------- */}
      <Dialog
        open={mobileMenuOpen}
        onClose={setMobileMenuOpen}
        className="lg:hidden"
      >
        <div className="fixed inset-0 z-50" />
        <DialogPanel className="fixed inset-y-0 right-0 z-50 w-full overflow-y-auto bg-white p-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10 dark:bg-gray-900 dark:sm:ring-gray-100/10">
          <div className="flex items-center justify-between">
            <Link href="/" className="-m-1.5 p-1.5">
              <span className="sr-only">eVaultz</span>
              <div className="flex h-16 shrink-0 items-center">
                <DatabaseZap className="h-8 w-auto dark:hidden" />
                <h1 className="font-bold h-8 pl-1">eVaultz Inc.</h1>
              </div>
            </Link>
            <button
              type="button"
              onClick={() => setMobileMenuOpen(false)}
              className="-m-2.5 rounded-md p-2.5 text-gray-700 dark:text-gray-400"
            >
              <span className="sr-only">Close menu</span>
              <XMarkIcon aria-hidden="true" className="size-6" />
            </button>
          </div>

          <div className="mt-6 flow-root">
            <div className="-my-6 divide-y divide-gray-500/10 dark:divide-white/10">
              <div className="space-y-2 py-6">
                <Disclosure as="div" className="-mx-3">
                  <DisclosureButton className="group flex w-full items-center justify-between rounded-lg py-2 pr-3.5 pl-3 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5">
                    Product
                    <ChevronDownIcon
                      aria-hidden="true"
                      className="size-5 flex-none group-data-open:rotate-180"
                    />
                  </DisclosureButton>
                  <DisclosurePanel className="mt-2 space-y-2">
                    {[...products, ...callsToAction].map((item) => (
                      <DisclosureButton
                        key={item.name}
                        as="a"
                        href={item.href}
                        onClick={() => setMobileMenuOpen(false)}
                        className="block rounded-lg py-2 pr-3 pl-6 text-sm/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                      >
                        {item.name}
                      </DisclosureButton>
                    ))}
                  </DisclosurePanel>
                </Disclosure>

                <Link
                  href="/"
                  className="-mx-3 block rounded-lg px-3 py-2 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                >
                  Solution
                </Link>
                <Link
                  href="/"
                  className="-mx-3 block rounded-lg px-3 py-2 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                >
                  Feature
                </Link>

                <Link
                  href="/"
                  className="-mx-3 block rounded-lg px-3 py-2 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                >
                  Pricing
                </Link>

                <Disclosure as="div" className="-mx-3">
                  <DisclosureButton className="group flex w-full items-center justify-between rounded-lg py-2 pr-3.5 pl-3 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5">
                    About eVaultz
                    <ChevronDownIcon
                      aria-hidden="true"
                      className="size-5 flex-none group-data-open:rotate-180"
                    />
                  </DisclosureButton>
                  <DisclosurePanel className="mt-2 space-y-2">
                    {about.map((item) => (
                      <DisclosureButton
                        key={item.name}
                        as="a"
                        href={item.href}
                        onClick={() => setMobileMenuOpen(false)}
                        className="block rounded-lg py-2 pr-3 pl-6 text-sm/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                      >
                        {item.name}
                      </DisclosureButton>
                    ))}
                  </DisclosurePanel>
                </Disclosure>
              </div>

              <div className="py-6">
                <Link
                  href="/signin"
                  className="-mx-3 block rounded-lg px-3 py-2.5 text-base/7 font-semibold text-gray-900 hover:bg-gray-50 dark:text-white dark:hover:bg-white/5"
                >
                  Sign in
                </Link>
              </div>
            </div>
          </div>
        </DialogPanel>
      </Dialog>
    </header>
  );
}
