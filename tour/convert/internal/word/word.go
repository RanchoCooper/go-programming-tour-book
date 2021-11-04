package word

import (
    "strings"
    "unicode"
)

/**
 * @author Rancho
 * @date 2021/11/4
 */

func ToUpper(s string) string {
    return strings.ToUpper(s)
}

func ToLower(s string) string {
    return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string {
    s = strings.ReplaceAll(s, "_", " ")
    s = strings.Title(s)
    return strings.ReplaceAll(s, " ", "")
}

func UnderscoreToLowerCamelCase(s string) string {
    s = UnderscoreToUpperCamelCase(s)
    return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
    var output []rune
    for i, r := range s {
        if i == 0 {
            output = append(output, unicode.ToLower(r))
            continue
        }
        if unicode.IsUpper(r) {
            output = append(output, '_')
        }
        output = append(output, unicode.ToLower(r))
    }
    return string(output)
}