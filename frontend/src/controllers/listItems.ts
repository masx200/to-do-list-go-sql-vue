import axios from "axios";
import { todoitemurl } from "./createItems";

export interface ToDoItemFull {
    id: number;
    author: string;
    content: string;
    completed: boolean;
}
export type ToDoItemResult = { data: ToDoItemFull[]; count: number };
export async function listItems(
    params: QueryParameters
): Promise<ToDoItemResult> {
    const config = {
        params: params,
        method: "get",
        url: todoitemurl,
        headers: {
            "Content-Type": "application/json",
        },
    };

    return axios(config).then(function (response) {
        console.log(response.data);
        return response.data;
    });
}
export interface QueryParameters {
    id?: number;
    limit?: number;
    page?: number;
    author?: string;
    order?: string;
    direction?: string;
    completed?: boolean;
}
