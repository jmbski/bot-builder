// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appState, err := UnmarshalAppState(bytes)
//    bytes, err = appState.Marshal()

package models

import "encoding/json"

func UnmarshalAppState(data []byte) (AppState, error) {
	var r AppState
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppState) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Used for tracking the overall state of the app
//
// Common properties for all schema objects
type AppState struct {
	// Open editor session tabs                                                                                  
	ActiveSessions                                                                              []EditorSession  `json:"active_sessions"`
	// Log of recent changes in app                                                                              
	ChangeStore                                                                                 []AppStateChange `json:"change_store"`
	// Version of the app                                                                                        
	AppVersion                                                                                  string           `json:"app_version"`
	// User who created the object                                                                               
	CreatedBy                                                                                   *string          `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle                 
	// "format": "date-time"                                                                                     
	CreatedOn                                                                                   string           `json:"created_on"`
	// Description of the object                                                                                 
	Description                                                                                 *string          `json:"description,omitempty"`
	// User who last modified the object                                                                         
	LastModifiedBy                                                                              *string          `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":                 
	// "date-time"                                                                                               
	LastModifiedOn                                                                              string           `json:"last_modified_on"`
	// Version of the schema                                                                                     
	SchemaVersion                                                                               string           `json:"schema_version"`
	// Lifecycle state of the object                                                                             
	Status                                                                                      *Status          `json:"status,omitempty"`
	// Optional user-defined tags                                                                                
	Tags                                                                                        []string         `json:"tags,omitempty"`
	// Unique identifier for the object                                                                          
	UUID                                                                                        string           `json:"uuid"`
	// Auto save the current session                                                                             
	AutoSave                                                                                    *bool            `json:"auto_save,omitempty"`
	// Interval in seconds for auto save                                                                         
	AutoSaveInterval                                                                            *int64           `json:"auto_save_interval,omitempty"`
	MetaFields                                                                                  *CommonProps     `json:"meta_fields,omitempty"`
	// Save the current session on edit                                                                          
	SaveOnEdit                                                                                  *bool            `json:"save_on_edit,omitempty"`
	// Save the current session on exit                                                                          
	SaveOnExit                                                                                  *bool            `json:"save_on_exit,omitempty"`
}

// Defines a workspace for a single file
//
// Common properties for all schema objects
type EditorSession struct {
	// Path to the in-work file for the session                                                                
	FilePath                                                                                    string         `json:"file_path"`
	// Tab order for the editor session panel                                                                  
	Order                                                                                       float64        `json:"order"`
	// What type of editor session is it (i.e, character, world lore, etc..)                                   
	SessionType                                                                                 string         `json:"session_type"`
	LayoutTemplate                                                                              LayoutTemplate `json:"layout_template"`
	LayoutInstance                                                                              LayoutTemplate `json:"layout_instance"`
	// Version of the app                                                                                      
	AppVersion                                                                                  string         `json:"app_version"`
	// User who created the object                                                                             
	CreatedBy                                                                                   *string        `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle               
	// "format": "date-time"                                                                                   
	CreatedOn                                                                                   string         `json:"created_on"`
	// Description of the object                                                                               
	Description                                                                                 *string        `json:"description,omitempty"`
	// User who last modified the object                                                                       
	LastModifiedBy                                                                              *string        `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":               
	// "date-time"                                                                                             
	LastModifiedOn                                                                              string         `json:"last_modified_on"`
	// Version of the schema                                                                                   
	SchemaVersion                                                                               string         `json:"schema_version"`
	// Lifecycle state of the object                                                                           
	Status                                                                                      *Status        `json:"status,omitempty"`
	// Optional user-defined tags                                                                              
	Tags                                                                                        []string       `json:"tags,omitempty"`
	// Unique identifier for the object                                                                        
	UUID                                                                                        string         `json:"uuid"`
	Tooltip                                                                                     *string        `json:"tooltip,omitempty"`
}

