package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/litsoftware/litmedia/pkg/d"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.Int("operator_id").Optional(),
		field.Text("encrypted_operator_rsa_public_key").Optional().Comment("用户公钥（operator上传)"),
		field.Text("encrypted_app_private_key").Optional().Comment("本平台为当前应用生成的应用私钥"),
		field.Text("encrypted_app_public_key").Optional().Comment("本平台为当前应用生成的应用公钥"),
		field.String("title").Optional(),
		field.String("description").Optional(),
		field.String("conf").Optional().Default("general").Comment("使用的配置名称"),
		field.String("app_id").Optional().Comment("应用id"),
		field.String("app_secret").Optional().Comment("应用密钥"),
		field.Int("status").Optional().Default(1).Comment("是否启用。 0： 未设置， 1： 启用， 2：不启用"),
		field.JSON("ip_whitelist", []d.H{}).Optional().Comment("允许发起支付的ip白名单"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "apps"},
	}
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
