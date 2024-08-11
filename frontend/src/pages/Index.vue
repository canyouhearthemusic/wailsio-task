<script setup>
import { PencilIcon, TrashIcon, CheckIcon } from "@heroicons/vue/24/outline";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { CreateTodo, ListTodos } from "_wailsjs";
import { onBeforeMount, reactive, ref } from "vue";

const form = reactive({
    body: "",
    priority: "",
    datetime: "",
});

const todos = ref([]);

onBeforeMount(async () => {
    await fetchTodos();
});

async function fetchTodos() {
    try {
        const result = await ListTodos();
        todos.value = result;
    } catch (error) {
        console.error("Failed to fetch todos:", error);
    }
}

async function createTodo() {
    try {
        await CreateTodo(form.body, form.priority, form.datetime);
        await fetchTodos();
    } catch (error) {
        console.error("Failed to create todo:", error);
    }
}
</script>

<template>
    <div>
        <h1 class="text-2xl font-semibold text-center">Todo List</h1>
    
        <Input
            id="body"
            v-model="form.body"
            type="text"
            class="mt-4 bg-white"
            placeholder="What to do?"
        />
    
        <div class="w-full mt-4 flex gap-x-3">
            <Select v-model="form.priority">
                <SelectTrigger class="w-full">
                    <SelectValue placeholder="Select a priority" />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Priority</SelectLabel>
                        <SelectItem value="low"> Low </SelectItem>
                        <SelectItem value="mid"> Mid </SelectItem>
                        <SelectItem value="high"> High </SelectItem>
                    </SelectGroup>
                </SelectContent>
            </Select>
        </div>
    
        <Input
            id="date"
            v-model="form.datetime"
            type="datetime-local"
            class="mt-4 bg-white"
            placeholder="Deadline"
        />
    
        <Button class="w-full mt-3" @click="createTodo"> Create </Button>
    </div>

    <div class="w-full grid grid-cols-4 gap-10 mt-6">
        <div
            v-for="todo in todos"
            :key="todo.id"
            class="border border-gray-300 rounded-md px-4 py-2.5 w-64 max-w-64 max-h-[400px] overflow-scroll"
        >
            <header class="flex justify-between items-center">
                <p class="font-semibold text-2xl">{{ todo.body }}</p>

                <div class="flex gap-x-2">
                    <PencilIcon
                        class="size-5 text-black hover:text-gray-500 transition-all"
                    />
                    <TrashIcon
                        class="size-5 text-black hover:text-red-600 transition-all"
                    />
                    <CheckIcon
                        class="size-5 text-black hover:text-green-500 transition-all"
                    />
                </div>
            </header>

            <main class="mt-3">
                <p>
                    Status:
                    <span>{{ todo.status }}</span>
                </p>
            </main>

            <footer class="mt-4 border-t border-gray-800 pt-2">
                <p class="text-sm italic text-gray-700">
                    Created at: {{ todo.created_at }}
                </p>
                <p class="text-sm italic text-gray-700">
                    Last changed at: {{ todo.updated_at }}
                </p>
                <p class="text-sm italic text-gray-700">
                    Deadline: {{ todo.deadline }}
                </p>
            </footer>
        </div>
    </div>
</template>
