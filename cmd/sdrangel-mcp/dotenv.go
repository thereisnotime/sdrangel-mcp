package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// loadDotenv loads environment variables from a .env file.
// Priority: explicit path > ./.env > $XDG_CONFIG_HOME/sdrangel-mcp/.env > ~/.config/sdrangel-mcp/.env
func loadDotenv(explicit string) error {
	paths := dotenvCandidates(explicit)
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return readDotenv(p)
		}
	}
	return nil
}

func dotenvCandidates(explicit string) []string {
	if explicit != "" {
		return []string{explicit}
	}
	candidates := []string{".env"}
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		candidates = append(candidates, filepath.Join(xdg, "sdrangel-mcp", ".env"))
	}
	home, err := os.UserHomeDir()
	if err == nil {
		candidates = append(candidates, filepath.Join(home, ".config", "sdrangel-mcp", ".env"))
	}
	return candidates
}

func readDotenv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)
		if os.Getenv(key) == "" {
			if err := os.Setenv(key, value); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}
