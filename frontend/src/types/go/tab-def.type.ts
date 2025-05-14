import { models } from '@go/models';
import { nanoid } from 'nanoid';
import { TabDefType } from '../auto-enums';
import type { Widget } from './widget.type';
import type { Ref } from 'vue';

export class TabDef extends models.TabDef {
    public type: TabDefType = TabDefType.Compound;
    public children: Widget[] = [];

    constructor(init?: Partial<TabDef> | Partial<models.TabDef>) {
        super(init)
        if (init) {
            Object.assign(this, init);
        }
        this.uuid ??= nanoid();
    }
}
export type TabDefRef = Ref<TabDef>