import { isRef, type Ref } from 'vue';
import { expandedTypes, type ExpandedTypes, type Orderable, type WeakObject } from './types';

export default class TypeGuards {
    public static isString(value: unknown): value is string {
        return typeof value === 'string';
    }

    public static isStringRef(value: unknown): value is Ref<string> {
        return isRef(value) && TypeGuards.isString(value.value)
    }

    public static isNumber(value: unknown): value is number {
        return typeof value === 'number';
    }

    public static isNumberRef(value: unknown): value is Ref<number> {
        return isRef(value) && TypeGuards.isNumber(value.value)
    }

    public static isBoolean(value: unknown): value is boolean {
        return typeof value === 'boolean';
    }

    public static isBooleanRef(value: unknown): value is Ref<boolean> {
        return isRef(value) && TypeGuards.isBoolean(value.value)
    }

    public static isObject(value: unknown): value is object {
        return typeof value === 'object' && value !== null;
    }

    public static isObjectRef(value: unknown): value is Ref<object> {
        return isRef(value) && TypeGuards.isObject(value.value)
    }

    public static isFunction(value: unknown): value is (...args: unknown[]) => unknown {
        return typeof value === 'function';
    }

    public static isOrderable(value: unknown): value is Orderable {
        return (
            TypeGuards.isObject(value) &&
            Object.hasOwn(value, 'order') &&
            TypeGuards.isNumber((value as Orderable).order)
        );
    }

    public static isOrderableArray(value: unknown): value is Orderable[] {
        return (
            Array.isArray(value) &&
            value.every((item) => TypeGuards.isOrderable(item))
        );
    }

    public static isOrderableArrayRef(value: unknown): value is Ref<Orderable[]> {
        return (
            isRef(value) && TypeGuards.isOrderableArray(value.value)
        );
    }

    public static isOrderableRef(value: unknown): value is Ref<Orderable> {
        return isRef(value) && TypeGuards.isOrderable(value.value)
    }

    public static isOrderableRefArray(value: unknown): value is Ref<Orderable>[] {
        return Array.isArray(value) && value.every(i => TypeGuards.isOrderableRef(i))
    }

    public static isOrderableRefArrayRef(value: unknown): value is Ref<Ref<Orderable>[]> {
        return isRef(value) && TypeGuards.isOrderableArrayRef(value.value)
    }

    /**
     * Checks if the input is a string indexed object.
     * 
     * @param input- The value to check if it is a string indexed object
     * @returns true if the input is a string indexed object, false otherwise
     */
    public static isWeakObject(input: unknown): input is WeakObject {
        return typeof input === 'object' && input != null && !Array.isArray(input) && Object.keys(input).every(key => typeof key === 'string');
    }

    /**
     * Checks if the input is an object that is not an array.
     * While this will typically result in the same value as isWeakObject, it is a quicker check,
     * but also less strict.
     * 
     * @param input - The value to check if it is an object that is not an array
     * @returns true if the input is an object that is not an array, false otherwise
     */
    public static isWeakObjectQuick(input: unknown): input is WeakObject {
        return typeof input === 'object' && input != null && !Array.isArray(input);
    }



    public static isWeakObjectRef(value: unknown): value is Ref<WeakObject> {
        return isRef(value) && TypeGuards.isWeakObject(value.value)
    }

    public static isExpandedTypeOf(value: unknown): value is ExpandedTypes {
        return TypeGuards.isString(value) && expandedTypes.includes(value as ExpandedTypes)
    }
}