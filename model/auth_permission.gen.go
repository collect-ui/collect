// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAuthPermission = "auth_permission"

// AuthPermission mapped from table <auth_permission>
type AuthPermission struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string `gorm:"column:name;not null" json:"name"`
	ContentTypeID int32  `gorm:"column:content_type_id;not null" json:"content_type_id"`
	Codename      string `gorm:"column:codename;not null" json:"codename"`
}

// TableName AuthPermission's table name
func (*AuthPermission) TableName() string {
	return TableNameAuthPermission
}

func (*AuthPermission) PrimaryKey() []string {
	return []string{"id"}
}