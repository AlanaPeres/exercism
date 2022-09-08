package techpalace

import (
    "fmt"
    "strings"
)
// WelcomeMessage retorna uma mensagem de boas-vindas para o cliente.
func WelcomeMessage(customer string) string {
 	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
	panic("Please implement the WelcomeMessage() function")
}
// AddBorder adiciona uma borda a uma mensagem de boas-vindas.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    borda := ""
	for i := 0; i < numStarsPerLine; i++ {
		borda += "*"
	}
	return fmt.Sprintf("%s\n%s\n%s", borda, welcomeMsg, borda)
	panic("Please implement the AddBorder() function")
}
// CleanupMessage limpa uma antiga mensagem de marketing.
func CleanupMessage(oldMsg string) string {
    linha := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(linha)
	panic("Please implement the CleanupMessage() function")
}
