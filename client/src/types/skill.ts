import type { Goal } from "./goal";

export type skill = {
    id: string;
    userID: string;
    Name: string;
    Level: "Beginner" | "Intermediate" | "Advanced";
    proof?: string;
    createdAt: string;
    goals: Goal[]

}