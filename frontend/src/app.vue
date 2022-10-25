<script lang="ts" setup>
import { onMounted, ref } from "vue";
import authorInput from "./author-input.vue";
import { createItem } from "./controllers/createItem";
import { listItems, ToDoItemFull } from "./controllers/listItems";
onMounted(() => {
    onquery();
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
    content.value = "";
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
    console.log(multipleSelection);
};
</script>
<style>
div#app {
    width: 100%;
}
.options {
    margin-top: 0 !important;
}
body,
* {
    font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB,
        Microsoft YaHei, \5fae\8f6f\96c5\9ed1, Arial, sans-serif !important;
}
</style>
<template>
    <main class="container" data-v-167ca4dc="">
        <header class="heading" data-v-73841b6c="" data-v-167ca4dc="">
            <div class="img-wrapper" data-v-73841b6c="">
                <img src="/note.75134fb0.svg" alt="Note" data-v-73841b6c="" />
            </div>
            <div class="title" data-v-73841b6c="">To-Do List</div>
        </header>

        <div class="form-field" data-v-5f8a7fba="" data-v-167ca4dc="">
            <authorInput :input="author" @change="onchange" />
            <h1 class="title" data-v-5f8a7fba="">今天日程:</h1>
            <form class="form-wrapper" data-v-5f8a7fba="">
                <div class="form-input" data-v-5f8a7fba="">
                    <input
                        v-model="content"
                        placeholder="Add new todo..."
                        autofocus="true"
                        data-v-5f8a7fba=""
                    />
                </div>
                <button
                    type="submit"
                    class="submit-btn"
                    data-v-5f8a7fba=""
                    @click="onsubmit"
                >
                    <span data-v-5f8a7fba="">提交</span>
                </button>
            </form>
        </div>
        <hr />
        <br />
        <div style="display: flex; justify-content: space-around">
            <span>页数</span>
            <el-input-number
                :stepStrictly="true"
                :min="0"
                :step="1"
                v-model="page"
                placeholder="page"
            />
            <el-button size="large" @click="onquery">查询</el-button>
        </div>
        <div class="options" data-v-975e0b72="" data-v-167ca4dc="">
            <div class="filters" data-v-975e0b72="">
                <span class="option active" data-v-975e0b72="">全部</span
                ><span class="option" data-v-975e0b72="">未完成</span
                ><span class="option" data-v-975e0b72="">已完成</span>
            </div>
            <span class="option" data-v-975e0b72="">删除</span>
        </div>
        <br />
        <div>
            <el-table
                :data="listdata"
                style="width: 100%; font-size: 20px"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" />
                <el-table-column property="id" label="序号" width="120" />
                <el-table-column
                    property="completed"
                    label="完成"
                    width="120"
                />
                <el-table-column property="author" label="作者" width="120" />
                <el-table-column property="content" label="内容" />
            </el-table>
        </div>

        <hr />

        <br />
    </main>
</template>
