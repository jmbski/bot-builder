import { models } from '@go/models';
import { nanoid } from 'nanoid';

export class AppStateChange extends models.AppStateChange {
    constructor(init?: Partial<AppStateChange> | Partial<models.AppStateChange>) {
        super(init);
        if (init) {
            Object.assign(this, init);
        }
        this.uuid ??= nanoid();
    }
}