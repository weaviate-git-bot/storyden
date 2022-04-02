// Code generated by entc, DO NOT EDIT.

package post

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldShort holds the string denoting the short field in the database.
	FieldShort = "short"
	// FieldFirst holds the string denoting the first field in the database.
	FieldFirst = "first"
	// FieldPinned holds the string denoting the pinned field in the database.
	FieldPinned = "pinned"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldRootPostId holds the string denoting the rootpostid field in the database.
	FieldRootPostId = "root_post_id"
	// FieldReplyPostId holds the string denoting the replypostid field in the database.
	FieldReplyPostId = "reply_post_id"
	// FieldCategoryId holds the string denoting the categoryid field in the database.
	FieldCategoryId = "category_id"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeRoot holds the string denoting the root edge name in mutations.
	EdgeRoot = "root"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeReplyTo holds the string denoting the replyto edge name in mutations.
	EdgeReplyTo = "replyTo"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeReacts holds the string denoting the reacts edge name in mutations.
	EdgeReacts = "reacts"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "categories"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "post_category"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "users"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "post_author"
	// RootTable is the table that holds the root relation/edge.
	RootTable = "posts"
	// RootColumn is the table column denoting the root relation/edge.
	RootColumn = "post_posts"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "post_posts"
	// ReplyToTable is the table that holds the replyTo relation/edge. The primary key declared below.
	ReplyToTable = "post_replyTo"
	// RepliesTable is the table that holds the replies relation/edge. The primary key declared below.
	RepliesTable = "post_replies"
	// TagsTable is the table that holds the tags relation/edge.
	TagsTable = "tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TagsColumn is the table column denoting the tags relation/edge.
	TagsColumn = "post_tags"
	// ReactsTable is the table that holds the reacts relation/edge.
	ReactsTable = "reacts"
	// ReactsInverseTable is the table name for the React entity.
	// It exists in this package in order to avoid circular dependency with the "react" package.
	ReactsInverseTable = "reacts"
	// ReactsColumn is the table column denoting the reacts relation/edge.
	ReactsColumn = "post_reacts"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldSlug,
	FieldBody,
	FieldShort,
	FieldFirst,
	FieldPinned,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserId,
	FieldRootPostId,
	FieldReplyPostId,
	FieldCategoryId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_posts",
	"post_posts",
	"react_post",
	"tag_posts",
	"user_posts",
}

var (
	// ReplyToPrimaryKey and ReplyToColumn2 are the table columns denoting the
	// primary key for the replyTo relation (M2M).
	ReplyToPrimaryKey = []string{"post_id", "replyTo_id"}
	// RepliesPrimaryKey and RepliesColumn2 are the table columns denoting the
	// primary key for the replies relation (M2M).
	RepliesPrimaryKey = []string{"post_id", "reply_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPinned holds the default value on creation for the "pinned" field.
	DefaultPinned bool
)
