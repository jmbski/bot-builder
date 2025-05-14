<script setup lang="ts">
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';

import { LoadTemplateSchema } from '@go';
import { onMounted, ref, type Ref } from 'vue';
import { models } from '@go/models';
import DynamicForm from './DynamicForm.vue';
import { type UILayout, type WidgetMap } from '@apptypes';
import Utils from '@/logic/utils';
import type { FormSubmitEvent } from '@primevue/forms';

const template: any = {};//ref<UILayout>(new models.LayoutTemplate())
const widgetMap: WidgetMap = {};
const refMap: Record<string, Ref> = {};

function onSubmit(event: FormSubmitEvent) {
    console.log('tab display', event)
}


/* function loadTemplate(templateName?: string) {
    templateName ??= "test.json";
    console.log('\nTEST\n')
    LoadTemplateSchema(templateName).then(result => {
        Object.assign(template.value, result)

        Utils.sortByOrder(template.value.tabs)

        template.value.tabs.forEach(tab => {
            widgetMap[tab.uuid] = []
        });

        template.value.widgets.forEach(widget => {
            if (!widgetMap[widget.tab]?.includes(widget)) {
                widgetMap[widget.tab].push(widget);
                refMap[widget.uuid] = ref(widget.default)

            }


        });

        Object.values(widgetMap).forEach(widgets => {
            Utils.sortByOrder(widgets)
        })
        console.log('widgetMap', widgetMap)
    })
} */

/* onMounted(() => {
    loadTemplate();
}) */
</script>

<template>
    <div class="card">
        <Tabs value="0">
            <TabList>
                <Tab v-for="tab in template.tabs" :key="tab.title" :value="tab.order">{{ tab.title }}</Tab>
            </TabList>
            <TabPanels>
                <TabPanel v-for="tab in template.tabs" :key="tab.field" :value="tab.order">
                    <DynamicForm :widgets="widgetMap[tab.uuid]" :ref-map="refMap" @submit=""></DynamicForm>
                </TabPanel>
            </TabPanels>
        </Tabs>
    </div>
</template>

<style scoped>
.tab-display {
    /* styles for TabDisplay */
}
</style>
