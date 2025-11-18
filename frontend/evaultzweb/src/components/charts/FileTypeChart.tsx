"use client";

import ReactEChart from "echarts-for-react";

export default function FileTypeChart() {
  const option = {
    tooltip: { trigger: "item" },
    series: [
      {
        name: "File Types",
        type: "pie",
        radius: "60%",
        data: [
          { value: 300, name: "PDF" },
          { value: 200, name: "Images" },
          { value: 150, name: "Videos" },
          { value: 100, name: "Docs" },
          { value: 80, name: "Others" },
        ],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        File Type Distribution
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
