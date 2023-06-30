import { configureStore, createSlice, PayloadAction } from "@reduxjs/toolkit";

interface DataState {
  severity?: string;
  message?: string;
}

const dataSlice = createSlice({
  name: "data",
  initialState: {},
  reducers: {
    setData: (state, action: PayloadAction<DataState>) => {
      return action.payload;
    }
  }
});

const store = configureStore({
  reducer: {
    data: dataSlice.reducer
  }
});

export const { setData } = dataSlice.actions;

export default store;


