import {forwardRef} from "react";
import {ElementOf} from "../types/utils";
import {InputLabel, MenuItem, Select, SelectProps, TextField} from "@mui/material";

export const roles = ["general", "admin", "banned"] as const;
export type Role = ElementOf<typeof roles>;

type RoleSelectProps = Omit<SelectProps<any>, "defaultValue"> & {
	defaultValue?: Role;
	value?: Role;
};

export const RoleSelect = forwardRef<HTMLSelectElement, RoleSelectProps>((props, ref) => {
	return (
		<TextField
			select
			label="Role"
			defaultValue={roles[0]}
			variant="filled"
		>
			{roles.map((role) => (
				<MenuItem key={role} value={role}>
					{role}
				</MenuItem>
			))}
		</TextField>
	);
});

RoleSelect.displayName = "RoleSelect";
