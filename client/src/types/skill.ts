import type { Goal } from "./goal";

export type Skill = {
    id: string;
    userID: string;
    Name: string;
    Level: "Beginner" | "Intermediate" | "Advanced";
    proof?: string;
    createdAt: string;
    goals: Goal[]

}

export type CreateSkillData = Omit<Skill, "id" | "createdAt" | "goals">;
export type UpdateSkillData = Omit<Skill, "id" | "createdAt" | "goals" | "userID">;