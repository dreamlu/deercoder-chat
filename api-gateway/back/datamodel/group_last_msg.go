package datamodel

//import (
//	"github.com/chenhg5/go-admin/template/types"
//	"github.com/chenhg5/go-admin/plugins/admin/models"
//)
//
//func GetGroup_last_msgTable() (group_last_msgTable models.Table) {
//
//	group_last_msgTable.Info.FieldList = []types.Field{{
//			Head:     "Id",
//			Field:    "id",
//			TypeName: "int",
//			Sortable: false,
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Group_id",
//			Field:    "group_id",
//			TypeName: "varchar",
//			Sortable: false,
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Uid",
//			Field:    "uid",
//			TypeName: "int",
//			Sortable: false,
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Last_group_msg_uuid",
//			Field:    "last_group_msg_uuid",
//			TypeName: "varchar",
//			Sortable: false,
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Is_read",
//			Field:    "is_read",
//			TypeName: "tinyint",
//			Sortable: false,
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},}
//
//	group_last_msgTable.Info.Table = "group_last_msg"
//	group_last_msgTable.Info.Title = "Group_last_msg"
//	group_last_msgTable.Info.Description = "Group_last_msg"
//
//	group_last_msgTable.Form.FormList = []types.Form{{
//			Head:     "Id",
//			Field:    "id",
//			TypeName: "int",
//			Default:  "",
//			Editable: false,
//			FormType: "default",
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Group_id",
//			Field:    "group_id",
//			TypeName: "varchar",
//			Default:  "",
//			Editable: false,
//			FormType: "text",
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Uid",
//			Field:    "uid",
//			TypeName: "int",
//			Default:  "",
//			Editable: false,
//			FormType: "text",
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Last_group_msg_uuid",
//			Field:    "last_group_msg_uuid",
//			TypeName: "varchar",
//			Default:  "",
//			Editable: false,
//			FormType: "text",
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},{
//			Head:     "Is_read",
//			Field:    "is_read",
//			TypeName: "tinyint",
//			Default:  "",
//			Editable: false,
//			FormType: "text",
//			ExcuFun: func(model types.RowModel) interface{} {
//				return model.Value
//			},
//		},	}
//
//	group_last_msgTable.Form.Table = "group_last_msg"
//	group_last_msgTable.Form.Title = "Group_last_msg"
//	group_last_msgTable.Form.Description = "Group_last_msg"
//
//	group_last_msgTable.ConnectionDriver = "mysql"
//
//	return
//}