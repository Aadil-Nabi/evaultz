"use client";
import { TextGenerateEffect } from "@/components/ui/text-generate-effect";

const words = ` A next-gen data vault that keeps your sensitive documents,
                encrypted files, and business assets fully protected —
                accessible only by you."
`;

export function WelcomeTextSectionPage() {
  return <TextGenerateEffect words={words} />;
}
