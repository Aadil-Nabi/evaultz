import {
  ArrowPathIcon,
  CloudArrowUpIcon,
  Cog6ToothIcon,
  FingerPrintIcon,
  LockClosedIcon,
  ServerIcon,
} from "@heroicons/react/20/solid";

const features = [
  {
    name: "Secure File Storage.",
    description:
      "A hardened encrypted repository built for confidentiality and resilience",
    icon: CloudArrowUpIcon,
  },
  {
    name: "Analytics & Monitoring.",
    description:
      "Full audit visibility to track every file, event, and security anomaly.",
    icon: LockClosedIcon,
  },
  {
    name: "Data Classification & Integrity Checks.",
    description:
      "Intelligent classification and tamper detection for regulated workloads",
    icon: ArrowPathIcon,
  },
  {
    name: "Advanced security.",
    description:
      "Zero-trust policies, MFA, and compliance-ready security frameworks",
    icon: FingerPrintIcon,
  },
  {
    name: "Encryption on the fly",
    description:
      "Transparent real-time encryption powered by industry-standard cryptography.",
    icon: Cog6ToothIcon,
  },
  {
    name: "Mutitenant file share.",
    description:
      "Isolated tenant architecture for safe collaboration across organizations",
    icon: ServerIcon,
  },
];

export default function FeatureSection() {
  return (
    <div className="bg-white py-24 sm:py-32 dark:bg-gray-900">
      <div className="mx-auto max-w-7xl px-6 lg:px-8">
        <div className="mx-auto max-w-2xl sm:text-center">
          <h2 className="text-base/7 font-semibold text-indigo-600 dark:text-indigo-400">
            Everything you need
          </h2>
          <p className="mt-2 text-4xl font-semibold tracking-tight text-pretty text-gray-900 sm:text-5xl sm:text-balance dark:text-white">
            Trustworthy Security. Smarter Data Control.
          </p>
          <p className="mt-6 text-lg/8 text-gray-600 dark:text-gray-300">
            eVaultz empowers teams to protect their critical files with
            transparent, reliable, and verifiable security.
          </p>
        </div>
      </div>
      <div className="relative overflow-hidden pt-16">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <img
            alt="App screenshot"
            src="/eVaultz-dash.png"
            width={2432}
            height={1442}
            className="mb-[-12%] rounded-xl shadow-2xl ring-1 ring-gray-900/10 dark:hidden dark:ring-white/10"
          />
          <img
            alt="App screenshot"
            src="/eVaultz-dash.png"
            width={2432}
            height={1442}
            className="mb-[-12%] rounded-xl shadow-2xl ring-1 ring-gray-900/10 not-dark:hidden dark:ring-white/10"
          />
          <div aria-hidden="true" className="relative">
            <div className="absolute -inset-x-20 bottom-0 bg-linear-to-t from-white pt-[7%] dark:from-gray-900" />
          </div>
        </div>
      </div>
      <div className="mx-auto mt-16 max-w-7xl px-6 sm:mt-20 md:mt-24 lg:px-8">
        <dl className="mx-auto grid max-w-2xl grid-cols-1 gap-x-6 gap-y-10 text-base/7 text-gray-600 sm:grid-cols-2 lg:mx-0 lg:max-w-none lg:grid-cols-3 lg:gap-x-8 lg:gap-y-16 dark:text-gray-400">
          {features.map((feature) => (
            <div key={feature.name} className="relative pl-9">
              <dt className="inline font-semibold text-gray-900 dark:text-white">
                <feature.icon
                  aria-hidden="true"
                  className="absolute top-1 left-1 size-5 text-indigo-600 dark:text-indigo-400"
                />
                {feature.name}
              </dt>{" "}
              <dd className="inline">{feature.description}</dd>
            </div>
          ))}
        </dl>
      </div>
    </div>
  );
}
