"use client";

import { Calendar } from "@/components/ui/calendar";
import { useState } from "react";

export function StreakCalender() {
  const [streak, setStreak] = useState([
    new Date(), // today
    new Date(2024, 5, 20)
  ]);

  const toggleDate = (date) => {
    const exists = streak.some(d => d.toDateString() === date.toDateString());
    setStreak((prev) =>
      exists ? prev.filter(d => d.toDateString() !== date.toDateString()) : [...prev, date]
    );
  };

  return (
    <div className="bg-zinc-50 dark:bg-zinc-800 rounded-lg p-4 h-full">
      <h2 className="text-xl font-semibold mb-4">Streak Calendar</h2>
      <div className="flex justify-center">
        <Calendar
          streak={streak}
          onToggleDate={toggleDate}
          className="w-full max-w-md"
        />
      </div>
    </div>
  );
}
