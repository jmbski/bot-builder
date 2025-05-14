<script setup lang="ts">

import TextArea from 'primevue/textarea';
import { useBasicFieldProps } from './shared';
import type { models } from '@go/models';
import { ref } from 'vue';
import type { Widget } from '@/types';

const props = defineProps<{
    widget: Widget
}>()
const value = ref(props.widget.default);


</script>

<template>
    <label class="col-span-2" for="username">{{ props.widget.label }}</label>
    <TextArea class="col-span-8" v-model="value" />
</template>

<style scoped></style>
