import axios from "axios";
import { todoitemurl } from "./createItem";

export async function deleteItems(params: { id: number }[]) {
    var data = JSON.stringify(params);

    var config = {
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
