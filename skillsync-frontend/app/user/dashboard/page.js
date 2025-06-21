import { ContentLayout } from "@/components/admin-panel/content-layout";
import { cn } from "@/lib/utils";
import { Roboto } from "next/font/google";
import { SectionCards } from "@/components/dashboard/section-card";
import { StreakCalender } from "@/components/dashboard/streak-card";


const roboto = Roboto({
  subsets: ["latin"],
  weight: ["400", "500", "700"],
  variable: "--font-roboto",
});


export default function Page() {
  return (
    <ContentLayout title="DashBoard">
      <div>
        <h1 className={cn(
          "text-3xl font-medium text-zinc-900 dark:text-zinc-100",
          roboto.className,
        )}>Hi, Subhadip..</h1>

        <div className="flex gap-6 ">
          <div className="w-[80%]">
            <SectionCards />
          </div>
          <div className="w-[20%]">
            <div className="bg-zinc-50 dark:bg-zinc-800 rounded-lg h-full">
              <StreakCalender />
            </div>
          </div>
        </div>
      </div>
    </ContentLayout>
  );
}