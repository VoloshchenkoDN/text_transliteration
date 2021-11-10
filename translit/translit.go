package translit

import (
	"strings"
)

type TransliterationMapLower map[string]string
type TransliterationMapFull map[rune]string

func (t TransliterationMapLower) Convert(s string) string {
	var out, ch, lower_ch string

	for _, rune := range s {
		ch = string(rune)
		lower_ch = strings.ToLower(ch)

		if v, ok := t[lower_ch]; ok {
			if ch == lower_ch {
				out += v
			} else {
				out += strings.Title(v)
			}

		} else {
			out += ch
		}
	}
	return out
}

func (t TransliterationMapFull) Convert(s string) string {
	var out string

	for _, rune := range s {
		if v, ok := t[rune]; ok {
			out += v
		} else {
			out += string(rune)
		}
	}
	return out
}

func (t TransliterationMapFull) ConvertUsingByteSlice(s string) string {
	var buf []byte

	for _, rune := range s {
		if v, ok := t[rune]; ok {
			buf = append(buf, v...)
		} else {
			buf = append(buf, string(rune)...)
		}
	}
	return string(buf)
}

func (t TransliterationMapFull) ConvertUsingBuilder(s string) string {
	var b strings.Builder

	for _, rune := range s {
		if v, ok := t[rune]; ok {
			//fmt.Fprintf(&b, "%v", v)
			b.WriteString(v)
		} else {
			//fmt.Fprintf(&b, "%c", rune)
			b.WriteRune(rune)
		}
	}
	return b.String()
}

var RU_EN_LowerChars = TransliterationMapLower{
	"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e", "ё": "yo", "ж": "zh",
	"з": "z", "и": "i", "й": "j", "к": "k", "л": "l", "м": "m", "н": "n", "о": "o",
	"п": "p", "р": "r", "с": "s", "т": "t", "у": "u", "ф": "f", "х": "kh", "ц": "ts",
	"ч": "ch", "ш": "sh", "щ": "shch", "ъ": "", "ы": "y", "ь": "", "э": "e", "ю": "ju",
	"я": "ja",
}

var RU_EN = TransliterationMapFull{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo", 'ж': "zh",
	'з': "z", 'и': "i", 'й': "j", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
	'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts",
	'ч': "ch", 'ш': "sh", 'щ': "shch", 'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "ju",
	'я': "ja",
	'А': "A", 'Б': "B", 'В': "V", 'Г': "G", 'Д': "D", 'Е': "E", 'Ё': "Yo", 'Ж': "Zh",
	'З': "Z", 'И': "I", 'Й': "J", 'К': "K", 'Л': "L", 'М': "M", 'Н': "N", 'О': "O",
	'П': "P", 'Р': "R", 'С': "S", 'Т': "T", 'У': "U", 'Ф': "F", 'Х': "Kh", 'Ц': "Ts",
	'Ч': "Ch", 'Ш': "Sh", 'Щ': "Shch", 'Ъ': "", 'Ы': "Y", 'Ь': "", 'Э': "E", 'Ю': "Ju",
	'Я': "Ja",
}

// ----------------------------------------------
// my brother's implementation of transliteration
type TranslitMap map[rune]string

func (t TranslitMap) Translate(s string) string {
	out := make([]string, len(s))
	for i, rune := range s {
		if v, ok := t[rune]; ok {
			out[i] = v
		} else {
			out[i] = string(rune)
		}
	}
	return strings.Join(out, "")
}

var CyrillicLatin = TranslitMap{
	'а': "a", 'А': "A",
	'б': "b", 'Б': "B",
	'в': "v", 'В': "V",
	'г': "g", 'Г': "G",
	'д': "d", 'Д': "D",
	'е': "e", 'Е': "E",
	'ё': "yo", 'Ё': "Yo",
	'ж': "zh", 'Ж': "Zh",
	'з': "z", 'З': "Z",
	'и': "i", 'И': "I",
	'й': "j", 'Й': "J",
	'к': "k", 'К': "K",
	'л': "l", 'Л': "L",
	'м': "m", 'М': "M",
	'н': "n", 'Н': "N",
	'о': "o", 'О': "O",
	'п': "p", 'П': "P",
	'р': "r", 'Р': "R",
	'с': "s", 'С': "S",
	'т': "t", 'Т': "T",
	'у': "u", 'У': "U",
	'ф': "f", 'Ф': "F",
	'х': "kh", 'Х': "Kh",
	'ц': "ts", 'Ц': "Ts",
	'ч': "ch", 'Ч': "Ch",
	'ш': "sh", 'Ш': "Sh",
	'щ': "shch", 'Щ': "Shch",
	'ъ': "", 'Ъ': "",
	'ы': "y", 'Ы': "Y",
	'ь': "", 'Ь': "",
	'э': "e", 'Э': "E",
	'ю': "ju", 'Ю': "Ju",
	'я': "ja", 'Я': "Ja",
}
