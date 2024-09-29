import { Line, LineChart, ResponsiveContainer, XAxis, YAxis } from 'recharts';

import {
  ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart';
import { useEffect, useState } from 'react';

interface ChartData {
  id: number;
  leftclick: number;
  rightclick: number;
  keypress: number;
  time: string;
}

interface ActivityChartProps {
  data: ChartData[];
}

const chartConfig = {
  leftclick: {
    label: 'Left Click',
    color: '#FFFF00',
  },
  rightclick: {
    label: 'Right Click',
    color: '#0000FF',
  },
  keypress: {
    label: 'Key Press',
    color: '#FFA500',
  },
} satisfies ChartConfig;

export default function ActivityChart({ data }: ActivityChartProps) {
  const [isSmallScreen, setIsSmallScreen] = useState(false);

  useEffect(() => {
    const checkScreenSize = () => {
      setIsSmallScreen(window.innerWidth < 640); // Adjust this value as needed
    };

    checkScreenSize();
    window.addEventListener('resize', checkScreenSize);

    return () => window.removeEventListener('resize', checkScreenSize);
  }, []);

  const chartData = data.map((item) => {
    const date = new Date(item.time);
    return {
      ...item,
      formattedTime: date.toLocaleTimeString([], {
        hour: '2-digit',
        minute: '2-digit',
        hour12: false,
      }),
      leftclick: item.leftclick,
      rightclick: item.rightclick,
      keypress: item.keypress,
    };
  });

  return (
    <ChartContainer
      config={chartConfig}
      className="md:h-[300px] sm:h-[200px] md:w-full sm:w-[200px]"
    >
      <ResponsiveContainer>
        <LineChart
          data={chartData}
          margin={{ top: 5, right: 10, left: 0, bottom: 5 }}
        >
          <XAxis dataKey="formattedTime" />
          <YAxis />
          <ChartTooltip cursor={false} content={<ChartTooltipContent />} />
          <ChartLegend content={<ChartLegendContent />} />
          <Line
            type="monotone"
            dataKey="keypress"
            stroke={chartConfig.keypress.color}
            strokeWidth={isSmallScreen ? 1 : 2}
            dot={false}
          />
          <Line
            type="monotone"
            dataKey="leftclick"
            stroke={chartConfig.leftclick.color}
            strokeWidth={isSmallScreen ? 1 : 2}
            dot={false}
          />
          <Line
            type="monotone"
            dataKey="rightclick"
            stroke={chartConfig.rightclick.color}
            strokeWidth={isSmallScreen ? 1 : 2}
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </ChartContainer>
  );
}
