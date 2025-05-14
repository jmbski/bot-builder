import type { models } from '@go/models';
import { defineProps, defineEmits, ref } from 'vue'

export function useBasicFieldProps() {
    const props = defineProps<{
        widget: models.WidgetDef
    }>()
    const value = ref(props.widget.default);

    return { props, value }
}
