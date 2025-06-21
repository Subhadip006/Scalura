import {
  TrendingUp,
  TrendingDown,
  BookOpenCheck,
  FolderKanban,
  GraduationCap,
} from "lucide-react";

import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export function SectionCards() {
  return (
    <div className=" *:data-[slot=card]:from-black/75 text-gray-200 *:data-[slot=card]:to-black dark:*:data-[slot=card]:bg-card grid sm:grid-cols-1 md:grid-cols-3 mt-4 gap-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:shadow-xs @xl/main:grid-cols-2 @5xl/main:grid-cols-4">
      
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Skill Progress</CardDescription>
          <CardTitle className="text-2xl font-semibold @[250px]/card:text-3xl tabular-nums">
            72%
          </CardTitle>
          <CardAction>
            <Badge variant="outline">
              <TrendingUp className="size-4 mr-1" />
              +8% this week
            </Badge>
          </CardAction>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="flex gap-2 font-medium">
            React, GoLang improving
            <TrendingUp className="size-4" />
          </div>
          <div className="text-muted-foreground">Based on logged activities</div>
        </CardFooter>
      </Card>

      {/* Projects Active */}
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Active Projects</CardDescription>
          <CardTitle className="text-2xl font-semibold @[250px]/card:text-3xl tabular-nums">
            3
          </CardTitle>
          <CardAction>
            <Badge variant="outline">
              <FolderKanban className="size-4 mr-1" />
              1 new this week
            </Badge>
          </CardAction>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="flex gap-2 font-medium">
            SkillFlow Backend in progress
          </div>
          <div className="text-muted-foreground">Updated 2 days ago</div>
        </CardFooter>
      </Card>

      {/* Academic Tasks */}
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Academic Tasks</CardDescription>
          <CardTitle className="text-2xl font-semibold @[250px]/card:text-3xl tabular-nums">
            5 Due
          </CardTitle>
          <CardAction>
            <Badge variant="outline">
              <TrendingDown className="size-4 mr-1" />
              +2 pending
            </Badge>
          </CardAction>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="flex gap-2 font-medium">
            Prioritize Research Logbook
            <BookOpenCheck className="size-4" />
          </div>
          <div className="text-muted-foreground">Due this weekend</div>
        </CardFooter>
      </Card>
      
    </div>
  );
}
