package utils

import (
	"strings"

	"github.com/EdwardAndrew/MovieBuff/pkg/config"
)

const BlankChar = "\u200B"

func FormatResponse(message string) string {
	return BlankChar + message
}

func RemoveBotPrefix(content string, botId string) string {
	command := strings.TrimPrefix(content, config.Get().Prefix+" ")
	command = strings.TrimPrefix(command, "<@"+botId+">")
	command = strings.TrimSpace(command)
	return command
}
