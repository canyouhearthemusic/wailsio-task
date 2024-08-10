<script setup>
import { PencilIcon, TrashIcon, CheckIcon } from "@heroicons/vue/24/outline";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
</script>

<template>
    <div>
        <h1 class="text-2xl font-semibold text-center">Todo List</h1>

        <div class="flex w-full max-w-sm items-center gap-1.5 mt-5">
            <Input id="email" type="email" placeholder="What to do?" />
            <Button type="submit"> Create </Button>
        </div>
    </div>

    <div class="w-full grid grid-cols-4 gap-8 mt-6">
        <div
            class="border border-gray-300 rounded-md px-4 py-2.5 w-64 max-w-64 max-h-[400px] overflow-scroll"
        >
            <header class="flex justify-between items-center">
                <p class="font-semibold text-2xl">Status</p>

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
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Repudiandae, aliquid sit! Illum obcaecati blanditiis eaque
                excepturi ea officiis, porro maxime. Maxime, beatae inventore ad
                iure incidunt provident deleniti doloribus aperiam.
            </main>

            <footer class="mt-4 border-t border-gray-800 pt-2">
                <p class="text-sm italic text-gray-700">
                    Created at: 2022-01-01
                </p>
                <p class="text-sm italic text-gray-700">
                    Last changed at: 2022-01-01
                </p>
            </footer>
        </div>
    </div>
</template>
