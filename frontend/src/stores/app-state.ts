// stores/appState.ts
import { defineStore } from 'pinia';
import { models } from '@go/models';
import { LoadAppState, SaveAppState } from '@go';
import { AppState } from '@/types';



export const useAppState = defineStore('appState', {
    state: () => (new AppState()),
    actions: {
        markDirty() {
            this.isDirty = true;
        },
        saveState() {
            /** Assuming `window.backend.SaveState` is your Wails call */

            SaveAppState(new models.AppState(this.$state)).then(_ => {

                this.isDirty = false;
            })
        },
        loadState(data: models.AppState) {
            Object.assign(this.$state, data);
        }
    }
});
