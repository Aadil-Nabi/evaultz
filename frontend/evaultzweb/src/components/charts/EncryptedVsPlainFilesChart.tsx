"use client";

import ReactEChart from "echarts-for-react";

export default function EncryptedVsPlainFilesChart() {
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
        padAngle: 5,
        itemStyle: {
          borderRadius: 10,
        },
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
          { value: 1048, name: "Total" },
          { value: 735, name: "Encrypted" },
          { value: 580, name: "Plain" },
        ],
      },
    ],
  };

  return (
    <div className="flex flex-col h-full">
      <h3 className="text-base font-semibold text-center pb-2">
        Encrypted vs Plain
      </h3>

      <div className="flex-1">
        <ReactEChart
          option={option}
          style={{ height: "100%", width: "100%" }}
        />
      </div>
    </div>
  );
}
