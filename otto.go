package otto

import (
	"os"
	"path/filepath"
)

func GetOTTO(input string) ([]byte, error) {
	var tokens []string
	split := splitStr(replaceYSDD(input))

	for _, s := range split {
		if isChinese(s) {
			tokens = append(tokens, pinyin(s)...)
		} else {
			if isYSDD(s) {
				tokens = append(tokens, s)
			} else {
				for _, s := range s {
					tokens = append(tokens, string(s))
				}
			}
		}
	}

	inputFiles, err := processTokens(tokens)
	if err != nil {
		return nil, err
	}
	return mergeWAV(inputFiles)
}

func SaveOTTO(data []byte, outputPath string) error {
	destDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}
