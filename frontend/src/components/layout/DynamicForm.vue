<script setup lang="ts">

import { models } from '@go/models';
import { WidgetDefType, type FormValues, type GoWidget, type GoWidgetRefArray, type Widget } from '@apptypes';
import { computed, onMounted, reactive, ref, type Ref } from 'vue'
import { useToast } from "primevue/usetoast";
import { z } from 'zod';
import TextInput from '../widgets/TextInput.vue'
import { InputText, Button, Message, ButtonGroup } from 'primevue';
import { zodResolver } from '@primevue/forms/resolvers/zod'
import { Form, type FormSubmitEvent } from '@primevue/forms';

import Toast from 'primevue/toast'
import { WidgetTypeDefaults } from '@/services';
import NumberWidget from '../widgets/NumberWidget.vue';
import ToggleWidget from '../widgets/ToggleWidget.vue';
import FormService from '@/services/form.service';


const props = defineProps<{
    widgets: Widget[],
    refMap?: Record<string, Ref>,
    formValues?: FormValues,
    showButtons?: boolean,
    onSubmit?: (formValues: FormValues) => void,
    onCancel?: (event?: unknown) => void,
}>()
const toast = useToast();
const showButtons: boolean = props.showButtons ?? true;
const widgets = ref(props.widgets)

const _initialValues = FormService.getWidgetInitVals(props.widgets)

const formValues = props.formValues ?? FormService.getFormValues(_initialValues)
const initialValues = reactive(_initialValues);

const resolver = ref(zodResolver(
    /** @todo implement custom validation handling */
    z.object({
        //username: z.string().min(1, { message: 'Username is required via Zod.' })
    })
));

const onSubmit = props.onSubmit;
const onCancel = props.onCancel;

/* const onFormSubmit = (event: FormSubmitEvent) => {
    const { onSubmit } = props;
    console.log('onSubmit', onSubmit)
    if (onSubmit) {
        return onSubmit(event);
    }
    if (event.valid) {
        toast.add({ severity: 'success', summary: 'Form is submitted.', life: 3000 });
    }
} */



</script>

<template>
    <div class="w-full">
          <Toast />

          <Form v-slot="$form" :initialValues :resolver="resolver"  class="w-full grid grid-cols-1 gap-6">
              <template v-for="widget in widgets">
                  <div class="grid grid-cols-10 col-auto items-center border-1 rounded min-h-11">
                    <template v-if="widget.type === 'text'">
                        <TextInput :widget="widget" :form-values="formValues"/>
                    </template>
                    <template v-if="widget.type === 'number'">
                        <NumberWidget :widget="widget" :form-values="formValues"/>
                    </template>
                    <template v-if="widget.type === 'toggle'">
                        <ToggleWidget :widget="widget" :form-values="formValues"/>
                    </template>
                  </div>
              </template>
              <ButtonGroup class="w-full gap-1" v-if="showButtons">
                <Button v-if="onSubmit" class="w-1/2" label="Submit" type="button" @click="onSubmit(formValues)"/>
                <Button v-if="onCancel" class="w-1/2" label="Cancel" type="button" @click="onCancel($event)"/>
              </ButtonGroup>

          </Form>
      </div>
</template>

<style scoped>
.widget-gen {
    /* styles for DynamicForm */
}
</style>
