package models

import "time"

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password,omitempty"`
    Balance   int       `json:"balance"`
    IsMember  bool      `json:"is_member"`
    CreatedAt time.Time `json:"created_at"`
}

type UserCreateRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Balance  int    `json:"balance"`
    IsMember bool   `json:"is_member"`
}

type UserUpdateRequest struct {
    Name     string `json:"name,omitempty"`
    Balance  int    `json:"balance,omitempty"`
    IsMember bool   `json:"is_member,omitempty"`
}