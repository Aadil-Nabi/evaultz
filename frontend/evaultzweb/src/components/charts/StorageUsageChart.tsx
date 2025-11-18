"use client";

import ReactEChart from "echarts-for-react";

export default function StorageUsageChart() {
  const option = {
    tooltip: {
      trigger: "item",
    },
    legend: {
      top: "5%",
      left: "center",
    },
    series: [
      {
        name: "Access From",
        type: "pie",
        radius: ["40%", "70%"],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: "center",
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 40,
            fontWeight: "bold",
          },
        },
        labelLine: {
          show: false,
        },
        data: [
          { value: 1048, name: "PDF" },
          { value: 735, name: "docs" },
          { value: 580, name: "emails" },
          { value: 484, name: "videos" },
          { value: 300, name: "pictures" },
        ],
      },
    ],
  };
  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        Storage Usage
      </h3>

      <div className="flex-1">
        <ReactEChart
          option={option}
          style={{ height: "100%", width: "100%" }}
          notMerge={true}
          lazyUpdate={true}
        />
      </div>
    </div>
  );
}
