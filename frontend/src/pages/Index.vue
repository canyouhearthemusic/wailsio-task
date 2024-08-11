<script setup>
import { TrashIcon, CheckIcon } from "@heroicons/vue/24/outline";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { CreateTodo, ListTodos, DeleteTodo, ToggleStatusTodo } from "_wailsjs";
import { onBeforeMount, reactive, ref, computed } from "vue";
import moment from "moment";

const form = reactive({
    body: "",
    priority: "",
    datetime: "",
});

const todos = ref([]);
const selectedStatus = ref("all");
const todoToDelete = ref(null);

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

async function deleteTodo(id) {
    try {
        await DeleteTodo(id);
        await fetchTodos();
    } catch (error) {
        console.error("Failed to delete todo:", error);
    }
}

async function toggleStatusTodo(id, status) {
    try {
        await ToggleStatusTodo(id, status);
        await fetchTodos();
    } catch (error) {
        console.error("Failed to complete a todo:", error);
    }
}

function formatDate(dateString) {
    return moment(dateString).utcOffset(0).format("D MMM YYYY, hh:mm");
}

const filteredTodos = computed(() => {
    if (selectedStatus.value === "all") return todos.value;
    return todos.value.filter((todo) => todo.status === selectedStatus.value);
});

function confirmDeleteTodo() {
    if (todoToDelete.value) {
        deleteTodo(todoToDelete.value);
        todoToDelete.value = null;
    }
}
</script>

<template>
    <div>
        <div class="flex justify-between items-center">
            <h1 class="flex-1 text-2xl font-semibold text-center">Todo List</h1>

            <Select v-model="selectedStatus">
                <SelectTrigger class="w-[80px] max-w-[80px]">
                    <SelectValue placeholder="Sort" />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectItem value="all"> All </SelectItem>
                        <SelectItem value="in-progress">
                            In-Progress
                        </SelectItem>
                        <SelectItem value="completed"> Completed </SelectItem>
                    </SelectGroup>
                </SelectContent>
            </Select>
        </div>

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

    <div class="w-full grid grid-cols-4 gap-12 mt-6">
        <div
            v-for="todo in filteredTodos"
            :key="todo.id"
            class="border border-gray-300 rounded-md px-4 py-2.5 min-w-64 max-w-64 max-h-[400px] overflow-scroll"
        >
            <header class="flex justify-between items-center">
                <p
                    class="font-semibold text-2xl"
                    :class="
                        todo.status == 'completed'
                            ? 'line-through decoration-from-font'
                            : ''
                    "
                >
                    {{ todo.body }}
                </p>

                <div class="flex gap-x-2">
                    <AlertDialog>
                        <AlertDialogTrigger as-child>
                            <TrashIcon
                                as="button"
                                class="size-5 text-black hover:text-red-600 transition-all"
                                @click="todoToDelete = todo.id"
                            />
                        </AlertDialogTrigger>
                        <AlertDialogContent>
                            <AlertDialogHeader>
                                <AlertDialogTitle>
                                    Are you absolutely sure?
                                </AlertDialogTitle>
                                <AlertDialogDescription>
                                    This action cannot be undone. This will
                                    permanently delete this task.
                                </AlertDialogDescription>
                            </AlertDialogHeader>
                            <AlertDialogFooter>
                                <AlertDialogCancel>Cancel</AlertDialogCancel>
                                <AlertDialogAction @click="confirmDeleteTodo">
                                    Continue
                                </AlertDialogAction>
                            </AlertDialogFooter>
                        </AlertDialogContent>
                    </AlertDialog>
                    <CheckIcon
                        as="button"
                        class="size-5 text-black hover:text-green-500 transition-all"
                        @click="toggleStatusTodo(todo.id, todo.status)"
                    />
                </div>
            </header>

            <main class="mt-3">
                <p>
                    Status:
                    <span>{{ todo.status }}</span>
                </p>

                <p>
                    Priority:
                    <span>{{ todo.priority }}</span>
                </p>
            </main>

            <footer class="mt-4 border-t border-gray-800 pt-2">
                <p class="text-sm italic text-gray-700">
                    Deadline: {{ formatDate(todo.deadline) }}
                </p>
            </footer>
        </div>
    </div>
</template>
