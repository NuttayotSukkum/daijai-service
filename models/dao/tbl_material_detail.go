package dao

type MaterialDetail struct {
	Id                    int `gorm:"primarykey"`
	MaterialId            int
	Material              Material `gorm:"foreignKey:MaterialId"`
	MaterialFieldDetailId int
	MaterialFieldDetail   MaterialFieldDetail `gorm:"foreignKey:MaterialFieldDetailId"`
	CodeOrder             int
}

func (MaterialDetail) TableName() string {
	return "tbl_material_details"
}
