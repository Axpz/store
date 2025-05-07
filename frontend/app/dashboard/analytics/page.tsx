"use client";

import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  PieChart,
  Pie,
  Cell,
  Legend,
  BarChart,
  Bar,
  LineChart,
  Line,
} from "recharts";
import { cn } from "@/lib/utils";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

// Example Data
const trendData = [
  { name: "Jan", value: 400 },
  { name: "Feb", value: 300 },
  { name: "Mar", value: 500 },
  { name: "Apr", value: 200 },
  { name: "May", value: 600 },
  { name: "Jun", value: 450 },
];

const insightData = [
  { category: "A", value: 30 },
  { category: "B", value: 40 },
  { category: "C", value: 20 },
  { category: "D", value: 10 },
];

const chartOverviewData = [
  { month: "Jan", sales: 2000, profit: 400, revenue: 2500 },
  { month: "Feb", sales: 1500, profit: 300, revenue: 2000 },
  { month: "Mar", sales: 2500, profit: 500, revenue: 3000 },
];

const userSegmentData = [
  { segment: "New", users: 150 },
  { segment: "Returning", users: 250 },
  { segment: "Active", users: 300 },
];

const reportData = [
  { date: "2024-01-01", views: 1200, downloads: 50 },
  { date: "2024-01-08", views: 1500, downloads: 75 },
  { date: "2024-01-15", views: 1800, downloads: 90 },
];

const settingsData = [
  { setting: "Option 1", value: "Enabled" },
  { setting: "Option 2", value: "Medium" },
  { setting: "Option 3", value: "7 days" },
];

const COLORS = ["#0088FE", "#00C49F", "#FFBB28", "#FF8042"];

interface PieLabelProps {
  cx: number;
  cy: number;
  midAngle: number;
  innerRadius: number;
  outerRadius: number;
  value: number;
  index: number;
}

import { useAuth } from "@/context/UserContext";

export default function AnalyticsPage() {
  const { user, isLoading } = useAuth();

  return (
    <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
        <h2 className="text-2xl font-bold text-gray-800 dark:text-gray-200">
          Analytics Dashboard Data
        </h2>

        <div className="grid grid-cols-1 gap-8">
          {/* Trend Analysis Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                Trend Analysis
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ResponsiveContainer width="100%" height={300}>
                <AreaChart data={trendData}>
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="name" />
                  <YAxis />
                  <Tooltip />
                  <Area
                    type="monotone"
                    dataKey="value"
                    stroke="#8884d8"
                    fill="#8884d8"
                  />
                </AreaChart>
              </ResponsiveContainer>
            </CardContent>
          </Card>

          {/* Data Insights Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                Data Insights
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ResponsiveContainer width="100%" height={300}>
                <PieChart>
                  <Pie
                    data={insightData}
                    cx="50%"
                    cy="50%"
                    labelLine={false}
                    outerRadius={80}
                    fill="#8884d8"
                    dataKey="value"
                    label={({
                      cx,
                      cy,
                      midAngle,
                      innerRadius,
                      outerRadius,
                      value,
                      index,
                    }: PieLabelProps) => {
                      const RADIAN = Math.PI / 180;
                      const radius =
                        25 + innerRadius + (outerRadius - innerRadius);
                      const x = cx + radius * Math.cos(-midAngle * RADIAN);
                      const y = cy + radius * Math.sin(-midAngle * RADIAN);

                      return (
                        <text
                          x={x}
                          y={y}
                          fill={COLORS[index % COLORS.length]}
                          textAnchor={x > cx ? "start" : "end"}
                          dominantBaseline="central"
                        >
                          {insightData[index].category} ({value})
                        </text>
                      );
                    }}
                  >
                    {insightData.map((entry, index) => (
                      <Cell
                        key={`cell-${index}`}
                        fill={COLORS[index % COLORS.length]}
                      />
                    ))}
                  </Pie>
                  <Legend />
                </PieChart>
              </ResponsiveContainer>
            </CardContent>
          </Card>

          {/* Chart Overview Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                Chart Overview
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ResponsiveContainer width="100%" height={300}>
                <BarChart data={chartOverviewData}>
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="month" />
                  <YAxis />
                  <Tooltip />
                  <Legend />
                  <Bar dataKey="sales" fill="#8884d8" />
                  <Bar dataKey="profit" fill="#82ca9d" />
                  <Bar dataKey="revenue" fill="#ffc658" />
                </BarChart>
              </ResponsiveContainer>
            </CardContent>
          </Card>

          {/* User Segmentation Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                User Segmentation
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ResponsiveContainer width="100%" height={300}>
                <LineChart data={userSegmentData}>
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="segment" />
                  <YAxis />
                  <Tooltip />
                  <Line
                    type="monotone"
                    dataKey="users"
                    stroke="#8884d8"
                    activeDot={{ r: 8 }}
                  />
                </LineChart>
              </ResponsiveContainer>
            </CardContent>
          </Card>

          {/* Report Generation Card */}
          {/* Report Generation Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                Report Generation
              </CardTitle>
            </CardHeader>
            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead className="text-left">Date</TableHead>
                    <TableHead className="text-left">Views</TableHead>
                    <TableHead className="text-left">Downloads</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {reportData.map((item, index) => (
                    <TableRow key={index}>
                      <TableCell className="text-left">{item.date}</TableCell>
                      <TableCell className="text-left">{item.views}</TableCell>
                      <TableCell className="text-left">
                        {item.downloads}
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CardContent>
          </Card>

          {/* Settings & Configuration Card */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-gray-700 dark:text-gray-300">
                Settings & Configuration
              </CardTitle>
            </CardHeader>
            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead className="text-left">Setting</TableHead>
                    <TableHead className="text-left">Value</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {settingsData.map((item, index) => (
                    <TableRow key={index}>
                      <TableCell className="text-left">
                        {item.setting}
                      </TableCell>
                      <TableCell className="text-left">{item.value}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CardContent>
          </Card>
        </div>
      </div>
  );
}