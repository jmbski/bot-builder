import { models } from '@go/models';
import { nanoid } from 'nanoid';
import { TabDef } from './tab-def.type';
import { Widget } from './widget.type';

export class UILayout extends models.LayoutTemplate {
    public tabs: TabDef[] = [];
    public widgets: Widget[] = [];

    constructor(init?: Partial<UILayout> | Partial<models.LayoutTemplate>) {
        super(init);
        if (init) {
            Object.assign(this, init);
        }

        if (this.tabs) {
            const newTabs: TabDef[] = [];
            this.tabs.forEach(tab => {
                newTabs.push(new TabDef(tab));
            });
            this.tabs = newTabs;
        }

        if (this.widgets) {
            const newWidgets: Widget[] = [];
            this.widgets.forEach(widget => {
                newWidgets.push(new Widget(widget));
            });
            this.widgets = newWidgets;
        }

        this.uuid ??= nanoid();
    }
}