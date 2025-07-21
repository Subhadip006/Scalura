import api from "../lib/axios";
import type { Skill, CreateSkillData, UpdateSkillData } from "../types/skill";

export const getSkills = async (userID: string): Promise<{ skills: Skill[] }> => {
  const response = await api.get(`/skills/user/${userID}`);
  return response.data;
};


export const createSkill = async (
  skillData: CreateSkillData
): Promise<{ skill: Skill }> => {
  const response = await api.post(`/skills`, skillData);
  return response.data;
};

export const updateSkill = async (
  skillId: string,
  skillData: UpdateSkillData
): Promise<{ skill: Skill }> => {
  const response = await api.put(`/skills/${skillId}`, skillData);
  return response.data;
};

export const deleteSkill = async (
  skillId: string
): Promise<{ message: string }> => {
  const response = await api.delete(`/skills/${skillId}`);
  return response.data;
};
