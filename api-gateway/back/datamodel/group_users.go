package datamodel

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetGroup_usersTable() (group_usersTable models.Table) {

	group_usersTable.Info.FieldList = []types.Field{{
			Head:     "Id",
			Field:    "id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Group_id",
			Field:    "group_id",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Uid",
			Field:    "uid",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "创建时间",
			Field:    "create_time",
			TypeName: "datetime",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	group_usersTable.Info.Table = "group_users"
	group_usersTable.Info.Title = "好友列表"
	group_usersTable.Info.Description = "Group_users"

	group_usersTable.Form.FormList = []types.Form{{
			Head:     "Id",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Group_id",
			Field:    "group_id",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Uid",
			Field:    "uid",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "创建时间",
			Field:    "create_time",
			TypeName: "datetime",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	group_usersTable.Form.Table = "group_users"
	group_usersTable.Form.Title = "好友列表"
	group_usersTable.Form.Description = "Group_users"

	group_usersTable.ConnectionDriver = "mysql"

	return
}