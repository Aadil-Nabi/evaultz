import {
  CloudArrowUpIcon,
  LockClosedIcon,
  ServerIcon,
} from "@heroicons/react/20/solid";
import Image from "next/image";

export default function Example() {
  return (
    <div className="relative bg-white dark:bg-gray-900">
      <div className="mx-auto max-w-7xl lg:flex lg:justify-between lg:px-8">
        {/* LEFT: Static Image */}
        <div className="hidden lg:block lg:w-1/2 sticky top-0 h-screen">
          <div className="relative w-full h-full">
            <Image
              alt="founder"
              src="/about-founder.jpeg"
              fill
              className="object-cover rounded-2xl bg-gray-50 dark:bg-gray-800"
            />
          </div>
        </div>

        {/* RIGHT: Scrollable Content */}
        <div className="w-full lg:w-1/2 px-6 overflow-y-auto">
          <div className="mx-auto max-w-2xl pt-16 pb-24 sm:pt-20 sm:pb-32 lg:max-w-lg lg:pt-32">
            <p className="text-base/7 font-semibold text-indigo-600 dark:text-indigo-400">
              Deploy faster
            </p>

            <h1 className="mt-2 text-4xl font-semibold tracking-tight text-gray-900 dark:text-white">
              A better workflow
            </h1>

            <p className="mt-6 text-xl/8 text-gray-700 dark:text-gray-300">
              Aliquet nec orci mattis amet quisque ullamcorper neque...
            </p>

            <div className="mt-10 text-base/7 text-gray-600 dark:text-gray-400">
              <p>
                Faucibus commodo massa rhoncus, volutpat. Dignissim sed eget
                risus enim...
              </p>

              <ul role="list" className="mt-8 space-y-8">
                <li className="flex gap-x-3">
                  <CloudArrowUpIcon className="size-5 mt-1 text-indigo-600 dark:text-indigo-400" />
                  <span>
                    <strong className="font-semibold text-gray-900 dark:text-white">
                      Push to deploy.
                    </strong>{" "}
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  </span>
                </li>

                <li className="flex gap-x-3">
                  <LockClosedIcon className="size-5 mt-1 text-indigo-600 dark:text-indigo-400" />
                  <span>
                    <strong className="font-semibold text-gray-900 dark:text-white">
                      SSL certificates.
                    </strong>{" "}
                    Anim aute id magna aliqua ad non deserunt sunt.
                  </span>
                </li>

                <li className="flex gap-x-3">
                  <ServerIcon className="size-5 mt-1 text-indigo-600 dark:text-indigo-400" />
                  <span>
                    <strong className="font-semibold text-gray-900 dark:text-white">
                      Database backups.
                    </strong>{" "}
                    Et magna sit morbi lobortis.
                  </span>
                </li>
              </ul>

              <p className="mt-8">
                Et vitae blandit facilisi magna lacus commodo. Vitae sapien duis
                odio id et...
              </p>

              <h2 className="mt-16 text-2xl font-bold text-gray-900 dark:text-white">
                No server? No problem.
              </h2>

              <p className="mt-6">
                Id orci tellus laoreet id ac. Dolor, aenean leo, ac etiam
                consequat in...
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
