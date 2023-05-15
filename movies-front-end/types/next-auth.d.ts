import {DefaultSession} from "next-auth";
import {DateTime} from "next-auth/providers/kakao";

declare module "next-auth" {
    interface Session {
        error: any,
        user: {
            id: string;
            accessToken: string;
            role: string;
        } & DefaultSession["user"];
    }
}


declare module "next-auth/jwt" {
    interface JWT {
        id: string;
        role?: string;
        accessToken: string;
        expiresAt: number;
        refreshToken?: string;
    }
}


declare module "next-auth/core/types" {
    interface DefaultUser {
        id: string;
        token: string;
        role?: string;
        preferred_username?: string;
    }

    interface Profile {
        preferred_username?: string;
        given_name?: string;
        family_name?: string;
    }
}