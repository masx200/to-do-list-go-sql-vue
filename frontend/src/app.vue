<script lang="ts" setup>
import { ref } from "vue";
import authorInput from "./author-input.vue";
import { createItem } from "./controllers/createItem";
import { ToDoItemFull } from "./controllers/listItems";

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

const listdata: ToDoItemFull[] = [
    { id: 1, completed: false, content: "haha", author: "djw " },
    { id: 21, completed: false, content: "h22aha", author: "d222jw " },
];

const limit = 10;

const page = ref(0);

const direction = "desc";

const order = "id";
</script>
<style>
div#app {
    width: 100%;
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
        <div style="display: flex">
            <span>页数</span>
            <el-input v-model="page" placeholder="page" />
            <el-button size="large">查询</el-button>
        </div>
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
        <div>
            <el-table :data="listdata" style="width: 100%; font-size: 20px">
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

        <footer class="options" data-v-975e0b72="" data-v-167ca4dc="">
            <span data-v-975e0b72="">2 item left</span>
            <div class="filters" data-v-975e0b72="">
                <span class="option active" data-v-975e0b72="">All</span
                ><span class="option" data-v-975e0b72="">Active</span
                ><span class="option" data-v-975e0b72="">Completed</span>
            </div>
            <span class="option" data-v-975e0b72="">Clear completed</span>
        </footer>
    </main>
</template>
