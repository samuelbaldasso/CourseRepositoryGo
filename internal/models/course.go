package models

type Course struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Duration    int     `json:"duration"` // Duration in hours
}