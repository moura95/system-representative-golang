package util

import (
	"regexp"
	"strings"
)

func ValidateCPF(cpf string) bool {
	const pattern = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	matched, _ := regexp.MatchString(pattern, cpf)
	return matched
}

func ValidateCNPJ(cnpj string) bool {
	const pattern = `^^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$$`
	matched, _ := regexp.MatchString(pattern, cnpj)
	return matched
}

func ValidateEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func UnMaskCPFCNPJ(cpfcnpj string) string {
	cpfcnpj = strings.Replace(cpfcnpj, ".", "", -1)
	cpfcnpj = strings.Replace(cpfcnpj, "-", "", -1)
	cpfcnpj = strings.Replace(cpfcnpj, "/", "", -1)

	return cpfcnpj
}
