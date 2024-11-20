package glossary

// sourceText string
// Текст на языке оригинала.
// Обязательно при переводе с глоссарием.

// translatedText string
// Текст на языке перевода.
// Обязательно при переводе с глоссарием.

// exact bool
// Игнорировать падежи и т.п.
// Поддерживаемые языки ar, bg, cs, de, en, es, fr, it, kk, pl, ru, tr, tt, uk

var Glossary = []map[string]any{
	{
		"sourceText":     "switch",
		"translatedText": "коммутатор",
		"exact":          false,
	},
	{
		"sourceText":     "listening",
		"translatedText": "прослушивание",
		"exact":          false,
	},
	{
		"sourceText":     "PVST",
		"translatedText": "PVST",
		"exact":          false,
	},
	{
		"sourceText":     "Rapid PVST",
		"translatedText": "Rapid PVST",
		"exact":          false,
	},
	{
		"sourceText":     "forwarding",
		"translatedText": "пересылка",
		"exact":          false,
	},
	{
		"sourceText":     "PVST+",
		"translatedText": "PVST+",
		"exact":          false,
	},
}
