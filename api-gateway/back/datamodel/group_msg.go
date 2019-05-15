package datamodel

import (
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template/types"
)

func GetGroup_msgTable() (group_msgTable models.Table) {

	group_msgTable.Info.FieldList = []types.Field{{
		Head:     "Id",
		Field:    "id",
		TypeName: "bigint",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Uuid",
		Field:    "uuid",
		TypeName: "varchar",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Group_id",
		Field:    "group_id",
		TypeName: "varchar",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "聊天内容",
		Field:    "content",
		TypeName: "blob",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "From_uid",
		Field:    "from_uid",
		TypeName: "int",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "创建时间",
		Field:    "create_time",
		TypeName: "datetime",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "聊天类型",
		Field:    "content_type",
		TypeName: "varchar",
		Sortable: false,
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	},}

	group_msgTable.Info.Table = "group_msg"
	group_msgTable.Info.Title = "聊天消息"
	group_msgTable.Info.Description = "Group_msg"

	group_msgTable.Form.FormList = []types.Form{{
		Head:     "Id",
		Field:    "id",
		TypeName: "bigint",
		Default:  "",
		Editable: false,
		FormType: "default",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Uuid",
		Field:    "uuid",
		TypeName: "varchar",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Group_id",
		Field:    "group_id",
		TypeName: "varchar",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Content",
		Field:    "content",
		TypeName: "blob",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "From_uid",
		Field:    "from_uid",
		TypeName: "int",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Create_time",
		Field:    "create_time",
		TypeName: "datetime",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Content_type",
		Field:    "content_type",
		TypeName: "varchar",
		Default:  "",
		Editable: false,
		FormType: "text",
		ExcuFun: func(model types.RowModel) interface{} {
			return model.Value
		},
	},}

	group_msgTable.Form.Table = "group_msg"
	group_msgTable.Form.Title = "聊天消息"
	group_msgTable.Form.Description = "Group_msg"

	group_msgTable.ConnectionDriver = "mysql"

	return
}
