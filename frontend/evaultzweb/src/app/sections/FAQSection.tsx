import {
  Disclosure,
  DisclosureButton,
  DisclosurePanel,
} from "@headlessui/react";
import { MinusSmallIcon, PlusSmallIcon } from "@heroicons/react/24/outline";

const faqs = [
  {
    question: "What is Evaultz?",
    answer:
      "Evaultz is a secure file storage, encryption, and digital vault platform designed for individuals and organizations that need enterprise-grade security for their data",
  },
  {
    question: "Who is Evaultz designed for?",
    answer:
      "Businesses, cybersecurity teams, developers, IT admins, and individuals who handle sensitive documents or want complete control over their data.",
  },
  {
    question: "What makes Evaultz different from other storage platforms",
    answer:
      "Evaultz focuses heavily on encryption, privacy, compliance, and user-controlled keys, making it more secure than traditional cloud storage systems.",
  },
  {
    question: "Can I upload any file type?",
    answer: "Yes, unless your organization has custom restrictions.",
  },
  {
    question: "Does Evaultz support file sharing?",
    answer:
      "Yes, with secure sharing links that can be time-bound and permission-controlled.",
  },
  {
    question: "How do I delete my Evaultz account?",
    answer:
      "You can delete it from the profile page, which permanently removes your vault & all encrypted files.",
  },
];

export default function FAQSection() {
  return (
    <div className="bg-white dark:bg-gray-900">
      <div className="mx-auto max-w-7xl px-6 py-24 sm:py-32 lg:px-8 lg:py-40">
        <div className="mx-auto max-w-4xl">
          <h2 className="text-4xl font-semibold tracking-tight text-gray-900 sm:text-5xl dark:text-white">
            Frequently asked questions
          </h2>
          <dl className="mt-16 divide-y divide-gray-900/10 dark:divide-white/10">
            {faqs.map((faq) => (
              <Disclosure
                key={faq.question}
                as="div"
                className="py-6 first:pt-0 last:pb-0"
              >
                <dt>
                  <DisclosureButton className="group flex w-full items-start justify-between text-left text-gray-900 dark:text-white">
                    <span className="text-base/7 font-semibold">
                      {faq.question}
                    </span>
                    <span className="ml-6 flex h-7 items-center">
                      <PlusSmallIcon
                        aria-hidden="true"
                        className="size-6 group-data-open:hidden"
                      />
                      <MinusSmallIcon
                        aria-hidden="true"
                        className="size-6 group-not-data-open:hidden"
                      />
                    </span>
                  </DisclosureButton>
                </dt>
                <DisclosurePanel as="dd" className="mt-2 pr-12">
                  <p className="text-base/7 text-gray-600 dark:text-gray-400">
                    {faq.answer}
                  </p>
                </DisclosurePanel>
              </Disclosure>
            ))}
          </dl>
        </div>
      </div>
    </div>
  );
}
