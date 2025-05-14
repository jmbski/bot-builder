import { WidgetDefType, type ExpandedTypes } from '@/types';
import { z } from 'zod';

export const WidgetTypeMap: Record<ExpandedTypes, WidgetDefType> = {
    string: WidgetDefType.Text,
    bigint: WidgetDefType.Number,
    boolean: WidgetDefType.Toggle,
    function: WidgetDefType.None,
    number: WidgetDefType.Number,
    object: WidgetDefType.None,
    symbol: WidgetDefType.Text,
    undefined: WidgetDefType.None,
    array: WidgetDefType.List,
    date: WidgetDefType.Calendar,
    map: WidgetDefType.Container,
    null: WidgetDefType.None,
    regexp: WidgetDefType.Text,
    set: WidgetDefType.List,
    weakmap: WidgetDefType.Container,
    weakset: WidgetDefType.List
}

export const WidgetToZodMap: Record<WidgetDefType, unknown> = {
    calendar: z.date,
    custom: z.string,
    dropdown: z.array,
    image: z.string,
    list: z.array,
    container: z.object,
    none: z.undefined,
    number: z.number,
    text: z.string,
    textarea: z.string,
    toggle: z.boolean,
}

export const WidgetTypeDefaults: Record<WidgetDefType, unknown> = {
    calendar: '',
    custom: '',
    dropdown: [],
    image: '',
    list: [],
    container: {},
    none: undefined,
    number: 0,
    text: '',
    textarea: '',
    toggle: false,
}