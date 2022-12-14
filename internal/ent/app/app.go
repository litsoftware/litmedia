// Code generated by entc, DO NOT EDIT.

package app

import (
	"time"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldOperatorID holds the string denoting the operator_id field in the database.
	FieldOperatorID = "operator_id"
	// FieldEncryptedOperatorRsaPublicKey holds the string denoting the encrypted_operator_rsa_public_key field in the database.
	FieldEncryptedOperatorRsaPublicKey = "encrypted_operator_rsa_public_key"
	// FieldEncryptedAppPrivateKey holds the string denoting the encrypted_app_private_key field in the database.
	FieldEncryptedAppPrivateKey = "encrypted_app_private_key"
	// FieldEncryptedAppPublicKey holds the string denoting the encrypted_app_public_key field in the database.
	FieldEncryptedAppPublicKey = "encrypted_app_public_key"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldConf holds the string denoting the conf field in the database.
	FieldConf = "conf"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldAppSecret holds the string denoting the app_secret field in the database.
	FieldAppSecret = "app_secret"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIPWhitelist holds the string denoting the ip_whitelist field in the database.
	FieldIPWhitelist = "ip_whitelist"
	// Table holds the table name of the app in the database.
	Table = "apps"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldOperatorID,
	FieldEncryptedOperatorRsaPublicKey,
	FieldEncryptedAppPrivateKey,
	FieldEncryptedAppPublicKey,
	FieldTitle,
	FieldDescription,
	FieldConf,
	FieldAppID,
	FieldAppSecret,
	FieldStatus,
	FieldIPWhitelist,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultConf holds the default value on creation for the "conf" field.
	DefaultConf string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
)
