<script lang="ts" setup>
import { ref } from "vue";
import authorInputVue from "./author-input.vue";

const author = ref("");
function onchange(target: string) {
    // console.log(target)
    author.value = target;

    // console.log(author)
}
</script>
<template>
    <main class="container" data-v-167ca4dc="">
        <header class="heading" data-v-73841b6c="" data-v-167ca4dc="">
            <div class="img-wrapper" data-v-73841b6c="">
                <img src="/note.75134fb0.svg" alt="Note" data-v-73841b6c="" />
            </div>
            <div class="title" data-v-73841b6c="">To-Do List</div>
        </header>
        <div class="form-field" data-v-5f8a7fba="" data-v-167ca4dc="">
            <authorInputVue :input="author" @change="onchange" />
            <h1 class="title" data-v-5f8a7fba="">~ Today I need to ~</h1>
            <form class="form-wrapper" data-v-5f8a7fba="">
                <div class="form-input" data-v-5f8a7fba="">
                    <input
                        placeholder="Add new todo..."
                        autofocus="true"
                        data-v-5f8a7fba=""
                    />
                </div>
                <button type="submit" class="submit-btn" data-v-5f8a7fba="">
                    <span data-v-5f8a7fba="">Submit</span>
                </button>
            </form>
        </div>
        <ul class="todos" data-v-132cabf7="" data-v-167ca4dc="">
            <li class="todos-item" data-v-132cabf7="">
                <button class="checkbox-btn" data-v-132cabf7="">
                    <div class="checkbox" data-v-132cabf7=""></div>
                </button>
                <p class="content" data-v-132cabf7="">3gggggg</p>
                <button class="close-btn" data-v-132cabf7="">
                    <svg
                        class="svg-inline--fa fa-times fa-w-11 close"
                        aria-hidden="true"
                        focusable="false"
                        data-prefix="fas"
                        data-icon="times"
                        role="img"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 352 512"
                        data-v-132cabf7=""
                    >
                        <path
                            class=""
                            fill="currentColor"
                            d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"
                        ></path>
                    </svg>
                </button>
            </li>
            <li class="todos-item" data-v-132cabf7="">
                <button class="checkbox-btn" data-v-132cabf7="">
                    <div class="checkbox" data-v-132cabf7=""></div>
                </button>
                <p class="content" data-v-132cabf7="">dddddddd</p>
                <button class="close-btn" data-v-132cabf7="">
                    <svg
                        class="svg-inline--fa fa-times fa-w-11 close"
                        aria-hidden="true"
                        focusable="false"
                        data-prefix="fas"
                        data-icon="times"
                        role="img"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 352 512"
                        data-v-132cabf7=""
                    >
                        <path
                            class=""
                            fill="currentColor"
                            d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"
                        ></path>
                    </svg>
                </button>
            </li>
            <li class="todos-item completed" data-v-132cabf7="">
                <button class="check-btn" data-v-132cabf7="">
                    <svg
                        class="svg-inline--fa fa-check fa-w-16 check"
                        aria-hidden="true"
                        focusable="false"
                        data-prefix="fas"
                        data-icon="check"
                        role="img"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 512 512"
                        data-v-132cabf7=""
                    >
                        <path
                            class=""
                            fill="currentColor"
                            d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                        ></path>
                    </svg>
                </button>
                <p class="content" data-v-132cabf7="">11111111111222gg</p>
                <button class="close-btn" data-v-132cabf7="">
                    <svg
                        class="svg-inline--fa fa-times fa-w-11 close"
                        aria-hidden="true"
                        focusable="false"
                        data-prefix="fas"
                        data-icon="times"
                        role="img"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 352 512"
                        data-v-132cabf7=""
                    >
                        <path
                            class=""
                            fill="currentColor"
                            d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"
                        ></path>
                    </svg>
                </button>
            </li>
        </ul>
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
