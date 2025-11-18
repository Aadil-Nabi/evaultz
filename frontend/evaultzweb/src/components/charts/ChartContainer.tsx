"use client";

export function ChartContainer({ children }: { children: React.ReactNode }) {
  return (
    <div className="bg-muted/50 rounded-xl p-4">
      <div className="relative h-[260px] w-full">{children}</div>
    </div>
  );
}
// className = "bg-muted/50 rounded-xl p-4 flex items-center justify-center";
