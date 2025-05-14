<script setup lang="ts">
import { onMounted, ref } from 'vue';
import HelloWorld from './components/HelloWorld.vue'
import { Button } from 'primevue';
import TheWelcome from './components/TheWelcome.vue'
import TabDisplay from './components/layout/TabDisplay.vue';
import DynamicForm from './components/layout/DynamicForm.vue';
import TextInput from './components/widgets/TextInput.vue';
import TestWidget from './components/TestWidget.vue';
import TopNav from './components/layout/TopNav.vue';
import { InputText } from 'primevue';

import { AppState, type GoWidgetRef, type GoWidgetRefArray } from './types';
import { models } from '@go/models';
import { useAppState } from './stores/app-state';
import { LoadAppState } from '@go';
import FormService from './services/form.service';

const widgets: GoWidgetRefArray = ref<models.WidgetDef[]>([])
const appState = useAppState()


onMounted(() => {
    LoadAppState().then(result => {
        appState.loadState(new AppState(result))
    })

})

function saveState() {
    appState.auto_save = false;
    appState.saveState();
}
</script>

<template>
    <main>
        <TopNav/>
        <Button label="Test" @click="saveState()"/>
        <TabDisplay/>
    </main>

    <!-- <div class="rounded-2xl border-1 m-2 p-4 shadow-md">
        <div id="a-form" class="test-form min-h-10 grid grid-cols-1 w-full ">
            <label class="mb-1">Label 1</label>
            <InputText/>
        </div>
    </div> -->
</template>

<style scoped>
.test-form {
    background-color: grey;
}
</style>
