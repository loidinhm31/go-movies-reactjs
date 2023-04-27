import {boolean} from "boolean";
import type {AuthOptions, DefaultUser} from "next-auth";
import NextAuth from "next-auth";
import {Provider} from "next-auth/providers";
import CredentialsProvider from "next-auth/providers/credentials";
import KeycloakProvider from "next-auth/providers/keycloak";


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
            signIn: "/auth/signin"
            // TODO error: "/auth/error",
        },
        callbacks: {
            signIn: async function ({user, account, profile, credentials}) {
                if (account?.provider === "debug") {
                    if (boolean(process.env.DEBUG_LOGIN) || process.env.NODE_ENV === "development") {
                        return true;
                    }
                } else {
                    const payload = {
                        username: user.name,
                    };
                    const res = await fetch(`${process.env.API_BASE_URL}/auth/login`, {
                        method: "POST",
                        body: JSON.stringify(payload),
                        headers: {
                            "Content-Type": "application/json"
                        }
                    });
                    const userResponse = await res.json();
                    if (res.ok && userResponse) {
                        user.id = userResponse.username;
                        user.name = userResponse.first_name + ' ' + userResponse.last_name;
                        user.role = userResponse.role.role_name
                        return true;
                    }
                }
                return false
            },
            async jwt({token, user, account}) {
                if (account && user) {
                    return {
                        ...token,
                        accessToken: user.token,
                        role: user.role
                    };
                }
                return token;
            },
            async session({session, token}) {
                session.user.id = token.name!.toString();
                session.user.role = token.role!;
                session.user.name = token.name;
                session.user.accessToken = token.accessToken;
                return session;
            },

        },
        events: {
            async signIn() {
                console.log("handle event sign in");
            }
            ,
        },
        session: {
            strategy: "jwt"
        }
    }
;

export default NextAuth(authOptions);
