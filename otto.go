package otto

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetOTTO(input string) ([]byte, error) {
	var tokens []string
	split := splitStr(replaceYSDD(input))

	for _, s := range split {
		if isChinese(s) {
			for _, c := range pinyin(s) {
				tokens = append(tokens, fmt.Sprintf("(%s)", c))
			}
		} else {
			if isYSDD(s) || isSingle(s) {
				tokens = append(tokens, s)
			} else {
				for _, c := range s {
					tokens = append(tokens, string(c))
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
