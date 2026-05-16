package main

import (
	"fmt"
	"iris/utils"
	"os"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded config: %+v\n", cfg)

	// Data Verification
	fmt.Printf("Theme Profile: %s\n", cfg.Window.ThemeProfile)
	fmt.Printf("Window Border Style: %s\n", cfg.Window.BorderStyle)
	fmt.Printf("Default Music Location: %s\n", cfg.Player.DefaultMusicDir)

	fmt.Printf("Base Color: %s\n", cfg.Colors.Base)
	fmt.Printf("Panel Color: %s\n", cfg.Colors.Panel)

	songs, err := utils.ScanAudioFiles(cfg.Player.DefaultMusicDir)
	if err != nil {
		fmt.Println("Error scanning audio files:", err)
		return
	}

	fmt.Printf("Audio Files Found: %d\n", len(songs))

	for i, song := range songs {
		fmt.Printf(" [%d] %s\n", i+1, song)
	}
}
