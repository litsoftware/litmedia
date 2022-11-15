package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type Media struct {
	ent.Schema
}

type MediaSize struct {
	Size     int64         `json:"size" bson:"size"`
	Width    int           `json:"width" bson:"width"`
	Height   int           `json:"height" bson:"height"`
	Length   int           `json:"length" bson:"length"`
	Duration time.Duration `json:"duration" bson:"duration"`
	WhRate   float64       `json:"wh_rate"` // 长宽比
}

// Fields of the Refund.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").Optional(),
		field.Int("user_id").Optional().Comment("本系统用户id"),
		field.Int("app_id").Optional(),
		field.String("sn").Optional(),
		field.String("org_file_name").Comment("原始文件名"),
		field.String("file_name").Comment("保存的文件名称"),
		field.JSON("size", MediaSize{}).Optional().Comment("尺寸信息"),
		field.String("mime").Optional().Comment("文件mime"),
		field.String("ext").Optional().Comment("扩展名"),
		field.Int("ref_count").Optional().Default(0).Comment("引用计数"),
		field.Int("level").Optional().Default(0).Comment("资源的级别 普通， 敏感资料"), // enum MediaLevel
		field.Int("type").Optional().Default(0),                            // enum MediaType
		field.Int("status").Optional().Default(0),                          // enum MediaStatus
		field.Int("reason_type").Optional().Default(0),                     // enum MediaReasonType
		field.String("save_path").Optional().Comment("保存路径"),
		field.String("full_path").Optional(),
		field.Bool("is_encrypted").Optional().Default(false),
		field.String("reason").Optional().Default("").Comment("图片被强制操作或审核失败的理由/原因"),
		field.Int64("out_user_id").Optional().Comment("外部用户id"),
		field.Int64("company_id").Optional().Comment("公司id"),
	}
}

// Edges of the Refund.
func (Media) Edges() []ent.Edge {
	return nil
}

func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "medias"},
	}
}

func (Media) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
