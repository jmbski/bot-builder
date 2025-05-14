<script setup lang="ts">
import type { FormValues, GoWidget, Widget } from '@/types';
import DynamicForm from '../layout/DynamicForm.vue';
import { inject, onMounted, type Ref } from 'vue';
import type { DynamicDialogInstance } from 'primevue/dynamicdialogoptions';
import FormService from '@/services/form.service';

const dialogRef: Ref<DynamicDialogInstance> | undefined = inject('dialogRef');


//let widgets: Ref<SimpleWidget[]> = ref<SimpleWidget[]>([]);

const widgets: Widget[] = dialogRef?.value.data.widgets ?? [];

const emit = defineEmits(['cancel', 'save'])
function cancelDialog() {
    emit('cancel', { user: 'primetime' });
}

function saveDialog() {
    emit('save', { user: 'primetime' });

}

function onSubmit(formRefs: FormValues) {
    const formValues = FormService.parseFormValues(formRefs);
    console.log('values:', formValues);
}

function onCancel(event?: unknown) {
    dialogRef?.value.close()
}

function typeSafe<T>(obj: unknown): obj is T {
    return true
}

onMounted(() => {
    /* const dataWidgets = dialogRef?.value.data.widgets;
    if (typeSafe<SimpleWidget[]>(dataWidgets)) {

    } */
})
</script>

<template>
    <div class="form-dialog">
        <DynamicForm v-if="widgets" :widgets="widgets" @submit="onSubmit($event)" @cancel="onCancel($event)"/>
        <!-- <div class="mt-2">
            <ButtonGroup>
                <Button label="Save" icon="pi pi-check" @click="saveDialog()"/>
                <Button label="Cancel" icon="pi pi-times" @click="cancelDialog()"/>
            </ButtonGroup>
        </div> -->
    </div>
</template>

<style scoped>
.form-dialog {
    /* styles for FormDialog */
}
</style>
