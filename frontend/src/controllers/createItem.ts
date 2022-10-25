import config from "../../config.json";
export const todoitemurl = config.todoitem;
export async function createItem(item: ToDoItemNew) {
    const myHeaders = new Headers();

    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify([item]);

    const requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow",
    } as const;

    return fetch(todoitemurl, requestOptions)
        .then((response) => response.json())
        .then((result) => console.log(result))
        .catch((error) => console.log("error", error));
}

export interface ToDoItemNew {
    author: string;
    content: string;
    completed: boolean;
}
