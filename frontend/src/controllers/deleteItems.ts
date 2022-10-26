import axios from "axios";
import { todoitemurl } from "./createItems";

export async function deleteItems(params: { id: number }[]) {
    const data = JSON.stringify(params);

    const config = {
        method: "delete",
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
