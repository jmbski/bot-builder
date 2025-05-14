<script setup lang="ts">

import { ref, type Ref } from 'vue';
import { TypeGuards, type FormValues, type GoWidget, type Widget } from '@/types';
import { InputNumber } from 'primevue';

const props = defineProps<{
    widget: Widget,
    formValues: FormValues,
}>()
//const value = ref(props.widget.default);

const defaultValue = 0;

function getStringRef(): Ref<number> {
    const val = props.formValues[props.widget.field ?? defaultValue];
    if (TypeGuards.isNumberRef(val)) {
        return val
    }
    else if (TypeGuards.isNumber(val)) {
        return ref(val)
    }
    return ref(props.widget.default ?? defaultValue)
}
const value = getStringRef()


</script>

<template>
    <label class="col-span-2" for="number-widget">{{ props.widget.label }}</label>
    <InputNumber id="number-widget" class="col-span-8" v-model="value"/>
</template>

<style scoped></style>
