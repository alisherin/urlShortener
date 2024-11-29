package models

type Link struct {
	ID         uint64  `json:"id" db:"id"`
	Url        string  `json:"url" db:"url"`
	Short      string  `json:"short" db:"short"`
	UserId     uint64  `json:"userId" db:"users_id"`
	ExpiredAt  string  `json:"expiredAt" db:"expired_at"`
	Counter    uint64  `json:"counter" db:"counter"`
	LastUsedAt *string `json:"lastUsedAt" db:"last_used_at"`
}

type LinkSave struct {
	Url       string `json:"url"`
	Short     string `json:"short"`
	UserId    uint64 `json:"userId"`
	ExpiredAt string `json:"expiredAt"`
}
