package utilities

// FS - достаем наименование файла с нужным шрифтом
func FS(args []string) string {
	for _, font := range args[1:] {
		switch font {
		case "standard", "shadow", "thinkertoy":
			return font + ".txt"
		}
	}

	return "standard.txt"
}
