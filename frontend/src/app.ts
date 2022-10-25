import { ElNotification } from "element-plus";
import { defineComponent, onMounted, ref } from "vue";
import authorInput from "./author-input.vue";
import { createItem } from "./controllers/createItem";
import { deleteItems } from "./controllers/deleteItems";
import { listItems, ToDoItemFull } from "./controllers/listItems";
export function notifyerror(error: unknown) {
    ElNotification({
        title: "Error",
        message: String(error),
        type: "error",
        position: "top-left",
        duration: 10000,
    });
}
export default defineComponent({
    components: { authorInput },
    setup() {
        onMounted(async () => {
            await onquery();
        });
        const author = ref("");
        function onchange(target: string): void {
            author.value = target;
        }

        async function onsubmit(event: Event) {
            event.preventDefault();

            await createItem({
                author: author.value,
                completed: false,
                content: content.value,
            });

            await onquery();
        }

        const content = ref("");

        const listdata = ref([] as ToDoItemFull[]);

        const limit = 20;

        const page = ref(0);

        const direction = "desc";

        const order = "id";

        async function onquery() {
            listdata.value = await listItems({
                order,
                direction,
                page: page.value,
                limit,
            });
        }
        const multipleSelection = ref<ToDoItemFull[]>([]);
        const handleSelectionChange = (val: ToDoItemFull[]) => {
            multipleSelection.value = val;
        };
        async function ondelete() {
            if (multipleSelection.value.length === 0) return;
            await deleteItems(
                multipleSelection.value.map((a) => ({ id: a.id }))
            );
            multipleSelection.value = [];
            await onquery();
        }
        return {
            onquery,
            ondelete,
            handleSelectionChange,
            author,
            onchange,
            content,
            onsubmit,
            page,
            listdata,
        };
    },
});
