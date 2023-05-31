import {boolean} from "boolean";
import type {AuthOptions, DefaultUser} from "next-auth";
import NextAuth from "next-auth";
import {OAuthConfig, Provider} from "next-auth/providers";
import CredentialsProvider from "next-auth/providers/credentials";
import KeycloakProvider, {KeycloakProfile} from "next-auth/providers/keycloak";
import {NextApiRequest} from "next";
import {MimetypesKind} from "video.js/dist/types/utils/mimetypes";
import oga = MimetypesKind.oga;
import {JWT} from "next-auth/jwt";

async function refreshAccessToken(token) {
    try {
        console.log("Refreshing token...");
        const response = await fetch(`${process.env.KEYCLOAK_ISSUER}/protocol/openid-connect/token`, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
                "Authorization": `Basic ${btoa(`${process.env.KEYCLOAK_CLIENT_ID}:${process.env.KEYCLOAK_CLIENT_SECRET}`)}`,
            },
            body: new URLSearchParams({
                grant_type: "refresh_token",
                refresh_token: token.refreshToken!,
            }),
            method: "POST",
        })
        const refreshedTokens = await response.json();
        if (!response.ok) {
            throw refreshedTokens
        }

        return {
            ...token, // Keep the previous token properties
            accessToken: refreshedTokens.access_token,
            expiresAt: Date.now() + refreshedTokens.expires_in * 1000,
            // Fall back to old refresh token, but note that
            // many providers may only allow using a refresh token once.
            refreshToken: refreshedTokens.refresh_token,
        }
    } catch (error) {
        console.error("Error refreshing access token", error)
        // The error property will be used client-side to handle the refresh token error
        return {...token, error: "RefreshAccessTokenError" as const}
    }
}

const providers: Provider[] = [];
if (boolean(process.env.DEBUG_LOGIN) || process.env.NODE_ENV === "development") {
    providers.push(
        CredentialsProvider({
            id: "debug",
            name: "Debug Credentials",
            credentials: {
                username: {label: "Username", type: "text"},
                role: {label: "Role", type: "text"}
            },
            async authorize(credentials) {
                const user: DefaultUser = {
                    id: credentials!.username,
                    name: credentials!.username,
                    role: credentials!.role,
                    token: ""
                };
                return user;
            }
        })
    );
}

providers.push(
    CredentialsProvider({
        id: "server",
        name: "Credentials",
        credentials: {
            username: {label: "Username", type: "text"},
            password: {label: "Password", type: "text"}
        },
        async authorize(credentials) {
            const payload = {
                username: credentials!.username,
                password: credentials!.password
            };
            const res = await fetch(`${process.env.API_BASE_URL}/api/v1/auth/login`, {
                method: "POST",
                body: JSON.stringify(payload),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            const userWithToken = await res.json();
            if (res.ok && userWithToken) {
                const user: DefaultUser = {
                    id: userWithToken.user.username,
                    name: userWithToken.user.username,
                    role: "general", // TODO,
                    token: userWithToken.token
                };
                return user;
            }
            // Return null if user data could not be retrieved
            return null;
        }
    })
);

if (process.env.KEYCLOAK_CLIENT_ID) {
    providers.push(
        KeycloakProvider({
            clientId: process.env.KEYCLOAK_CLIENT_ID,
            clientSecret: process.env.KEYCLOAK_CLIENT_SECRET!,
            issuer: process.env.KEYCLOAK_ISSUER,
        })
    );
}

export const authOptions: AuthOptions = {
        providers,
        pages: {
            signIn: "/auth/signin",
            // TODO error: "/auth/error",
        },
        callbacks: {
            signIn: async function ({user, account, profile, credentials}) {
                if (account?.provider === "debug") {
                    if (boolean(process.env.DEBUG_LOGIN) || process.env.NODE_ENV === "development") {
                        return true;
                    }
                } else {
                    if (profile) {
                        const payload = {
                            username: profile.preferred_username,
                        };
                        const res = await fetch(`${process.env.API_BASE_URL}/auth/login`, {
                            method: "POST",
                            body: JSON.stringify(payload),
                            headers: {
                                "Content-Type": "application/json",
                                "Authorization": `Bearer ${account?.access_token}`
                            }
                        });
                        const userResponse = await res.json();
                        if (res.ok && userResponse) {
                            user.id = userResponse.username;
                            user.name = userResponse.first_name + ' ' + userResponse.last_name;
                            user.role = userResponse.role.role_name;
                            user.token = account?.access_token!;
                            return true;
                        } else {    // Create pending user into the database
                            const registerPayload = {
                                username: profile.preferred_username,
                                first_name: profile.given_name,
                                last_name: profile.family_name,
                                email: profile.email,
                                is_new: true,
                            };
                            const registerRes = await fetch(`${process.env.API_BASE_URL}/auth/signup`, {
                                method: "POST",
                                body: JSON.stringify(registerPayload),
                                headers: {
                                    "Content-Type": "application/json",
                                    "Authorization": `Bearer ${account?.access_token}`
                                }
                            });
                        }
                    }
                }
                return false;
            },
            async jwt({token, user, account}) {
                if (account && user) {
                    return {
                        ...token,
                        id_token: account.id_token,
                        id: user.id,
                        accessToken: user.token,
                        expiresAt: account.expires_at!,
                        refreshToken: account.refresh_token,
                        role: user.role,
                    };
                } else if (Date.now() < token.expiresAt!) {
                    // If the access token has not expired yet, return it
                    return token;
                }
                return refreshAccessToken(token);
            },
            async session({session, token }) {
                if (token) {
                    session.user.id = token.id;
                    session.user.role = token.role!;
                    session.user.name = token.name;
                    session.user.accessToken = token.accessToken;
                    session.error = token.error
                }
                return session;
            },
        },
        events: {
            async signIn() {
                console.log("handle event sign in");
            },
            async signOut({session, token}) {
                const issuerUrl = (authOptions.providers.find(p => p.id === "keycloak") as OAuthConfig<KeycloakProfile>).options!.issuer!
                const logOutUrl = new URL(`${issuerUrl}/protocol/openid-connect/logout`)
                logOutUrl.searchParams.set("id_token_hint", token.id_token!)
                await fetch(logOutUrl);
            }
        },
        session: {
            strategy: "jwt"
        }
    }
;

const getIp = (req: NextApiRequest) => {
    try {
        // https://stackoverflow.com/questions/66111742/get-the-client-ip-on-nextjs-and-use-ssr
        const forwarded = req.headers["x-forwarded-for"];
        return typeof forwarded === "string" ? forwarded.split(/, /)[0] : req.socket.remoteAddress;
    } catch {
        return "";
    }
};


export default NextAuth(authOptions);
