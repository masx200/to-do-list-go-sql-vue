import { createApp } from "vue";
import { notifyerror } from "./app";

import App from "./app.vue";

const app = createApp(App);
app.mount("#app");
app.config.errorHandler = function (error) {
    setTimeout(() => {
        throw error;
    });
    notifyerror(error);
};
