import { models } from '@go/models';
import { nanoid } from 'nanoid';
import type { WeakObject } from '../types';
import { EditSession } from './edit-session.type';

const AppStateKeys = Object.keys(new models.AppState());

export class AppState extends models.AppState {
    public activeSessions: EditSession[] = []
    public changeStore: models.AppStateChange[] = [];
    public isDirty: boolean = false;


    constructor(init?: Partial<AppState>) {
        super(init);
        if (init) {
            Object.assign(this, init);
        }

        this.activeSessions = this.active_sessions?.map(session => {
            return new EditSession(session);
        }) ?? [];
        this.changeStore ??= this.change_store;

        this.uuid ??= nanoid();
    }

    public static getSaveContent(state: WeakObject): WeakObject {
        const content: WeakObject = {};
        AppStateKeys.forEach(key => {
            content[key] = state[key]
        })
        return content;
    }
}