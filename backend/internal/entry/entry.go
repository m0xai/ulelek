package entry

type Entry struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateEntryRequest struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	IsDeleted bool   `json:"is_deleted"`
}

type CreateEntryRequest struct {
	Content   string `json:"content" binding:"required"`
	IsDeleted bool   `json:"is_deleted"`
}
