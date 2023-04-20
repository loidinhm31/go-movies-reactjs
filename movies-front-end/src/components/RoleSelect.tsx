import {forwardRef} from "react";
import {ElementOf} from "../types/utils";
import {FormControl, InputLabel, MenuItem, Select, SelectProps, TextField} from "@mui/material";

export const roles = ["general", "admin", "banned"] as const;
export type Role = ElementOf<typeof roles>;

type RoleSelectProps = Omit<SelectProps<any>, "defaultValue"> & {
	defaultValue?: Role;
	value?: Role;
};

export const RoleSelect = forwardRef<HTMLSelectElement, RoleSelectProps>((props, ref) => {
	return (
		<>
			<FormControl>
				<InputLabel id="demo-simple-select-standard-label">Role</InputLabel>
				<Select
					id="demo-simple-select-standard-label"
					label="Role"
					defaultValue={roles[0]}
					{...props}
					ref={ref}
				>

					{roles.map((role) => (
						<MenuItem key={role} value={role}>
							{role}
						</MenuItem>
					))}
				</Select>
			</FormControl>
		</>
	);
});

RoleSelect.displayName = "RoleSelect";
