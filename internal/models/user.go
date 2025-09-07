package models

// User representa um usuário da plataforma
// ID é autoincrementado (em memória, igual ao Course)
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}
