import { Alert, AlertColor, Snackbar, SnackbarOrigin } from "@mui/material";
import { useEffect } from "react";

export interface NotifyState extends SnackbarOrigin {
  open?: boolean;
  message?: string;
  severity?: AlertColor;
}

interface NotifyProps {
  state: NotifyState;
  setState: (state: NotifyState) => void;
}

export function sleep(delay = 0) {
  return new Promise((resolve) => {
    setTimeout(resolve, delay);
  });
}

export default function NotifySnackbar({ state, setState }: NotifyProps) {
  const { vertical, horizontal, open, message, severity } = state;

  useEffect(() => {
    if (!open) {
      (async () => {
        await setState(state);
        await sleep(5000);
        await setState({ ...state, open: false });
      })();
    }
  }, [state]);

  return (
    <div>
      <Snackbar anchorOrigin={{ vertical, horizontal }} open={open} key={vertical + horizontal}>
        <Alert severity={severity}>{message}</Alert>
      </Snackbar>
    </div>
  );
}
