import { Terminal, LineChart as LineChartIcon } from 'lucide-react';
import ActivityChart from './components/activity-chart';
import useSWR from 'swr';

const fetcher = (url: string) => fetch(url).then((res) => res.json());

export default function PersonalWebsite() {
  const { data, isLoading } = useSWR(
    'http://localhost:8080/chartinfo',
    fetcher,
  );

  if (isLoading) return <div>loading...</div>;

  return (
    <div className="h-screen bg-gray-900 text-green-500 font-mono flex flex-col">
      <div className="flex-grow flex flex-col overflow-hidden">
        <div className="bg-gray-700 p-2 flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <Terminal size={16} />
          </div>
        </div>

        <div className="flex-grow overflow-auto p-4 bg-gray-800">
          <div className="mb-8">
            Lorem ipsum dolor, sit amet consectetur adipisicing elit. Distinctio
            at qui quam mollitia aut sapiente nemo vel asperiores, minus
            corporis!
          </div>
          <div>Blogs (if i write any)</div>
          <div>Other pages links would go here maybe</div>

          <div className="bg-gray-700 p-4 rounded-lg mb-8">
            <div className="flex items-center space-x-2 mb-2">
              <LineChartIcon size={20} />
              <h2 className="text-lg font-bold">my activity</h2>
            </div>
            <div className="w-full md:h-[300px] sm:h-[200px]">
              <ActivityChart data={data} />
            </div>
          </div>
        </div>
      </div>

      {/* Status bar */}
      <footer className="bg-green-500 text-gray-900 p-1 text-sm">
        <div className="flex justify-between items-center">
          <span>NORMAL</span>
          <span>
            <a href="https://github.com/sama-004">github.com/sama-004</a>
          </span>
          <span>Line 1, Col 1</span>
        </div>
      </footer>
    </div>
  );
}
