package translit_test

import (
	"testing"

	"github.com/VoloshchenkoDN/text_transliteration/translit"
)

func TestTranslate(t *testing.T) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for src, res := range test_map {
		if translit.RU_EN.Convert(src) != res {
			t.Errorf("source: %s, result: %s", src, res)
		}
	}
}

func Benchmark_RUENConvertUsingBuilder(b *testing.B) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for i := 0; i < b.N; i++ {
		for in, out := range test_map {
			if res := translit.RU_EN.ConvertUsingBuilder(in); res != out {
				b.Errorf("input: %s, output: %s, result: %s", in, out, res)
			}
		}
	}
}

func Benchmark_RUENConvertUsingByteSlice(b *testing.B) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for i := 0; i < b.N; i++ {
		for in, out := range test_map {
			if res := translit.RU_EN.ConvertUsingByteSlice(in); res != out {
				b.Errorf("input: %s, output: %s, result: %s", in, out, res)
			}
		}
	}
}

func Benchmark_RUENConvert(b *testing.B) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for i := 0; i < b.N; i++ {
		for in, out := range test_map {
			if res := translit.RU_EN.Convert(in); res != out {
				b.Errorf("input: %s, output: %s, result: %s", in, out, res)
			}
		}
	}
}

func Benchmark_RUENLowerCharsConvert(b *testing.B) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for i := 0; i < b.N; i++ {
		for in, out := range test_map {
			if res := translit.RU_EN_LowerChars.Convert(in); res != out {
				b.Errorf("input: %s, output: %s, result: %s", in, out, res)
			}
		}
	}
}

func Benchmark_CyrillicLatinTranslat(b *testing.B) {
	src := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	res := "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
	test_map := map[string]string{
		src: res,
	}

	for i := 0; i < b.N; i++ {
		for in, out := range test_map {
			if res := translit.CyrillicLatin.Translate(in); res != out {
				b.Errorf("input: %s, output: %s, result: %s", in, out, res)
			}
		}
	}
}
