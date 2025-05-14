<script setup lang="ts">

import InputText from 'primevue/inputtext';
import { ref, type Ref } from 'vue';
import { TypeGuards, type FormValues, type GoWidget, type Widget } from '@/types';

const props = defineProps<{
    widget: Widget,
    formValues: FormValues,
}>()
//const value = ref(props.widget.default);

function getStringRef(): Ref<string> {
    const val = props.formValues[props.widget.field ?? ''];
    if (TypeGuards.isStringRef(val)) {
        return val
    }
    else if (TypeGuards.isString(val)) {
        return ref(val)
    }
    return ref(props.widget.default ?? '')
}
const value = getStringRef()


</script>

<template>
    <label class="col-span-2" for="username">{{ props.widget?.label }}</label>
    <InputText class="col-span-8" type="text" v-model="value" />
</template>

<style scoped></style>
