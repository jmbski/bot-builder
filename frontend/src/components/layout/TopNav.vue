<script setup lang='ts'>
import { useAppState } from '@/stores/app-state';
import { LoadAppState, LoadTemplateSchema, ResetAppState } from '@go';
import { models } from '@go/models';
import Menubar from 'primevue/menubar';
import type { MenuItem } from 'primevue/menuitem';
import { defineAsyncComponent, ref, type Ref } from 'vue';
import { useDialog } from 'primevue/usedialog';
import FormService from '@/services/form.service';
import { Button, DynamicDialog } from 'primevue';
import { Widget } from '@/types';


const SimpleFormDialog = defineAsyncComponent(() => import('@/components/forms/SimpleFormDialog.vue'));
const dialog = useDialog();
const appState = useAppState()

function newWidget() {
    const widgetObj = new Widget(undefined, 'blank', '')
    const widgetForm = FormService.objToFormSimple(widgetObj)
    console.log('form', widgetForm)
    return dialog.open(SimpleFormDialog, {
        props: {
            header: 'New Widget',
            appendTo: 'body',
            closable: true,
            draggable: true,
            maximizable: true,
            style: {
                width: '50vw',
            },
            breakpoints: {
                '960px': '75vw',
                '640px': '90vw'
            },
            modal: true
        },
        data: {
            widgets: widgetForm
        },
        onClose: (opts) => {
            console.log(opts)
        }
    });
}

const model = ref<MenuItem[]>([
    {
        label: 'File',
        items: [
            { label: 'New Character', command: () => { } },
            { label: 'Open Character...', command: () => { } },
            { label: 'Save', command: () => { } },
            { label: 'Save As...', command: () => { } },
            { separator: true },
            { label: 'New Lorebook', command: () => { } },
            { label: 'Open Lorebook...', command: () => { } },
            { label: 'Save Lorebook', command: () => { } },
            { separator: true },
            {
                label: 'Reset State',
                command: () => {
                    ResetAppState().then(() => {
                        LoadAppState().then(result => {
                            appState.loadState(result);
                        }).catch(err => { })
                    })
                }
            },
            { label: 'Exit', command: () => { } }
        ]
    },
    {
        label: 'Edit',
        items: [
            { label: 'Undo', command: () => { } },
            { label: 'Redo', command: () => { } },
            { separator: true },
            { label: 'Cut', command: () => { } },
            { label: 'Copy', command: () => { } },
            { label: 'Paste', command: () => { } }
        ]
    },
    /* {
        label: 'View',
        items: [
            { label: 'Toggle Sidebar', command: () => { } },
            { label: 'Toggle Preview', command: () => { } },
            { label: 'Zoom In', command: () => { } },
            { label: 'Zoom Out', command: () => { } },
            { label: 'Reset Layout', command: () => { } }
        ]
    }, */
    {
        label: 'Template',
        items: [
            { label: 'Load Template...', command: () => { } },
            { label: 'Save Template', command: () => { } },
            { label: 'Clone Template', command: () => { } },
            { separator: true },
            { label: 'Edit Current Template', command: () => { } },
            {
                label: 'Create New Template',
                command: () => {
                    console.log("Template new")
                    const dialogRef = newWidget()
                    console.log(dialogRef)
                }
            }
        ]
    },
    {
        label: 'Tools',
        items: [
            { label: 'Validate Schema', command: () => { } },
            { label: 'Analyze Card Format', command: () => { } },
            { label: 'AI Message Generator', command: () => { } },
            { label: 'Lore Entry Formatter', command: () => { } }
        ]
    },
    {
        label: 'Preferences',
        items: [
            {
                label: 'Theme',
                items: [
                    { label: 'System', command: () => { } },
                    { label: 'Light', command: () => { } },
                    { label: 'Dark', command: () => { } }
                ]
            },
            {
                label: 'Autosave',
                items: [
                    { label: 'Enabled', command: () => { } },
                    { label: 'Interval...', command: () => { } }
                ]
            },
            { label: 'Default Template...', command: () => { } }
        ]
    },
    {
        label: 'Help',
        items: [
            { label: 'User Guide', command: () => { } },
            { label: 'View Logs', command: () => { } },
            { label: 'About', command: () => { } }
        ]
    }
])
</script>

<template>
    <div class="top-nav">
        <Menubar :model='model' />
        <DynamicDialog/>
    </div>
</template>

<style scoped>
.top-nav {
    /* styles for TopNav */
}
</style>
