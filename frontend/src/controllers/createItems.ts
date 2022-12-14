import config from "../../config.json";
import axios from "axios";
export const todoitemurl = config.todoitem;
export async function createItems(item: ToDoItemNew[]) {
    const data = JSON.stringify(item);

    const config = {
        method: "post",
        url: todoitemurl,
        headers: {
            "Content-Type": "application/json",
        },
        data: data,
    };

    return axios(config).then(function (response) {
        console.log(response.data);
    });
}

export interface ToDoItemNew {
    author: string;
    content: string;
    completed: boolean;
}
