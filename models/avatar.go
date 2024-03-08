package models

import "github.com/akhil-is-watching/sainterview-backend/types"

type Avatar struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	AvatarID       string `json:"avatar_id"`
	RenderAvatarID string `json:"render_avatar_id"`
	Name           string `json:"name"`
	Gender         int    `json:"gender"`
	PreviewURL     string `json:"preview_url"`
	Avatar         string `json:"avatar"`
	Thumbnail      string `json:"thumbnail"`
	ThumbnailPath  string `json:"thumbnail_path"`
	AvatarPath     string `json:"avatar_path"`
}

func (m Avatar) AvatarInput() types.AvatarData {
	return types.AvatarData{
		ID:                m.ID,
		AvatarID:          m.AvatarID,
		OnlyShowOwned:     0,
		IsFullBody:        0,
		MonthlyInfo:       nil,
		RenderAvatarID:    m.RenderAvatarID,
		IsVerticalPercent: 0,
		Name:              m.Name,
		Gender:            m.Gender,
		PreviewURL:        m.PreviewURL,
		Avatar:            m.Avatar,
		NewUserOnly:       0,
		Thumbnail:         m.Thumbnail,
		ThumbnailPath:     m.ThumbnailPath,
		AvatarPath:        m.AvatarPath,
	}
}
