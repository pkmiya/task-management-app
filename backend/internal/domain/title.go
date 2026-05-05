package domain

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Title はタスクタイトルの値オブジェクト（要件 5.2）。
type Title struct {
	value string
}

func (t Title) String() string {
	return t.value
}

// ParseTitle は入力を検証し Title を構築する。
// 先頭末尾は unicode.IsSpace に基づきトリムする（strings.TrimFunc と整合）。
func ParseTitle(raw string) (Title, error) {
	trimmed := strings.TrimFunc(raw, unicode.IsSpace)
	if trimmed == "" {
		return Title{}, fmt.Errorf("%w: title is empty after trim", ErrValidation)
	}
	if utf8.RuneCountInString(trimmed) > 255 {
		return Title{}, fmt.Errorf("%w: title exceeds 255 Unicode code points", ErrValidation)
	}
	return Title{value: trimmed}, nil
}
