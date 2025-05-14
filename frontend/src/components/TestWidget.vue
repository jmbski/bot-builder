<script setup lang="ts">
import { models } from '@go/models';
import { type GoWidgetRefArray } from '@apptypes';
import { computed, onMounted, reactive, ref } from 'vue'
import { useToast } from "primevue/usetoast";
import { z } from 'zod';
import TextInput from './widgets/TextInput.vue'
import { InputText, Button, Message } from 'primevue';
import { zodResolver } from '@primevue/forms/resolvers/zod'
import { Form, type FormSubmitEvent } from '@primevue/forms';

const props = defineProps<{
    widgets: GoWidgetRefArray
}>()


const initialValues = reactive({
    username: ''
});



/* const resolver = ref(zodResolver(
    z.object({
        username: z.string().min(1, { message: 'Username is required via Zod.' })
    })
)); */

/* const onFormSubmit = (event: FormSubmitEvent) => {
    if (event.valid) {
        toast.add({ severity: 'success', summary: 'Form is submitted.', life: 3000 });
    }
} */
/* const test = ref("Not loaded"); */

/* onMounted(() => {
    test.value = "Loaded";
}) */
</script>

<template>
    <div class="test-widget">
        This is a testd
    </div>
</template>

<style scoped>
.test-widget {
    /* styles for TestWidget */
}
</style>
