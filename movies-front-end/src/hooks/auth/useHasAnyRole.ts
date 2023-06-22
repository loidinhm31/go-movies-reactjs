import { useSession } from "next-auth/react";
import { Role } from "@/components/RoleSelect";

export const useHasAnyRole = (roles: Role[]) => {
  const { data: session } = useSession();

  return roles.some((role) => role === session?.user?.role);
};
