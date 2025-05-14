<script setup lang="ts">

import { ref, type Ref } from 'vue';
import { TypeGuards, type FormValues, type GoWidget, type Widget } from '@/types';
import { ToggleSwitch } from 'primevue';

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
    <label class="col-span-2" for="toggle-widget">{{ props.widget?.label }}</label>
    <div class="col-span-8 flex justify-center align-center">
        <ToggleSwitch id="toggle-widget" class="" v-model="value" />
    </div>
</template>


<style scoped></style>
