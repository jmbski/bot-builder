import { models } from '@go/models';
import { type PropType, type Ref } from 'vue';
import { TabDefType, WidgetDefType } from './auto-enums';
import Utils from '@/logic/utils';
import { WidgetTypeMap } from '@/services';
import { nanoid } from 'nanoid';
import type { Widget } from './go';

export type Orderable = { order: number }
export type OrderableArray = Orderable[] | Ref<Orderable[]> | Ref<Orderable>[];

export type GoWidget = models.WidgetDef;
export type GoTab = models.TabDef;
export type GoLayout = models.LayoutTemplate;

export type GoWidgetRef = Ref<GoWidget>
export type GoTabRef = Ref<GoTab>
export type GoLayoutRef = Ref<GoLayout>

export type GoWidgetRefArray = Ref<GoWidget[]>
export type GoTabRefArray = Ref<GoTab[]>
export type GoLayoutRefArray = Ref<GoLayout[]>

export type FormValues = Record<string, Ref>

export type WidgetMap = Record<string, Widget[]>
/**
 * Function that generates a union type of string literals.
 */
export function stringLiterals<T extends string>(...args: T[]): T[] {
    return args;
}

/**
 * Union type built from a list of string literals.
 */
export type UnionTypeOf<T extends ReadonlyArray<unknown>> = T extends ReadonlyArray<infer ElementType>
    ? ElementType
    : never;

/** Change any to unknown */
export type WeakObject = Record<string, any>;
export type TypeOfValues = 'string' | 'number' | 'boolean' | 'undefined' | 'object' | 'function' | 'symbol' | 'bigint';


export type MetaData = models.CommonProps
export type MetaFields = keyof MetaData

export type ModifyPartial<T, K extends keyof T, R> = Partial<Omit<T, K>> & {
    [P in K]?: R;
};
export type ModdableType<T, M> = Omit<T, Extract<keyof T, keyof M>> & M;

export type TypedWidget = ModifyPartial<models.WidgetDef, 'type', WidgetDefType>
export type TypedTab = ModifyPartial<models.TabDef, 'type', TabDefType>




export const expandedTypes = stringLiterals(
    'string',
    'null',
    'array',
    'date',
    'regexp',
    'map',
    'set',
    'weakmap',
    'weakset',
    'object',
    'function',
    'symbol',
    'bigint',
    'undefined',
    'number',
    'boolean',
)

export type ExpandedTypes = UnionTypeOf<typeof expandedTypes>;