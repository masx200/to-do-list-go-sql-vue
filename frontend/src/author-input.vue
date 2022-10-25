<template>
    <h1>author:</h1>
    <el-input
        v-model="author"
        placeholder=" input your name"
        @input="onchange"
    />
</template>
<script lang="ts" setup>
import { ref, watch } from "vue";
const props = defineProps<{
    input: string;
}>();
const emits = defineEmits<{ (event: "change", value: string): void }>();
function onchange(target: string) {
    emits("change", target);
}
const author = ref("");

watch(
    () => props.input,
    (v) => (author.value = v)
);
</script>
