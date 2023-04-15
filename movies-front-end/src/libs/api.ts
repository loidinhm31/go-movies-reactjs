import axios from "axios";

export interface Message {
	createdAt: string;
	id: string;
	sender: string;
	receiver: string;
	content: string;
	topicId: string;
}

export type InputMessage = {
    sender: string;
    receiver: string;
    content: string;
    topicId?: string;
}


const headers = {
    "Content-Type": "application/json",
};


const api = axios.create({
    headers,
});

export const get = (url: string) => api.get(url).then((res) => res.data);


export const post = (url: string, { arg: data }) => api.post(url, data).then((res) => res.data);

api.interceptors.response.use(
    (response) => response,
    (error) => {
        // TODO
    }
);

export default api;