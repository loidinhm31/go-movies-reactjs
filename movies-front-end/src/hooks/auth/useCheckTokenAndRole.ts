import { useRouter } from "next/router";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { Role } from "@/components/RoleSelect";

export const useCheckTokenAndRole = (roles: Role[]) => {
  const router = useRouter();
  const { data: session, status } = useSession();

  const [isInvalid, setIsInvalid] = useState(false);

  useEffect(() => {
    if (status === "loading") {
      return;
    }
    if (session?.error === "RefreshAccessTokenError") {
      setIsInvalid(true); // Force sign in to hopefully resolve error
      return;
    }

    if (!roles.some((role) => role === session?.user?.role)) {
      router.push("/");
    }
  }, [router.pathname, session, status]);

  return isInvalid;
};
