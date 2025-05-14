
package models

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
	

)

func Ptr[T any](v T) *T {
    return &v
}

func getUuid() string {
	id, err := gonanoid.New()
	if err != nil {
		id = ""
	}
	return id
}


func NewAppState() AppState {
	return AppState{
		ActiveSessions: make([]EditorSession, 0),
		ChangeStore: make([]AppStateChange, 0),
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		AutoSave: Ptr(false),
		AutoSaveInterval: Ptr(int64(0)),
		MetaFields: nil,
		SaveOnEdit: Ptr(false),
		SaveOnExit: Ptr(false),
	}
}

func NewEditorSession() EditorSession {
	return EditorSession{
		FilePath: "",
		Order: 0,
		SessionType: "",
		LayoutTemplate: LayoutTemplate{},
		LayoutInstance: LayoutTemplate{},
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		Tooltip: Ptr(""),
	}
}

func NewLayoutTemplate() LayoutTemplate {
	return LayoutTemplate{
		Name: "",
		Tabs: make([]TabDef, 0),
		Widgets: make([]WidgetDef, 0),
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		Theme: Ptr(""),
	}
}

func NewTabDef() TabDef {
	return TabDef{
		Type: "",
		Title: "",
		Order: 0,
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		Field: Ptr(""),
		Format: Ptr(""),
		Stringify: Ptr(false),
		Tooltip: Ptr(""),
	}
}

func NewWidgetDef() WidgetDef {
	return WidgetDef{
		Type: "",
		Label: "",
		Field: "",
		Tab: "",
		Order: 0,
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		Default: nil,
		Formatting: nil,
		ReadOnly: Ptr(false),
		Required: Ptr(false),
		Tooltip: Ptr(""),
		Undoable: Ptr(false),
		Validation: nil,
		VisibleIf: nil,
	}
}

func NewFormatting() Formatting {
	return Formatting{
		As: Ptr(""),
	}
}

func NewValidation() Validation {
	return Validation{
		MaxLength: Ptr(int64(0)),
		MinLength: Ptr(int64(0)),
		Pattern: Ptr(""),
	}
}

func NewVisibleIf() VisibleIf {
	return VisibleIf{
		Equals: nil,
		Field: Ptr(""),
	}
}

func NewAppStateChange() AppStateChange {
	return AppStateChange{
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
		EditorSessionID: Ptr(""),
		NewValue: nil,
		OldValue: nil,
		PropertyPath: make([]string, 0),
		TabID: Ptr(""),
		WidgetID: Ptr(""),
	}
}

func NewCommonProps() CommonProps {
	return CommonProps{
		AppVersion: "0.0.1",
		CreatedBy: Ptr("bot-builder"),
		CreatedOn: time.Now().Format(time.RFC3339),
		Description: Ptr(""),
		LastModifiedBy: Ptr("bot-builder"),
		LastModifiedOn: time.Now().Format(time.RFC3339),
		SchemaVersion: "0.0.1",
		Status: Ptr(Active),
		Tags: make([]string, 0),
		UUID: getUuid(),
	}
}
