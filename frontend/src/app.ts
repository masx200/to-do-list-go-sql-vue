import { ElNotification, ElMessage } from "element-plus";
import { defineComponent, onMounted, ref, watch } from "vue";
import authorInput from "./author-input.vue";
import { createItems } from "./controllers/createItems";
import { deleteItems } from "./controllers/deleteItems";
import {
    listItems,
    QueryParameters,
    ToDoItemFull,
} from "./controllers/listItems";
import { patchItems } from "./controllers/patchItems";

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
            if (!author.value || !content.value)
                return ElMessage.error("请输入内容,作者或者内容不能为空");
            await createItems([
                {
                    author: author.value,
                    completed: false,
                    content: content.value,
                },
            ]);

            await onquery();
        }
        const total = ref(0);
        const content = ref("");

        const listdata = ref([] as ToDoItemFull[]);

        const limit = 20;

        const currentPage4 = ref(1);

        const direction = "desc";

        const order = "id";
        let query: QueryParameters = {};
        async function onquery() {
            const result = await listItems({
                order,
                direction,
                page: currentPage4.value,
                limit: pageSize4.value,
                ...query,
            });
            listdata.value = result.data;
            total.value = result.count;
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

        const filterState = ref(0);
        watch(filterState, async () => await onquery());
        const clearquery = () => {
            query = {};
            filterState.value = 0;
        };
        function filternotcomplete() {
            query = { completed: false };
            filterState.value = 1;
        }
        function filtercomplete() {
            query = { completed: true };
            filterState.value = 2;
        }
        const handleToggle = async (row: ToDoItemFull) => {
            await patchItems([{ id: row.id, completed: !row.completed }]);
            await onquery();
        };
        const handleDelete = async (row: ToDoItemFull) => {
            await deleteItems([{ id: row.id }]);
            await onquery();
        };
        const pageSize4 = ref(limit);
        const handleCurrentChange = (val: number) => {
            return onquery();
        };
        const handleSizeChange = (val: number) => {
            return onquery();
        };
        return {
            handleCurrentChange,
            handleSizeChange,
            currentPage4,
            pageSize4,
            background: true,
            disabled: false,
            size: "large",
            handleToggle,
            handleDelete,
            filterState: filterState,
            clearquery,
            onquery,
            ondelete,
            handleSelectionChange,
            author,
            onchange,
            content,
            total,
            onsubmit,

            listdata,
            filtercomplete,
            filternotcomplete,
        };
    },
});
