"use client";

import ReactEChart from "echarts-for-react";

export default function SharingTrendChart() {
  const option = {
    tooltip: { trigger: "axis" },
    xAxis: { type: "category", data: ["Jan", "Feb", "Mar", "Apr", "May"] },
    yAxis: { type: "value" },
    series: [
      {
        name: "External Shares",
        type: "line",
        smooth: true,
        data: [20, 40, 35, 60, 80],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        External Sharing Trend
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
