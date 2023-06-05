import {useRouter} from "next/router";
import MovieCheckout from "src/components/Payment/MovieCheckout";
import EpisodeCheckout from "src/components/Payment/EpisodeCheckout";
import {useSession} from "next-auth/react";

export default function Payment() {
    const router = useRouter();
    const {data: session} = useSession();

    const {type, refId} = router.query;

    return (
        <>

            {type === "MOVIE" &&
                <MovieCheckout
                    refId={Number(refId)}
                    type={type}
                />
            }

            {type === "TV" &&
                <EpisodeCheckout
                    refId={Number(refId)}
                    type={type}
                />
            }
        </>
    );
}