// General template for the structure of the UI for an editor sessionm
//
// Common properties for all schema objects
type LayoutTemplate struct {
	Name                                                                                        string      `json:"name"`
	Tabs                                                                                        []TabDef    `json:"tabs"`
	Widgets                                                                                     []WidgetDef `json:"widgets"`
	// Version of the app                                                                                   
	AppVersion                                                                                  string      `json:"app_version"`
	// User who created the object                                                                          
	CreatedBy                                                                                   *string     `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle            
	// "format": "date-time"                                                                                
	CreatedOn                                                                                   string      `json:"created_on"`
	// Description of the object                                                                            
	Description                                                                                 *string     `json:"description,omitempty"`
	// User who last modified the object                                                                    
	LastModifiedBy                                                                              *string     `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":            
	// "date-time"                                                                                          
	LastModifiedOn                                                                              string      `json:"last_modified_on"`
	// Version of the schema                                                                                
	SchemaVersion                                                                               string      `json:"schema_version"`
	// Lifecycle state of the object                                                                        
	Status                                                                                      *Status     `json:"status,omitempty"`
	// Optional user-defined tags                                                                           
	Tags                                                                                        []string    `json:"tags,omitempty"`
	// Unique identifier for the object                                                                     
	UUID                                                                                        string      `json:"uuid"`
	Theme                                                                                       *string     `json:"theme,omitempty"`
}

// Defines the contents and purpose of a UI Tab panel
//
// Common properties for all schema objects
type TabDef struct {
	Type                                                                                        TabDefType `json:"type"`
	Title                                                                                       string     `json:"title"`
	Order                                                                                       int64      `json:"order"`
	// Version of the app                                                                                  
	AppVersion                                                                                  string     `json:"app_version"`
	// User who created the object                                                                         
	CreatedBy                                                                                   *string    `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle           
	// "format": "date-time"                                                                               
	CreatedOn                                                                                   string     `json:"created_on"`
	// Description of the object                                                                           
	Description                                                                                 *string    `json:"description,omitempty"`
	// User who last modified the object                                                                   
	LastModifiedBy                                                                              *string    `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":           
	// "date-time"                                                                                         
	LastModifiedOn                                                                              string     `json:"last_modified_on"`
	// Version of the schema                                                                               
	SchemaVersion                                                                               string     `json:"schema_version"`
	// Lifecycle state of the object                                                                       
	Status                                                                                      *Status    `json:"status,omitempty"`
	// Optional user-defined tags                                                                          
	Tags                                                                                        []string   `json:"tags,omitempty"`
	// Unique identifier for the object                                                                    
	UUID                                                                                        string     `json:"uuid"`
	Field                                                                                       *string    `json:"field,omitempty"`
	Format                                                                                      *string    `json:"format,omitempty"`
	Stringify                                                                                   *bool      `json:"stringify,omitempty"`
	Tooltip                                                                                     *string    `json:"tooltip,omitempty"`
}

// Atomic definition of a UI input widget for a given property
//
// Common properties for all schema objects
type WidgetDef struct {
	Type                                                                                        WidgetDefType `json:"type"`
	Label                                                                                       string        `json:"label"`
	Field                                                                                       string        `json:"field"`
	Tab                                                                                         string        `json:"tab"`
	Order                                                                                       int64         `json:"order"`
	// Version of the app                                                                                     
	AppVersion                                                                                  string        `json:"app_version"`
	// User who created the object                                                                            
	CreatedBy                                                                                   *string       `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle              
	// "format": "date-time"                                                                                  
	CreatedOn                                                                                   string        `json:"created_on"`
	// Description of the object                                                                              
	Description                                                                                 *string       `json:"description,omitempty"`
	// User who last modified the object                                                                      
	LastModifiedBy                                                                              *string       `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":              
	// "date-time"                                                                                            
	LastModifiedOn                                                                              string        `json:"last_modified_on"`
	// Version of the schema                                                                                  
	SchemaVersion                                                                               string        `json:"schema_version"`
	// Lifecycle state of the object                                                                          
	Status                                                                                      *Status       `json:"status,omitempty"`
	// Optional user-defined tags                                                                             
	Tags                                                                                        []string      `json:"tags,omitempty"`
	// Unique identifier for the object                                                                       
	UUID                                                                                        string        `json:"uuid"`
	Default                                                                                     interface{}   `json:"default"`
	Formatting                                                                                  *Formatting   `json:"formatting,omitempty"`
	ReadOnly                                                                                    *bool         `json:"read_only,omitempty"`
	Required                                                                                    *bool         `json:"required,omitempty"`
	Tooltip                                                                                     *string       `json:"tooltip,omitempty"`
	Undoable                                                                                    *bool         `json:"undoable,omitempty"`
	Validation                                                                                  *Validation   `json:"validation,omitempty"`
	VisibleIf                                                                                   *VisibleIf    `json:"visibleIf,omitempty"`
}

