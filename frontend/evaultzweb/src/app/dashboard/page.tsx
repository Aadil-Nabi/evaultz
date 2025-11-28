"use client";

import { useState } from "react";
import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
} from "@headlessui/react";
import {
  Bars3Icon,
  BellIcon,
  ChartPieIcon,
  Cog6ToothIcon,
  DocumentDuplicateIcon,
  FolderIcon,
  HomeIcon,
  UsersIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import {
  ChevronDownIcon,
  MagnifyingGlassIcon,
} from "@heroicons/react/20/solid";
import UploadsChart from "@/components/charts/UploadsChart";
import PublicVsPrivateChart from "@/components/charts/PublicVsPrivateChart";
import { ChartContainer } from "@/components/charts/ChartContainer";
import StorageUsageChart from "@/components/charts/StorageUsageChart";
import SharedFilesPage from "../files/@shared/page";

import { DatabaseZap } from "lucide-react";
import TenantUsageChart from "@/components/charts/TenantUsageChart";
import SharingTrendChart from "@/components/charts/SharingTrendChart";
import FileTypeChart from "@/components/charts/FileTypeChart";
import Link from "next/link";
import DropdownMenuPage from "@/components/DropDownMenu";
import EncryptedVsPlainFilesChart from "@/components/charts/EncryptedVsPlainFilesChart";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import axiosClient from "../api/axiosClient";
import { SignOutUser } from "../api/auth/signOut";

const navigation = [
  { name: "Dashboard", href: "/dashboard", icon: HomeIcon, current: true },
  { name: "Manage Files", href: "/files", icon: FolderIcon, current: false },
  { name: "Projects", href: "/dashboard", icon: FolderIcon, current: false },
];
const favourites = [
  { id: 1, name: "Design", href: "/dashboard", initial: "D", current: false },
  {
    id: 2,
    name: "Integrations",
    href: "/dashboard",
    initial: "I",
    current: false,
  },
  {
    id: 3,
    name: "Project Management",
    href: "/dashboard",
    initial: "P",
    current: false,
  },
];
const userNavigation = [
  { name: "Your profile", href: "/dashboard" },
  { name: "Sign out", action: "logout" },
];

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

export default function Page() {
  const router = useRouter();

  async function handleLogout() {
    await SignOutUser();
    router.push("/signin");
  }

  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <>
      <div>
        <Dialog
          open={sidebarOpen}
          onClose={setSidebarOpen}
          className="relative z-50 lg:hidden"
        >
          <DialogBackdrop
            transition
            className="fixed inset-0 bg-gray-900/80 transition-opacity duration-300 ease-linear data-closed:opacity-0"
          />

          <div className="fixed inset-0 flex">
            <DialogPanel
              transition
              className="relative mr-16 flex w-full max-w-xs flex-1 transform transition duration-300 ease-in-out data-closed:-translate-x-full"
            >
              <TransitionChild>
                <div className="absolute top-0 left-full flex w-16 justify-center pt-5 duration-300 ease-in-out data-closed:opacity-0">
                  <button
                    type="button"
                    onClick={() => setSidebarOpen(false)}
                    className="-m-2.5 p-2.5"
                  >
                    <span className="sr-only">Close sidebar</span>
                    <XMarkIcon
                      aria-hidden="true"
                      className="size-6 text-white"
                    />
                  </button>
                </div>
              </TransitionChild>

              {/* Sidebar component, swap this element with another sidebar if you like */}
              <div className="relative flex grow flex-col gap-y-5 overflow-y-auto bg-white px-6 pb-4 dark:bg-gray-900 dark:ring dark:ring-white/10 dark:before:pointer-events-none dark:before:absolute dark:before:inset-0 dark:before:bg-black/10">
                <div className="relative flex h-16 shrink-0 items-center">
                  <DatabaseZap />
                </div>
                <nav className="relative flex flex-1 flex-col">
                  <ul role="list" className="flex flex-1 flex-col gap-y-7">
                    <li>
                      <ul role="list" className="-mx-2 space-y-1">
                        {navigation.map((item) => (
                          <li key={item.name}>
                            <Link
                              href={item.href}
                              className={classNames(
                                item.current
                                  ? "bg-gray-50 text-indigo-600 dark:bg-white/5 dark:text-white"
                                  : "text-gray-700 hover:bg-gray-50 hover:text-indigo-600 dark:text-gray-400 dark:hover:bg-white/5 dark:hover:text-white",
                                "group flex gap-x-3 rounded-md p-2 text-sm/6 font-semibold"
                              )}
                            >
                              <item.icon
                                aria-hidden="true"
                                className={classNames(
                                  item.current
                                    ? "text-indigo-600 dark:text-white"
                                    : "text-gray-400 group-hover:text-indigo-600 dark:group-hover:text-white",
                                  "size-6 shrink-0"
                                )}
                              />
                              {item.name}
                            </Link>
                          </li>
                        ))}
                      </ul>
                    </li>
                    <li>
                      <div className="text-xs/6 font-semibold text-gray-400">
                        Your teams
                      </div>
                      <ul role="list" className="-mx-2 mt-2 space-y-1">
                        {favourites.map((favourite) => (
                          <li key={favourite.name}>
                            <Link
                              href={favourite.href}
                              className={classNames(
                                favourite.current
                                  ? "bg-gray-50 text-indigo-600 dark:bg-white/5 dark:text-white"
                                  : "text-gray-700 hover:bg-gray-50 hover:text-indigo-600 dark:text-gray-400 dark:hover:bg-white/5 dark:hover:text-white",
                                "group flex gap-x-3 rounded-md p-2 text-sm/6 font-semibold"
                              )}
                            >
                              <span
                                className={classNames(
                                  favourite.current
                                    ? "border-indigo-600 text-indigo-600 dark:border-white/20 dark:text-white"
                                    : "border-gray-200 text-gray-400 group-hover:border-indigo-600 group-hover:text-indigo-600 dark:border-white/10 dark:group-hover:border-white/20 dark:group-hover:text-white",
                                  "flex size-6 shrink-0 items-center justify-center rounded-lg border bg-white text-[0.625rem] font-medium dark:bg-white/5"
                                )}
                              >
                                {favourite.initial}
                              </span>
                              <span className="truncate">{favourite.name}</span>
                            </Link>
                          </li>
                        ))}
                      </ul>
                    </li>

                    <li className="mt-auto w-fit">
                      <DropdownMenuPage />
                    </li>
                  </ul>
                </nav>
              </div>
            </DialogPanel>
          </div>
        </Dialog>

        {/* Static sidebar for desktop */}
        <div className="hidden bg-gray-900 lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
          {/* Sidebar component, swap this element with another sidebar if you like */}
          <div className="flex grow flex-col gap-y-5 overflow-y-auto border-r border-gray-200 bg-white px-6 pb-4 dark:border-white/10 dark:bg-black/10">
            <Link href="/dashboard" className="flex h-16 shrink-0 items-center">
              <DatabaseZap className="h-8 w-auto dark:hidden" />
              <h1 className="font-bold h-8 pl-1">eVaultz Inc.</h1>
            </Link>
            <nav className="flex flex-1 flex-col">
              <ul role="list" className="flex flex-1 flex-col gap-y-7">
                <li>
                  <ul role="list" className="-mx-2 space-y-1">
                    {navigation.map((item) => (
                      <li key={item.name}>
                        <Link
                          href={item.href}
                          className={classNames(
                            item.current
                              ? "bg-gray-50 text-indigo-600 dark:bg-white/5 dark:text-white"
                              : "text-gray-700 hover:bg-gray-50 hover:text-indigo-600 dark:text-gray-400 dark:hover:bg-white/5 dark:hover:text-white",
                            "group flex gap-x-3 rounded-md p-2 text-sm/6 font-semibold"
                          )}
                        >
                          <item.icon
                            aria-hidden="true"
                            className={classNames(
                              item.current
                                ? "text-indigo-600 dark:text-white"
                                : "text-gray-400 group-hover:text-indigo-600 dark:group-hover:text-white",
                              "size-6 shrink-0"
                            )}
                          />
                          {item.name}
                        </Link>
                      </li>
                    ))}
                  </ul>
                </li>
                <li>
                  <div className="text-xs/6 font-semibold text-gray-400">
                    Your teams
                  </div>
                  <ul role="list" className="-mx-2 mt-2 space-y-1">
                    {favourites.map((favourite) => (
                      <li key={favourite.name}>
                        <Link
                          href={favourite.href}
                          className={classNames(
                            favourite.current
                              ? "bg-gray-50 text-indigo-600 dark:bg-white/5 dark:text-white"
                              : "text-gray-700 hover:bg-gray-50 hover:text-indigo-600 dark:text-gray-400 dark:hover:bg-white/5 dark:hover:text-white",
                            "group flex gap-x-3 rounded-md p-2 text-sm/6 font-semibold"
                          )}
                        >
                          <span
                            className={classNames(
                              favourite.current
                                ? "border-indigo-600 text-indigo-600 dark:border-white/20 dark:text-white"
                                : "border-gray-200 text-gray-400 group-hover:border-indigo-600 group-hover:text-indigo-600 dark:border-white/10 dark:group-hover:border-white/20 dark:group-hover:text-white",
                              "flex size-6 shrink-0 items-center justify-center rounded-lg border bg-white text-[0.625rem] font-medium dark:bg-white/5"
                            )}
                          >
                            {favourite.initial}
                          </span>
                          <span className="truncate">{favourite.name}</span>
                        </Link>
                      </li>
                    ))}
                  </ul>
                </li>

                <li className="mt-auto w-fit">
                  <DropdownMenuPage />
                </li>
              </ul>
            </nav>
          </div>
        </div>

        <div className="lg:pl-72">
          <div className="sticky top-0 z-40 flex h-16 shrink-0 items-center gap-x-4 border-b border-gray-200 bg-white px-4 shadow-xs sm:gap-x-6 sm:px-6 lg:px-8 dark:border-white/10 dark:bg-gray-900 dark:shadow-none">
            <button
              type="button"
              onClick={() => setSidebarOpen(true)}
              className="-m-2.5 p-2.5 text-gray-700 hover:text-gray-900 lg:hidden dark:text-gray-400 dark:hover:text-white"
            >
              <span className="sr-only">Open sidebar</span>
              <Bars3Icon aria-hidden="true" className="size-6" />
            </button>

            {/* Separator */}
            <div
              aria-hidden="true"
              className="h-6 w-px bg-gray-200 lg:hidden dark:bg-white/10"
            />

            <div className="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
              <form action="#" method="GET" className="grid flex-1 grid-cols-1">
                <input
                  name="search"
                  placeholder="Search"
                  aria-label="Search"
                  className="col-start-1 row-start-1 block size-full bg-white pl-8 text-base text-gray-900 outline-hidden placeholder:text-gray-400 sm:text-sm/6 dark:bg-gray-900 dark:text-white dark:placeholder:text-gray-500"
                />
                <MagnifyingGlassIcon
                  aria-hidden="true"
                  className="pointer-events-none col-start-1 row-start-1 size-5 self-center text-gray-400"
                />
              </form>
              <div className="flex items-center gap-x-4 lg:gap-x-6">
                <button
                  type="button"
                  className="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500 dark:hover:text-white"
                >
                  <span className="sr-only">View notifications</span>
                  <BellIcon aria-hidden="true" className="size-6" />
                </button>

                {/* Separator */}
                <div
                  aria-hidden="true"
                  className="hidden lg:block lg:h-6 lg:w-px lg:bg-gray-200 dark:lg:bg-white/10"
                />

                {/* Profile dropdown */}
                <Menu as="div" className="relative">
                  <MenuButton className="relative flex items-center">
                    <span className="absolute -inset-1.5" />
                    <span className="sr-only">Open user menu</span>
                    <img
                      alt=""
                      src="/globe.svg"
                      className="size-8 rounded-full bg-gray-50 outline -outline-offset-1 outline-black/5 dark:bg-gray-800 dark:outline-white/10"
                    />
                    <span className="hidden lg:flex lg:items-center">
                      <span
                        aria-hidden="true"
                        className="ml-4 text-sm/6 font-semibold text-gray-900 dark:text-white"
                      >
                        Aadil Nabi
                      </span>
                      <ChevronDownIcon
                        aria-hidden="true"
                        className="ml-2 size-5 text-gray-400 dark:text-gray-500"
                      />
                    </span>
                  </MenuButton>
                  {/* <MenuItems
                    transition
                    className="absolute right-0 z-10 mt-2.5 w-32 origin-top-right rounded-md bg-white py-2 shadow-lg outline-1 outline-gray-900/5 transition data-closed:scale-95 data-closed:transform data-closed:opacity-0 data-enter:duration-100 data-enter:ease-out data-leave:duration-75 data-leave:ease-in dark:bg-gray-800 dark:shadow-none dark:-outline-offset-1 dark:outline-white/10"
                  >
                    {userNavigation.map((item) => (
                      <MenuItem key={item.name}>
                        <Link
                          href={item.href}
                          className="block px-3 py-1 text-sm/6 text-gray-900 data-focus:bg-gray-50 data-focus:outline-hidden dark:text-white dark:data-focus:bg-white/5"
                        >
                          {item.name}
                        </Link>
                      </MenuItem>
                    ))}
                  </MenuItems> */}

                  <MenuItems
                    transition
                    className="absolute right-0 z-10 mt-2.5 w-32 origin-top-right rounded-md bg-white py-2 shadow-lg"
                  >
                    {userNavigation.map((item) => (
                      <MenuItem key={item.name}>
                        {item.href ? (
                          <Link
                            href={item.href}
                            className="block px-3 py-1 text-sm text-gray-900"
                          >
                            {item.name}
                          </Link>
                        ) : (
                          <button
                            onClick={handleLogout}
                            className="block w-full text-left px-3 py-1 text-sm text-gray-900 hover:bg-gray-50"
                          >
                            {item.name}
                          </button>
                        )}
                      </MenuItem>
                    ))}
                  </MenuItems>
                </Menu>
              </div>
            </div>
          </div>

          <main className="py-10">
            <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
              <div className="grid auto-rows-min gap-4 md:grid-cols-3">
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <EncryptedVsPlainFilesChart />
                  </ChartContainer>
                </div>
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <PublicVsPrivateChart />
                  </ChartContainer>
                </div>
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <UploadsChart />
                  </ChartContainer>
                </div>
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <TenantUsageChart />
                  </ChartContainer>
                </div>
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <SharingTrendChart />
                  </ChartContainer>
                </div>
                <div className="bg-muted/50 aspect-video rounded-xl">
                  <ChartContainer>
                    <FileTypeChart />
                  </ChartContainer>
                </div>
              </div>
              <div className="bg-muted/50 min-h-screen flex-1 rounded-xl md:min-h-min">
                <SharedFilesPage />
              </div>

              <div className="bg-muted/50 min-h-screen flex-1 rounded-xl md:min-h-min">
                <SharedFilesPage />
              </div>
            </div>
          </main>
        </div>
      </div>
    </>
  );
}
