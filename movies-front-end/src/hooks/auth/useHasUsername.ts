import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

export const useHasUsername = () => {
    const session = useSession();
    const [author, setAuthor] = useState("anonymous");

    useEffect(() => {
        if (session && session.data?.user) {
            setAuthor(session.data.user.id);
        }
    }, [session]);

    return author;
};
