import { useState, useEffect } from 'react';
import { Terminal, BarChart2 } from 'lucide-react';

export default function PersonalWebsite() {
  const [blinkCursor, setBlinkCursor] = useState(true);

  useEffect(() => {
    const interval = setInterval(() => {
      setBlinkCursor((prev) => !prev);
    }, 530);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="h-screen bg-gray-900 text-green-500 font-mono flex flex-col">
      <div className="flex-grow flex flex-col overflow-hidden">
        <div className="bg-gray-700 p-2 flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <Terminal size={16} />
            <span className="text-sm">samanyuroy@gmail.com</span>
          </div>
          {/*Neovim svg here*/}
          VIM LOGO GOES HERE
        </div>

        <div className="flex-grow overflow-auto p-4 bg-gray-800">
          <div className="mb-4">
            <span className="text-green-400">$</span> cat welcome.txt
          </div>
          <div className="mb-8">
            Welcome to my personal website! I'm a developer who loves Linux and
            Vim. Explore my projects and get in touch if you'd like to
            collaborate.
          </div>
          <div>Blogs (if i write any)</div>
          <div>Other pages links would go here maybe</div>

          <div className="bg-gray-700 p-4 rounded-lg mb-8">
            <div className="flex items-center space-x-2 mb-2 h-[500px]">
              <BarChart2 size={20} />
              <h2 className="text-lg font-bold">Activity Chart</h2>
            </div>
            <div className="flex items-center justify-center border border-dashed border-green-500 rounded">
              Chart coming soon...
            </div>
          </div>

          <div className="flex items-center">
            <span className="text-green-400">$</span>
            <span className="ml-2">{blinkCursor ? '|' : ' '}</span>
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
