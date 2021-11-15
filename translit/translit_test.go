package translit_test

import (
	"fmt"
	"testing"

	"github.com/VoloshchenkoDN/text_transliteration/translit"
)

var (
	InputText  = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	OutputText = "ABVGDEYoZhZIJKLMNOPRSTUFKhTsChShShchYEJuJa abvgdeyozhzijklmnoprstufkhtschshshchyejuja"
)

type ErrorMessage struct {
	in, out, res string
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("Input text: %s, output text: %s is not equal the func's result: %s", e.in, e.out, e.res)

}

func RunFunc(in, out string, f func(string) string) error {
	if res := f(in); out != res {
		return &ErrorMessage{in, out, res}
	}
	return nil
}

func TestTranslate(t *testing.T) {
	if err := RunFunc(InputText, OutputText, translit.RU_EN.Convert); err != nil {
		t.Errorf(err.Error())
	}
}

func Benchmark_RUENConvertUsingBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := RunFunc(InputText, OutputText, translit.RU_EN.ConvertUsingBuilder); err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Benchmark_RUENConvertUsingByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := RunFunc(InputText, OutputText, translit.RU_EN.ConvertUsingByteSlice); err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Benchmark_RUENConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := RunFunc(InputText, OutputText, translit.RU_EN.Convert); err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Benchmark_RUENLowerCharsConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := RunFunc(InputText, OutputText, translit.RU_EN_LowerChars.Convert); err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Benchmark_CyrillicLatinTranslat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := RunFunc(InputText, OutputText, translit.CyrillicLatin.Translate); err != nil {
			b.Errorf(err.Error())
		}
	}
}
