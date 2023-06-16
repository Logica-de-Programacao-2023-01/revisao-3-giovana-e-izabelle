package q5

import "strings"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	alfa := ""
	for _, caracter := range s {
		if isAlfa(caracter) {
			alfa += string(caracter)
		}
	}
	return IsPalindromeString(alfa)
}
func isAlfa(caracter rune) bool {
	return ('a' <= caracter && caracter <= 'z') || ('O' <= caracter && caracter <= '9')
}

func IsPalindromeString(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
