package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func UnixHomeDir(path string) string {
	if path == "" || path[0] != '~' {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}
	if len(path) == 1 {
		return home
	}
	return filepath.Join(home, path[1:])
}

func ScanAudioFiles(dirPath string) ([]string, error) {
	MusicPath := UnixHomeDir(dirPath)

	entries, err := os.ReadDir(MusicPath)
	if err != nil {
		return nil, err
	}

	var audioFiles []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == ".mp3" || ext == ".wav" || ext == ".flac" || ext == ".aac" || ext == ".ogg" || ext == ".m4a" {
			audioFiles = append(audioFiles, filepath.Join(MusicPath, entry.Name()))
		}
	}
	return audioFiles, nil
}
