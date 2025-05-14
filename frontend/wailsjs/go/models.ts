export namespace main {
	
	export class CharacterCard {
	    name: string;
	    avatar: string;
	    notes: string;
	
	    static createFrom(source: any = {}) {
	        return new CharacterCard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.avatar = source["avatar"];
	        this.notes = source["notes"];
	    }
	}

}

export namespace models {
	
	export class CommonProps {
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	
	    static createFrom(source: any = {}) {
	        return new CommonProps(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	    }
	}
	export class AppStateChange {
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    editor_session_id?: string;
	    new_value: any;
	    old_value: any;
	    property_path?: string[];
	    tab_id?: string;
	    widget_id?: string;
	
	    static createFrom(source: any = {}) {
	        return new AppStateChange(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.editor_session_id = source["editor_session_id"];
	        this.new_value = source["new_value"];
	        this.old_value = source["old_value"];
	        this.property_path = source["property_path"];
	        this.tab_id = source["tab_id"];
	        this.widget_id = source["widget_id"];
	    }
	}
	export class VisibleIf {
	    equals: any;
	    field?: string;
	
	    static createFrom(source: any = {}) {
	        return new VisibleIf(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.equals = source["equals"];
	        this.field = source["field"];
	    }
	}
	export class Validation {
	    maxLength?: number;
	    minLength?: number;
	    pattern?: string;
	
	    static createFrom(source: any = {}) {
	        return new Validation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.maxLength = source["maxLength"];
	        this.minLength = source["minLength"];
	        this.pattern = source["pattern"];
	    }
	}
	export class Formatting {
	    as?: string;
	
	    static createFrom(source: any = {}) {
	        return new Formatting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.as = source["as"];
	    }
	}
	export class WidgetDef {
	    type: string;
	    label: string;
	    field: string;
	    tab: string;
	    order: number;
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    default: any;
	    formatting?: Formatting;
	    read_only?: boolean;
	    required?: boolean;
	    tooltip?: string;
	    undoable?: boolean;
	    validation?: Validation;
	    visibleIf?: VisibleIf;
	
	    static createFrom(source: any = {}) {
	        return new WidgetDef(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.label = source["label"];
	        this.field = source["field"];
	        this.tab = source["tab"];
	        this.order = source["order"];
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.default = source["default"];
	        this.formatting = this.convertValues(source["formatting"], Formatting);
	        this.read_only = source["read_only"];
	        this.required = source["required"];
	        this.tooltip = source["tooltip"];
	        this.undoable = source["undoable"];
	        this.validation = this.convertValues(source["validation"], Validation);
	        this.visibleIf = this.convertValues(source["visibleIf"], VisibleIf);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TabDef {
	    type: string;
	    title: string;
	    order: number;
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    field?: string;
	    format?: string;
	    stringify?: boolean;
	    tooltip?: string;
	
	    static createFrom(source: any = {}) {
	        return new TabDef(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.title = source["title"];
	        this.order = source["order"];
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.field = source["field"];
	        this.format = source["format"];
	        this.stringify = source["stringify"];
	        this.tooltip = source["tooltip"];
	    }
	}
	export class LayoutTemplate {
	    name: string;
	    tabs: TabDef[];
	    widgets: WidgetDef[];
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    theme?: string;
	
	    static createFrom(source: any = {}) {
	        return new LayoutTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.tabs = this.convertValues(source["tabs"], TabDef);
	        this.widgets = this.convertValues(source["widgets"], WidgetDef);
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.theme = source["theme"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class EditorSession {
	    file_path: string;
	    order: number;
	    session_type: string;
	    layout_template: LayoutTemplate;
	    layout_instance: LayoutTemplate;
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    tooltip?: string;
	
	    static createFrom(source: any = {}) {
	        return new EditorSession(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file_path = source["file_path"];
	        this.order = source["order"];
	        this.session_type = source["session_type"];
	        this.layout_template = this.convertValues(source["layout_template"], LayoutTemplate);
	        this.layout_instance = this.convertValues(source["layout_instance"], LayoutTemplate);
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.tooltip = source["tooltip"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AppState {
	    active_sessions: EditorSession[];
	    change_store: AppStateChange[];
	    app_version: string;
	    created_by?: string;
	    created_on: string;
	    description?: string;
	    last_modified_by?: string;
	    last_modified_on: string;
	    schema_version: string;
	    status?: string;
	    tags?: string[];
	    uuid: string;
	    auto_save?: boolean;
	    auto_save_interval?: number;
	    meta_fields?: CommonProps;
	    save_on_edit?: boolean;
	    save_on_exit?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AppState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.active_sessions = this.convertValues(source["active_sessions"], EditorSession);
	        this.change_store = this.convertValues(source["change_store"], AppStateChange);
	        this.app_version = source["app_version"];
	        this.created_by = source["created_by"];
	        this.created_on = source["created_on"];
	        this.description = source["description"];
	        this.last_modified_by = source["last_modified_by"];
	        this.last_modified_on = source["last_modified_on"];
	        this.schema_version = source["schema_version"];
	        this.status = source["status"];
	        this.tags = source["tags"];
	        this.uuid = source["uuid"];
	        this.auto_save = source["auto_save"];
	        this.auto_save_interval = source["auto_save_interval"];
	        this.meta_fields = this.convertValues(source["meta_fields"], CommonProps);
	        this.save_on_edit = source["save_on_edit"];
	        this.save_on_exit = source["save_on_exit"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	
	
	
	

}

