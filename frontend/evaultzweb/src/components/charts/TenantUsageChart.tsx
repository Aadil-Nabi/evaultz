"use client";

import ReactEChart from "echarts-for-react";

export default function TenantUsageChart() {
  const option = {
    tooltip: { trigger: "axis" },
    xAxis: {
      type: "category",
      data: ["Tenant A", "Tenant B", "Tenant C", "Tenant D"],
    },
    yAxis: { type: "value" },
    series: [
      {
        name: "Storage (GB)",
        type: "bar",
        barWidth: "50%",
        data: [120, 200, 150, 80],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">Tenant Usage</h3>

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
