package runner

import "strings"

func ExtractParams(params []string) string {

	var sb strings.Builder

	for _, param := range params {
		sb.WriteString(param)
		sb.WriteString(" ")
	}

	return sb.String()
}
