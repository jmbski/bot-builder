import Utils from '@/logic/utils';
import { WidgetTypeMap } from '@/services';
import { models } from '@go/models';
import { nanoid } from 'nanoid';
import type { Ref } from 'vue';
import { WidgetDefType } from '../auto-enums';

export class Widget extends models.WidgetDef {
    // New/Modified properties

    public type: WidgetDefType = WidgetDefType.None;
    public children: Widget[] = [];

    // Defaults

    /* @todo */
    public default: any = '';

    constructor(init?: Partial<Widget> | Partial<models.WidgetDef>, propName?: string, prop?: unknown) {
        super(init)
        if (init) {
            Object.assign(this, init);
        }

        if (propName) {
            this.label = Utils.toLabelCase(propName);
        }

        if (prop != null) {
            this.type = WidgetTypeMap[Utils.typeOf(prop)];
        }
        this.uuid ??= nanoid();
    }
}
export type WidgetRef = Ref<Widget>;