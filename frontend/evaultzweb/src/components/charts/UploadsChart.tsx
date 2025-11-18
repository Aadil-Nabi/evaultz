"use client";

import ReactEChart from "echarts-for-react";

export default function UploadsChart() {
  const option = {
    tooltip: { trigger: "axis" },
    xAxis: {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
    },
    yAxis: { type: "value" },
    series: [
      {
        name: "Uploads",
        type: "line",
        smooth: true,
        areaStyle: {},
        data: [50, 120, 90, 140, 200, 180, 220],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        Upload Activity
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
