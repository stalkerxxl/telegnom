package types

type ParseMode string

const (
	ParseModeHTML       = ParseMode("HTML")
	ParseModeMarkdownV2 = ParseMode("MarkdownV2")
	ParseModeMarkdown   = ParseMode("Markdown")
)
