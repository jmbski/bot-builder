import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia';

import App from './App.vue'

import PrimeVue from 'primevue/config'
import Aura from '@primeuix/themes/aura'
import ToastService from 'primevue/toastservice';
import DialogService from 'primevue/dialogservice';


const app = createApp(App)
const pinia = createPinia();


app.config.errorHandler = (err, instance, info) => {
    // Handle the error, e.g., log it, display a user-friendly message, or report it to a service
    console.error('An error occurred:', err);
    console.warn('Vue instance:', instance);
    console.warn('Error info:', info);

    /** @todo implement better error logging, probably through Go */
};
app.use(pinia)
app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
})
app.use(ToastService)
app.use(DialogService)
app.mount('#app')
