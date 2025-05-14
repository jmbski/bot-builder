<template>
    <div class="layout-dialog">

    </div>
</template>

<script setup lang="ts">
import type { FormValues, GoWidget } from '@/types';
import DynamicForm from '../layout/DynamicForm.vue';
import { inject, onMounted, type Ref } from 'vue';
import type { DynamicDialogInstance } from 'primevue/dynamicdialogoptions';
import FormService from '@/services/form.service';

const dialogRef: Ref<DynamicDialogInstance> | undefined = inject('dialogRef');

const formValues: FormValues = {

}

function onSubmit(formRefs: FormValues) {
}

function onCancel(event?: unknown) {
    dialogRef?.value.close()
}
</script>

<style scoped>
.layout-dialog {
    /* styles for FormDialog */
}
</style>
