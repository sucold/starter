package table

type Model struct {
	ID        int64 `gorm:"primarykey" json:"id"`
	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}
