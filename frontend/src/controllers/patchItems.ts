import axios from "axios";
import { todoitemurl } from "./createItems";
import { ToDoItemFull } from "./listItems";

export async function patchItems(
    params: (Partial<ToDoItemFull> & { id: number })[]
) {
    const data = JSON.stringify(params);

    const config = {
        method: "PATCH",
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
