package gui

import (
	"slices"
	"strings"

	"github.com/Baipyrus/ProxySwitcher/util"
)

func getConfigByName(cfgPath string, name string) *util.Config {
	configs, _ := util.ReadConfigs(cfgPath)

	idx := slices.IndexFunc(
		configs,
		func(other *util.Config) bool {
			ownName :=
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ToLower(name),
						" ", ""),
					"-", "_")

			otherName :=
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ToLower(other.Name),
						" ", ""),
					"-", "_")

			return ownName == otherName
		},
	)

	if idx == -1 {
		return nil
	}

	return configs[idx]
}

func splitPath(path string) [][]string {
	parts := strings.Split(path, "-")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	var result [][]string
	for i := range parts {
		// Concatenate all paths up to current
		result = append(result, parts[:i+1])
	}

	return result
}

func parseCategories(cfgPath string) map[string][]string {
	// Read configs, split into categories based on name
	cats := make(map[string][]string)
	configs, _ := util.ReadConfigs(cfgPath)

	for _, cfg := range configs {
		for _, levels := range splitPath(cfg.Name) {
			// Extract hierarchy and item name
			parent := strings.Join(levels[:len(levels)-1], " - ")
			item := levels[len(levels)-1]

			// Initialize category slice if empty
			if _, ok := cats[parent]; !ok {
				cats[parent] = []string{item}
				continue
			}

			// Ignore duplicate entries
			if slices.Contains(cats[parent], item) {
				continue
			}

			cats[parent] = append(cats[parent], item)
		}
	}

	return cats
}
