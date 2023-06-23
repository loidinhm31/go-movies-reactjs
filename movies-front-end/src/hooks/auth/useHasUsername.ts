import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

export const useHasUsername = () => {
  const session = useSession();
  const [author, setAuthor] = useState<string>();

  useEffect(() => {
    if (session) {
      if (session.data?.user && session.data.user.role !== "banned") {
        setAuthor(session.data?.user.id);
      } else {
        setAuthor("anonymous");
      }
    }
  }, [session]);

  return author;
};
