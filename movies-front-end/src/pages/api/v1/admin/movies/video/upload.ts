import {withAnyRole} from "src/libs/auth";
import formidable from "formidable";
import {NextApiRequest, PageConfig} from "next";
import {Writable} from "stream";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    if (req.method !== "POST") return res.status(404).end();
    try {
        const chunks: never[] = [];
        const {fields, files} = await formidablePromise(req, {
            // consume this, otherwise formidable tries to save the file to disk
            fileWriteStreamHandler: () => fileConsumer(chunks),
        });
        
        const fileData = Buffer.concat(chunks); // or is it from? I always mix these up

        const form = new FormData();
        form.append("file", new Blob([fileData]));

        const headers = new Headers();
        headers.append("Authorization", `Bearer ${token.accessToken}`)

        const requestOptions = {
            method: "POST",
            headers: headers,
            body: form
        };

        const response = await fetch(`${process.env.API_BASE_URL}/integration/videos`, requestOptions);
        const message = await response.json();
        res.status(response.status).json(message.message! || "Failed to save movie");

    } catch (err) {
        return res.status(500).json({error: "Internal Server Error"});
    }
});

function formidablePromise(
    req: NextApiRequest,
    opts?: Parameters<typeof formidable>[0]
): Promise<{fields: formidable.Fields; files: formidable.Files}> {
    return new Promise((accept, reject) => {
        const form = formidable(opts);

        form.parse(req, (err, fields, files) => {
            if (err) {
                return reject(err);
            }
            return accept({fields, files});
        });
    });
}

const fileConsumer = <T = unknown>(acc: T[]) => {
    return new Writable({
        write: (chunk, _enc, next) => {
            acc.push(chunk);
            next();
        },
    });
};


export const config: PageConfig  = {
    api: {
        responseLimit: false,
        bodyParser: false,
    },
}


export default handler;
