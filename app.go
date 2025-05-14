package main

import (
	"bot-builder/backend/models"
	"bot-builder/backend/utils"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed assets/default_config.json
var defaultConfig embed.FS

//go:embed schemas/*
var embeddedSchemas embed.FS

type AppPaths struct {
	BaseDir      string
	ConfigFile   string
	CardsDir     string
	TemplatesDir string
	LoreBooks    string
	ExportsDir   string
	LogsDir      string
	AppState     string
	Schemas      string
}

func getAppPaths() (*AppPaths, error) {
	// Use OS-appropriate config/data directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	base := filepath.Join(configDir, "bot-builder")
	return &AppPaths{
		BaseDir:      base,
		ConfigFile:   filepath.Join(base, "config.json"),
		CardsDir:     filepath.Join(base, "cards"),
		LoreBooks:    filepath.Join(base, "lorebooks"),
		TemplatesDir: filepath.Join(base, "templates"),
		ExportsDir:   filepath.Join(base, "exports"),
		LogsDir:      filepath.Join(base, "logs"),
		AppState:     filepath.Join(base, "app-state.json"),
		Schemas:      filepath.Join(base, "schemas"),
	}, nil
}

func initializeAppDirs(paths *AppPaths) error {
	dirs := []string{
		paths.BaseDir,
		paths.CardsDir,
		paths.TemplatesDir,
		paths.ExportsDir,
		paths.LogsDir,
		paths.LoreBooks,
		paths.Schemas,
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create default config file if missing
	if utils.PathExists(paths.ConfigFile) {
		data, err := defaultConfig.ReadFile("assets/default_config.json")
		if err != nil {
			return fmt.Errorf("failed to read embedded default config: %w", err)
		}
		if err := os.WriteFile(paths.ConfigFile, data, 0644); err != nil {
			return fmt.Errorf("failed to write config file: %w", err)
		}
	}

	// Copy embedded schema files if they don't already exist
	entries, err := embeddedSchemas.ReadDir("schemas")
	if err != nil {
		return fmt.Errorf("failed to read embedded schemas directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip subdirectories
		}

		targetPath := filepath.Join(paths.Schemas, entry.Name())
		if utils.PathExists(targetPath) {
			if err = os.Remove(targetPath); err != nil {
				fmt.Println(err)
			}
		}

		data, err := embeddedSchemas.ReadFile("schemas/" + entry.Name())
		if err != nil {
			return fmt.Errorf("failed to read embedded schema %s: %w", entry.Name(), err)
		}

		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			return fmt.Errorf("failed to write schema file %s: %w", targetPath, err)
		}
	}

	// Create app state file if missing
	if !utils.PathExists(paths.AppState) {
		if err := utils.WriteJSON(paths.AppState, models.NewAppState()); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	return nil
}

// Dummy config type
type Config struct {
	Theme string `json:"theme"`
}

func loadConfig(paths *AppPaths) (*Config, error) {
	data, err := os.ReadFile(paths.ConfigFile)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// App struct
type App struct {
	ctx    context.Context
	Config *Config
	Paths  *AppPaths
}

// NewApp creates a new App application struct
func NewApp() *App {
	paths, err := getAppPaths()
	if err != nil {
		panic(err)
	}

	err = initializeAppDirs(paths)
	if err != nil {
		panic(err)
	}

	config, err := loadConfig(paths)
	if err != nil {
		panic(err)
	}

	return &App{
		Config: config,
		Paths:  paths,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Example struct for a character card
type CharacterCard struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Notes  string `json:"notes"`
}

// Wails method: Load character card from JSON file
func (a *App) LoadCharacterCard(filename string) (*CharacterCard, error) {
	path := filepath.Join(a.Paths.BaseDir, "cards", filename)
	return utils.ReadJSON[CharacterCard](path)
}

// Wails method: Save character card to JSON file
func (a *App) SaveCharacterCard(filename string, card *CharacterCard) error {
	path := filepath.Join(a.Paths.BaseDir, "cards", filename)
	return utils.WriteJSON(path, card)
}

func (a *App) LoadAppState() *models.AppState {

	appState := models.NewAppState()
	if !utils.PathExists(a.Paths.AppState) {
		return &appState
	}

	appStateBytes, err := os.ReadFile(a.Paths.AppState)
	if err != nil {
		fmt.Println(err)
		return &appState
	}

	appState, err = models.UnmarshalAppState(appStateBytes)
	if err != nil {
		fmt.Println(err)
		return &appState
	}
	return &appState
}

func (a *App) SaveAppState(appState any) error {
	fmt.Println("appState", appState)
	return utils.WriteJSON(a.Paths.AppState, appState)
}

func (a *App) LoadTemplateSchema(templateName string) (template *models.LayoutTemplate, err error) {

	path := filepath.Join(a.Paths.TemplatesDir, templateName)
	return utils.ReadJSON[models.LayoutTemplate](path)
}

func (a *App) ResetAppState() error {
	appState := models.NewAppState()
	if err := utils.WriteJSON(a.Paths.AppState, appState); err != nil {
		return err
	}
	return nil
}

func (a *App) GetUserHome() (path string) {
	if path, err := os.UserConfigDir(); err != nil {
		return ""
	} else {

		return path
	}
}

func (a *App) SaveBlankTemplate(name string) {
	template := models.NewLayoutTemplate()
	template.Tabs = append(template.Tabs, models.NewTabDef())
	template.Widgets = append(template.Widgets, models.NewWidgetDef())

	path := filepath.Join(a.Paths.TemplatesDir, name)
	/* data, err := template.Marshal()
	if err != nil {
		return
	} */
	if err := utils.WriteJSON(path, template); err != nil {
		fmt.Println("Error writing to file")
	}
}
