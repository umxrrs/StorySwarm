package sentiment

import (
    "strings"
)

var toneKeywords = map[string][]string{
    "epic":   {"heroic", "grand", "adventure", "brave"},
    "spooky": {"dark", "eerie", "mysterious", "haunted"},
    "funny":  {"hilarious", "joke", "funny", "witty"},
}

func MatchesTone(text, tone string) bool {
    text = strings.ToLower(text)
    keywords, exists := toneKeywords[tone]
    if !exists {
        return true
    }
    for _, keyword := range keywords {
        if strings.Contains(text, keyword) {
            return true
        }
    }
    return false
}
