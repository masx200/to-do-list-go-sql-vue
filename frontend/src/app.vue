<script lang="ts" src="./app.ts"></script>
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
                        required
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
                <span
                    :class="{ option: true, active: filterState === 0 }"
                    data-v-975e0b72=""
                    @click="clearquery"
                    >全部</span
                ><span
                    :class="{ option: true, active: filterState === 1 }"
                    data-v-975e0b72=""
                    @click="filternotcomplete"
                    >未完成</span
                ><span
                    :class="{ option: true, active: filterState === 2 }"
                    data-v-975e0b72=""
                    @click="filtercomplete"
                    >已完成</span
                >
            </div>
            <el-button size="large" @click="ondelete">
                <span class="option" data-v-975e0b72="">删除</span></el-button
            >
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

                <el-table-column property="author" label="作者" width="120" />
                <el-table-column property="content" label="内容" />
                <el-table-column
                    property="completed"
                    label="完成"
                    width="120"
                />
                <el-table-column label="操作">
                    <template #default="scope">
                        <el-button size="small" @click="handleToggle(scope.row)"
                            >切换</el-button
                        >
                        <el-button
                            size="small"
                            type="danger"
                            @click="handleDelete(scope.row)"
                            >删除</el-button
                        >
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <hr />

        <br />
    </main>
</template>
