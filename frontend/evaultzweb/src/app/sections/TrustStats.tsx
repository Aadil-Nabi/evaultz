const stats = [
  { id: 1, name: "Files Secured & Encrypted", value: "2M+" },
  { id: 2, name: "Zero Data Breaches", value: "100%" },
  { id: 3, name: "Platform Availability", value: "99.99%" },
  { id: 4, name: "Encryption Performance", value: "10ms Avg" },
];

export default function TrustStats() {
  return (
    <div className="bg-white py-24 sm:py-32 dark:bg-gray-900">
      <div className="mx-auto max-w-7xl px-6 lg:px-8">
        <div className="mx-auto max-w-2xl lg:max-w-none">
          <div className="text-center">
            <h2 className="text-4xl font-semibold tracking-tight text-balance text-gray-900 sm:text-5xl dark:text-white">
              Built for Trust. Secured for You.
            </h2>
            <p className="mt-4 text-lg/8 text-gray-600 dark:text-gray-300">
              eVaultz ensures your data stays protected with enterprise-grade
              encryption, continuous integrity checks, and zero-trust
              architecture.
            </p>
          </div>

          <dl className="mt-16 grid grid-cols-1 gap-0.5 overflow-hidden rounded-2xl text-center sm:grid-cols-2 lg:grid-cols-4">
            {stats.map((stat) => (
              <div
                key={stat.id}
                className="flex flex-col bg-gray-400/5 p-8 dark:bg-white/5"
              >
                <dt className="text-sm/6 font-semibold text-gray-600 dark:text-gray-300">
                  {stat.name}
                </dt>
                <dd className="order-first text-3xl font-semibold tracking-tight text-gray-900 dark:text-white">
                  {stat.value}
                </dd>
              </div>
            ))}
          </dl>
        </div>
      </div>
    </div>
  );
}
