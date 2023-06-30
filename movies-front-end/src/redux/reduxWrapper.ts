import { Context, createWrapper, MakeStore } from "next-redux-wrapper";
import { Store } from "redux";
import store from "@/redux/store";


const makeStore: MakeStore<Store> = (context: Context) => store;

const wrapper = createWrapper<Store<Store>>(makeStore);

export default wrapper;