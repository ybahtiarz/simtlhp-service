package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

func (LHP) TableName() string {
	return "data_umum_lhp"
}

type LHP struct {
	ID            uint64         `json:"id" gorm:"primaryKey;column:id"`
	NomorLHP      string         `json:"nomor_lhp" gorm:"column:nomor_lhp;uniqueIndex"`
	TanggalLHP    time.Time      `json:"tanggal_lhp" gorm:"column:tanggal_lhp"`
	NomorST       string         `json:"nomor_st" gorm:"column:nomor_surat_tugas"`
	TanggalST     time.Time      `json:"tanggal_st" gorm:"column:tanggal_surat_tugas"`
	TahunPKPT     uint64         `json:"tahun_pkpt" gorm:"column:tahun_pkpt"`
	TAObrik       uint64         `json:"ta_obrik" gorm:"column:tahun_anggaran_obrik"`
	Obrik         string         `json:"obrik" gorm:"column:nama_obrik"`
	KodeObrik     uint64         `json:"kode_obrik" gorm:"column:kode_obrik"`
	PJObrik       string         `json:"pj_obrik" gorm:"column:penanggung_jawab_obrik"`
	NIPPJObrik    string         `json:"nip_pj" gorm:"column:nip_penanggung_jawab_obrik"`
	KodeProvinsi  string         `json:"kode_provinsi" gorm:"column:kode_provinsi"`
	KodeKabupaten string         `json:"kode_kabupaten" gorm:"column:kode_kabupaten"`
	CreatedAt     time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Temuan        []Temuan       `json:"temuan,omitempty" gorm:"foreignKey:LHPID;referencesID"`
}

type CreateLHPRequest struct {
	NomorLHP      string `json:"nomor_lhp" binding:"required"`
	TanggalLHP    string `json:"tanggal_lhp" binding:"required"` // Format YYYY-MM-DD, diparse di Service
	NomorST       string `json:"nomor_surat_tugas"`
	TanggalST     string `json:"tanggal_surat_tugas" binding:"required"` // Format YYYY-MM-DD, diparse di Service
	TahunPKPT     uint64 `json:"tahun_pkpt"`
	TAObrik       int    `json:"tahun_anggaran_obrik"`
	NamaObrik     string `json:"nama_obrik" binding:"required"`
	KodeObrik     uint64 `json:"kode_obrik"`
	PJObrik       string `json:"pj_obrik"`
	NIPPJObrik    string `json:"nip_pj"`
	KodeProvinsi  string `json:"kode_provinsi"`
	KodeKabupaten string `json:"kode_kabupaten"`
}

type LHPFilter struct {
	Page      int
	Limit     int
	Search    *string
	StartDate *string
	Enddate   *string
	Status    *string
}

type LHPRepository interface {
	Create(ctx context.Context, lhp *LHP) error
	FindAll(ctx context.Context, filter LHPFilter) ([]LHP, int64, error)
	GetByID(ctx context.Context, id uint64) (*LHP, error)
	GetMatrik(ctx context.Context, filter LHPFilter) ([]LHP, int64, error)
}

type LHPService interface {
	CreateLHP(ctx context.Context, req CreateLHPRequest) (*LHP, error)
	GetLHPList(ctx context.Context, filter LHPFilter) ([]LHP, int64, error)
	GetLHPByID(ctx context.Context, id uint64) (*LHP, error)
	GetMatrik(ctx context.Context, filter LHPFilter) ([]LHP, int64, error)
}
