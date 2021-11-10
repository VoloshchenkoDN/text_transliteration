package main

import (
	"fmt"

	"github.com/VoloshchenkoDN/text_transliteration/translit"
)

func main() {
	s := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	fmt.Println(translit.RU_EN_LowerChars.Convert(s))
	fmt.Println(translit.RU_EN.Convert(s))
	fmt.Println(translit.RU_EN.ConvertUsingByteSlice(s))
	fmt.Println(translit.RU_EN.ConvertUsingBuilder(s))
}
