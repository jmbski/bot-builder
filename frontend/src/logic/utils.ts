import { stringLiterals, type ExpandedTypes, type OrderableArray, type UnionTypeOf } from '@/types';
import TypeGuards from '@/types/type-guards';
import type { Ref } from 'vue';

export default class Utils {
    public static sortByOrder(items: unknown) {
        if (TypeGuards.isOrderableRefArrayRef(items)) {
            items.value.sort((a, b) => {
                const aOrder = a.value.order;
                const bOrder = b.value.order;

                return aOrder < bOrder ? -1 : aOrder == bOrder ? 0 : 1;
            })
        }
        else if (TypeGuards.isOrderableRefArray(items)) {
            items.sort((a, b) => {
                const aOrder = a.value.order;
                const bOrder = b.value.order;

                return aOrder < bOrder ? -1 : aOrder == bOrder ? 0 : 1;
            })
        }
        else if (TypeGuards.isOrderableArray(items)) {
            items.sort((a, b) => {
                return a.order < b.order ? -1 : a.order == b.order ? 0 : 1;
            })

        }
    }

    public static toLabelCase(str: string): string {
        // Remove all non-word characters (like hyphens and underscores) and replace camel cased words
        const labelCase = str.replace(/[\W_]+(.)/g, ' $1').replace(/([a-z\d])([A-Z])/g, '$1 $2');

        // Split the string into words
        const words = labelCase.split(' ');

        // Capitalize the first letter of each word and join back into a string
        const capitalizedWords = words.map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');

        return capitalizedWords;
    }

    /**
     * Retrieves the type of the provided value, similar to the built-in typeOf
     * but with more types
     * 
     * @param value some value to retrieve the type of
     */
    public static typeOf(value: unknown): ExpandedTypes {
        if (value === null) {
            return 'null';
        }
        if (Array.isArray(value)) {
            return 'array';
        }
        if (value instanceof Date) {
            return 'date';
        }
        if (value instanceof RegExp) {
            return 'regexp';
        }
        if (value instanceof Map) {
            return 'map';
        }
        if (value instanceof Set) {
            return 'set';
        }
        if (value instanceof WeakMap) {
            return 'weakmap';
        }
        if (value instanceof WeakSet) {
            return 'weakset';
        }
        return typeof value;
    }
}

//export type UtilsTypeOfReturnTypes = ;