type Formatting struct {
	As *string `json:"as,omitempty"`
}

type Validation struct {
	MaxLength *int64  `json:"maxLength,omitempty"`
	MinLength *int64  `json:"minLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

type VisibleIf struct {
	Equals interface{} `json:"equals"`
	Field  *string     `json:"field,omitempty"`
}

// Tracks an atomic change in the app state data
//
// Common properties for all schema objects
type AppStateChange struct {
	// Version of the app                                                                                   
	AppVersion                                                                                  string      `json:"app_version"`
	// User who created the object                                                                          
	CreatedBy                                                                                   *string     `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle            
	// "format": "date-time"                                                                                
	CreatedOn                                                                                   string      `json:"created_on"`
	// Description of the object                                                                            
	Description                                                                                 *string     `json:"description,omitempty"`
	// User who last modified the object                                                                    
	LastModifiedBy                                                                              *string     `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":            
	// "date-time"                                                                                          
	LastModifiedOn                                                                              string      `json:"last_modified_on"`
	// Version of the schema                                                                                
	SchemaVersion                                                                               string      `json:"schema_version"`
	// Lifecycle state of the object                                                                        
	Status                                                                                      *Status     `json:"status,omitempty"`
	// Optional user-defined tags                                                                           
	Tags                                                                                        []string    `json:"tags,omitempty"`
	// Unique identifier for the object                                                                     
	UUID                                                                                        string      `json:"uuid"`
	// ID of the editor session the change belonged to                                                      
	EditorSessionID                                                                             *string     `json:"editor_session_id,omitempty"`
	NewValue                                                                                    interface{} `json:"new_value"`
	OldValue                                                                                    interface{} `json:"old_value"`
	// Nested property path to the changed value                                                            
	PropertyPath                                                                                []string    `json:"property_path,omitempty"`
	// ID of the editor session sub-tab the change belonged to                                              
	TabID                                                                                       *string     `json:"tab_id,omitempty"`
	// ID of the widget the change belonged to                                                              
	WidgetID                                                                                    *string     `json:"widget_id,omitempty"`
}

// Common properties for all schema objects
type CommonProps struct {
	// Version of the app                                                                                
	AppVersion                                                                                  string   `json:"app_version"`
	// User who created the object                                                                       
	CreatedBy                                                                                   *string  `json:"created_by,omitempty"`
	// ISO 8601 timestamp of when the object was created. @todo update tooling scripts to handle         
	// "format": "date-time"                                                                             
	CreatedOn                                                                                   string   `json:"created_on"`
	// Description of the object                                                                         
	Description                                                                                 *string  `json:"description,omitempty"`
	// User who last modified the object                                                                 
	LastModifiedBy                                                                              *string  `json:"last_modified_by,omitempty"`
	// ISO 8601 timestamp of last modification. @todo update tooling scripts to handle "format":         
	// "date-time"                                                                                       
	LastModifiedOn                                                                              string   `json:"last_modified_on"`
	// Version of the schema                                                                             
	SchemaVersion                                                                               string   `json:"schema_version"`
	// Lifecycle state of the object                                                                     
	Status                                                                                      *Status  `json:"status,omitempty"`
	// Optional user-defined tags                                                                        
	Tags                                                                                        []string `json:"tags,omitempty"`
	// Unique identifier for the object                                                                  
	UUID                                                                                        string   `json:"uuid"`
}

// Lifecycle state of the object
type Status string

const (
	Active   Status = "active"
	Archived Status = "archived"
	Deleted  Status = "deleted"
	Draft    Status = "draft"
)

type TabDefType string

const (
	Compound    TabDefType = "compound"
	Lorebook    TabDefType = "lorebook"
	Meta        TabDefType = "meta"
	SingleField TabDefType = "single_field"
	System      TabDefType = "system"
	TextList    TabDefType = "text_list"
)

type WidgetDefType string

const (
	Calendar  WidgetDefType = "calendar"
	Container WidgetDefType = "container"
	Custom    WidgetDefType = "custom"
	Dropdown  WidgetDefType = "dropdown"
	Image     WidgetDefType = "image"
	List      WidgetDefType = "list"
	None      WidgetDefType = "none"
	Number    WidgetDefType = "number"
	Text      WidgetDefType = "text"
	Textarea  WidgetDefType = "textarea"
	Toggle    WidgetDefType = "toggle"
)
