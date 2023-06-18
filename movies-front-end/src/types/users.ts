export type UserType = {
    id: number;
    username: string;
    email: string;
    first_name: string;
    last_name: string;
    is_new: boolean;
    role: RoleType;
    created_at: string | null;
    updated_at: string | null;
};

export type RoleType = {
    id?: number;
    role_name?: string;
    role_code: string;
};
