package tag

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Southclaws/storyden/api/src/infra/db"
	"github.com/Southclaws/storyden/api/src/resources/user"
)

func TestGetTags(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	ctx := context.Background()

	d := db.TestDB(t)
	user.Seed(user.New(d))
	repo := New(d)

	// Create two tags
	t1, err := d.Tag.Create().SetName("tag1").Save(ctx)
	r.NoError(err)
	t2, err := d.Tag.Create().SetName("tag2").Save(ctx)
	r.NoError(err)

	// Create four posts, the first two are tagged with tag1 the third one is
	// tagged with tag2 and the fourth is tagged with both.
	_, err = d.Post.Create().
		SetBody("").
		SetShort("").
		SetFirst(true).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAuthorID(uuid.UUID(user.SeedUser_02_User.ID)).
		AddTagIDs(t1.ID).Save(ctx)
	r.NoError(err)
	_, err = d.Post.Create().
		SetBody("").
		SetShort("").
		SetFirst(true).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAuthorID(uuid.UUID(user.SeedUser_02_User.ID)).
		AddTagIDs(t1.ID).Save(ctx)
	r.NoError(err)
	_, err = d.Post.Create().
		SetBody("").
		SetShort("").
		SetFirst(true).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAuthorID(uuid.UUID(user.SeedUser_02_User.ID)).
		AddTagIDs(t2.ID).Save(ctx)
	r.NoError(err)
	_, err = d.Post.Create().
		SetBody("").
		SetShort("").
		SetFirst(true).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAuthorID(uuid.UUID(user.SeedUser_02_User.ID)).
		AddTagIDs(t2.ID, t1.ID).Save(ctx)
	r.NoError(err)

	// Searching for the prefix "ta" should get all our tags
	tags, err := repo.GetTags(ctx, "ta")
	r.NoError(err)
	r.NotNil(tags)

	// All 2 of them
	r.Len(tags, 2)
	// With the first tag having 3 posts (1, 2 and 4)
	a.Equal(tags[0].Name, "tag1")
	a.Equal(tags[0].Posts, 3)

	// And the second tag having 2 posts (3 and 4).
	a.Equal(tags[1].Name, "tag2")
	a.Equal(tags[1].Posts, 2)

	// Now search exactly for tag1
	tags2, err := repo.GetTags(ctx, "tag1")
	r.NoError(err)
	r.NotNil(tags)
	r.Len(tags2, 1)
	a.Equal(tags[0].Name, "tag1")
	a.Equal(tags[0].Posts, 3)

	// No results means a non-nil, empty list.
	none, err := repo.GetTags(ctx, "no")
	r.NoError(err)
	a.NotNil(none)
	a.Empty(none)
}
