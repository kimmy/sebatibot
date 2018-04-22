package models

type TelegramUpdate struct {
    UpdateId  int64   `json:"update_id"`
    Message   Message `json:"message"`
}

type Message struct {
    MessageId int  `json:"message_id"`
    Text      string `json:"text"`
    Chat      Chat   `json:"chat"`
    From      From   `json:"from"`
}

type Chat struct {
    Id    int64  `json:"id"`
    Title string `json:"title"`
    Type  string `json:"supergroup"`
}

type From struct {
    Id           int64  `json:"id"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    IsBot        bool   `json:"is_bot"`
    LanguageCode string `json:"language_code"`
}