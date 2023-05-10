import axios from "axios";
import {ClientError} from "./api_client";

const headers = {
    "Content-Type": "application/json",
};

const formApi = axios.create();

export const postForm = (url: string, data?: { arg: any }) => formApi.post(url, data?.arg).then((res) => res.data);

const api = axios.create({headers});

export const get = (url: string) => api.get(url).then((res) => res.data);

export const post = (url: string, data?: { arg: any }) => api.post(url, data?.arg).then((res) => res.data);

export const del = (url: string) => api.delete(url).then((res) => res.data);

export const put = (url: string, data?: { arg: any }) => api.put(url, data?.arg).then((res) => res.data);

api.interceptors.response.use(
    (response) => response,
    (error) => {
        const err = error?.response;
        throw new ClientError({
            message: err?.data || "",
            errorCode: err?.errorCode,
            httpStatusCode: error?.status || -1,
        });
    }
);
export default api;