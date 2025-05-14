import Utils from '@/logic/utils';
import { TypeGuards, Widget, WidgetDefType, type FormValues, type GoWidget, type WeakObject } from '@/types';
import { WidgetTypeMap, WidgetTypeDefaults } from '@/consts/widget-maps';
import { nanoid } from 'nanoid';
import { ref, type Ref } from 'vue';

const READ_ONLY_FIELDS: string[] = [
    'uuid',
    'created_on',
    'last_modified_on'
]

export default class FormService {

    public static getWidgetType(value: unknown): WidgetDefType {
        return WidgetTypeMap[Utils.typeOf(value)];
    }

    public static getWidgetInitVals(widgets: Widget[]): WeakObject {
        const initialValues: Record<string, any> = {}

        widgets.forEach(widget => {
            const { field, type } = widget;
            if (!field) {
                return
            }
            const defaultVal = WidgetTypeDefaults[type ?? WidgetDefType.None]
            if (defaultVal != null) {
                initialValues[field] = defaultVal;
            }
        })

        return initialValues;
    }

    public static getFormValues(initialValues: Record<string, any>): FormValues {
        const formValues: Record<string, Ref> = {}
        Object.keys(initialValues).forEach(key => {
            const value = initialValues[key];
            switch (FormService.getWidgetType(value)) {
                case WidgetDefType.Container:
                    formValues[key] = ref(FormService.getFormValues(value));
                default:
                    formValues[key] = ref(value)
            }
        })
        return formValues
    }

    public static parseFormValues(obj: Record<string, Ref>): WeakObject {
        const values: WeakObject = {}
        Object.keys(obj).forEach(key => {
            values[key] = obj[key].value;
        })
        return values;
    }

    public static objToFormSimple(obj: unknown): Widget[] | undefined {
        if (!TypeGuards.isWeakObject(obj)) {
            return undefined;
        }

        const widgets: Widget[] = [];

        Object.keys(obj).forEach((key, index) => {
            const prop = obj[key];
            widgets.push(new Widget({ order: index }, key, prop))
        })

        return widgets;
    }
}