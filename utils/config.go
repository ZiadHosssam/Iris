package utils

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//go:embed assets/config/config.toml
var defaultConfigContent string

//go:embed assets/config/themes/default.toml
var defaultThemeContent string

type ThemeColors struct {
	Base      string `toml:"base"`
	Panel     string `toml:"panel"`
	Hover     string `toml:"hover"`
	Border    string `toml:"border"`
	Muted     string `toml:"muted"`
	Sub       string `toml:"sub"`
	Text      string `toml:"text"`
	White     string `toml:"white"`
	Accent    string `toml:"accent"`
	AccDim    string `toml:"acc_dim"`
	AccBright string `toml:"acc_bright"`
}

type MainConfig struct {
	Window struct {
		ThemeProfile         string `toml:"theme_profile"`
		BorderStyle          string `toml:"border_style"`
		FontStyle            string `toml:"font_style"`
		ProgressBarChar      string `toml:"progress_bar_char"`
		ProgressBarEmptyChar string `toml:"progress_bar_empty_char"`
		CursorSymbol         string `toml:"cursor_symbol"`
		ShowLyricsByDefault  bool   `toml:"show_lyrics_by_default"`
		ShowQueueByDefault   bool   `toml:"show_queue_by_default"`
		TimeFormat           string `toml:"time_format"`
		SidebarAlignment     string `toml:"sidebar_alignment"`
		ScrollbarSymbol      string `toml:"scrollbar_symbol"`
	} `toml:"window"`

	Player struct {
		DefaultMusicDir  string `toml:"default_music_dir"`
		DefaultVolume    int    `toml:"default_volume"`
		VolumeSteps      int    `toml:"volume_steps"`
		SeekStepsSeconds int    `toml:"seek_steps_seconds"`
	} `toml:"player"`

	Colors ThemeColors `toml:"theme"`
}

func CheckConfig() (string, string, error) {
	ConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", "", err
	}

	IrisDir := filepath.Join(ConfigDir, "iris")
	themesDir := filepath.Join(IrisDir, "themes")
	masterConfigPath := filepath.Join(IrisDir, "config.toml")

	if _, err := os.Stat(masterConfigPath); os.IsNotExist(err) {
		if err := os.MkdirAll(themesDir, 0755); err != nil {
			return "", "", err
		}
		if err := os.WriteFile(masterConfigPath, []byte(defaultConfigContent), 0644); err != nil {
			return "", "", err
		}
		defaultThemePath := filepath.Join(themesDir, "default.toml")
		if err := os.WriteFile(defaultThemePath, []byte(defaultThemeContent), 0644); err != nil {
			return "", "", err
		}
	}
	return masterConfigPath, themesDir, nil
}

func LoadConfig() (*MainConfig, error) {
	masterConfigPath, themesDir, err := CheckConfig()
	if err != nil {
		return nil, err
	}

	var config MainConfig
	if _, err := toml.DecodeFile(masterConfigPath, &config); err != nil {
		return nil, err
	}

	if config.Window.ThemeProfile == "" {
		config.Window.ThemeProfile = "default"
	}
	themePath := filepath.Join(themesDir, config.Window.ThemeProfile+".toml")
	if _, err := toml.DecodeFile(themePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
