import { models } from '@go/models';
import { nanoid } from 'nanoid';
import { UILayout } from './layout.type';

export class EditSession extends models.EditorSession {
    public layoutTemplate: UILayout = new UILayout();
    public layoutInstance: UILayout = new UILayout();

    constructor(init?: Partial<EditSession> | Partial<models.LayoutTemplate>) {
        super(init);
        if (init) {
            Object.assign(this, init);
        }

        this.layoutTemplate = new UILayout(this.layout_template);
        this.layoutInstance = new UILayout(this.layout_instance);

        this.uuid ??= nanoid();
    }
}