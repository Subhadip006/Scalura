import type { Task } from "./task";
export type Goal = {
  id: string;
  skill_id: string;
  title: string;
  description: string;
  target_date: string;     
  completed: boolean;
  created_at: string;
  tasks: Task[];           
};

export type CreateGoalInput = Omit<Goal, "id" | "created_at" | "tasks">;

export type UpdateGoalInput = Partial<Omit<Goal, "id" | "created_at" | "skill_id" | "tasks">>;
