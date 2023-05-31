import {withAnyRole} from "src/libs/auth";
import httpProxyMiddleware from "next-http-proxy-middleware";


const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    httpProxyMiddleware(req, res, {
        target: `${process.env.API_BASE_URL}/auth/blobs/file`,
        pathRewrite: [{
            patternStr: req.url!,
            replaceStr: "",
        }],
        headers: {
            'Authorization': `Bearer ${token.accessToken}`,
        },
    });
});

// For preventing header corruption, specifically Content-Length header
export const config = {
    api: {
        bodyParser: false,
        externalResolver: true,
    },
}

export default handler;
