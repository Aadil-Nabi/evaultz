"use client";

import ReactEChart from "echarts-for-react";

export default function PublicVsPrivateChart() {
  const option = {
    tooltip: { trigger: "item" },
    series: [
      {
        name: "Visibility",
        type: "pie",
        radius: ["40%", "70%"],
        data: [
          { value: 180, name: "Private Files" },
          { value: 60, name: "Public Files" },
        ],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        Public vs Private
      </h3>

      <div className="flex-1">
        <ReactEChart
          option={option}
          style={{ height: "100%", width: "100%" }}
          notMerge
          lazyUpdate
        />
      </div>
    </div>
  );
}
