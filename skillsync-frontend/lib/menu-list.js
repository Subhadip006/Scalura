import {
  LayoutGrid,
  BadgeCheck,
  FolderKanban,
  GraduationCap,
  BookOpen,
  Settings
} from "lucide-react";

export function getMenuList(pathname) {
  return [
    {
      groupLabel: "",
      menus: [
        {
          href: "/dashboard",
          label: "Dashboard",
          icon: LayoutGrid,
          submenus: []
        }
      ]
    },
    {
      groupLabel: "SkillFlow Tools",
      menus: [
        {
          href: "/skills",
          label: "Skills",
          icon: BadgeCheck,
          submenus: []
        },
        {
          href: "/projects",
          label: "Projects",
          icon: FolderKanban,
          submenus: []
        },
        {
          href: "/academics",
          label: "Academic Planner",
          icon: GraduationCap,
          submenus: []
        },
      ]
    },
    {
      groupLabel: "Settings",
      menus: [
        {
          href: "/account",
          label: "Profile",
          icon: Settings,
          submenus: [
            {
              href: "/account/profile",
              label: "Profile Info"
            },
            {
              href: "/account/security",
              label: "Security"
            }
          ]
        }
      ]
    }
  ];
}
