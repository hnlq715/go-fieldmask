package fieldmask

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hnlq715/go-fieldmask/testproto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/protobuf/field_mask"
)

var testUserFull *testproto.User
var testUserPartial *testproto.User

func init() {
	ts := &timestamp.Timestamp{
		Seconds: 5,
		Nanos:   6,
	}
	serializedTs, _ := proto.Marshal(ts)

	friend1 := &testproto.User{
		Id:          2,
		Username:    "friend",
		Role:        testproto.Role_REGULAR,
		Meta:        map[string]string{"foo": "bar"},
		Deactivated: true,
		Permissions: []testproto.Permission{testproto.Permission_READ, testproto.Permission_WRITE},
		Avatar: &testproto.Image{
			OriginalUrl: "original.jpg",
			ResizedUrl:  "resized.jpg",
		},
		Images: []*testproto.Image{
			{
				OriginalUrl: "FRIEND original_image1.jpg",
				ResizedUrl:  "FRIEND resized_image1.jpg",
			},
			{
				OriginalUrl: "FRIEND original_image2.jpg",
				ResizedUrl:  "FRIEND resized_image2.jpg",
			},
		},
		Tags: []string{"FRIEND tag1", "FRIEND tag2", "FRIEND tag3"},
		Name: &testproto.User_FemaleName{FemaleName: "Maggy"},
	}

	testUserFull = &testproto.User{
		Id:          1,
		Username:    "username",
		Role:        testproto.Role_ADMIN,
		Meta:        map[string]string{"foo": "bar"},
		Deactivated: true,
		Permissions: []testproto.Permission{testproto.Permission_READ, testproto.Permission_WRITE},
		Avatar: &testproto.Image{
			OriginalUrl: "original.jpg",
			ResizedUrl:  "resized.jpg",
		},
		Images: []*testproto.Image{
			{
				OriginalUrl: "original_image1.jpg",
				ResizedUrl:  "resized_image1.jpg",
			},
			{
				OriginalUrl: "original_image2.jpg",
				ResizedUrl:  "resized_image2.jpg",
			},
		},
		Tags:    []string{"tag1", "tag2", "tag3"},
		Friends: []*testproto.User{friend1},
		Name:    &testproto.User_MaleName{MaleName: "John"},
		Details: []*any.Any{
			{
				TypeUrl: "example.com/example/" + proto.MessageName(ts),
				Value:   serializedTs,
			},
		},
	}

	extraUser, err := ptypes.MarshalAny(testUserFull)
	if err != nil {
		panic(err)
	}

	testUserFull.ExtraUser = extraUser
	testUserPartial = &testproto.User{
		Id:       1,
		Username: "username",
	}
}

func validate(t *testing.T, test *testproto.User, req *testproto.UpdateUserRequest) {
	assert.Equal(t, testUserFull.Id, req.User.Id)
	assert.Equal(t, testUserFull.Permissions, req.User.Permissions)
	assert.Equal(t, testUserFull.Tags, req.User.Tags)
	assert.Equal(t, testUserFull.Images, req.User.Images)
	assert.Equal(t, testUserFull.Avatar.OriginalUrl, req.User.Avatar.OriginalUrl)
	assert.Equal(t, testUserFull.Details, req.User.Details)
}

func TestFieldMask(t *testing.T) {

	t.Run("normal", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: new(testproto.User),
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details"},
			},
		}

		err := Merge(testUserFull, req.User, req.GetFieldMask())
		assert.Nil(t, err)
		validate(t, testUserFull, req)
	})

	t.Run("all", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: new(testproto.User),
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"*"},
			},
		}

		err := Merge(testUserFull, req.User, req.GetFieldMask())
		assert.Nil(t, err)
		validate(t, testUserFull, req)
	})

	t.Run("sub all", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: new(testproto.User),
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Avatar.*"},
			},
		}

		err := Merge(testUserFull, req.User, req.GetFieldMask())
		assert.Nil(t, err)
		assert.NotEqual(t, testUserFull, req.User)
		assert.Equal(t, testUserFull.Avatar, req.User.Avatar)
	})

	t.Run("invalid path", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: new(testproto.User),
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details", "Tags."},
			},
		}

		err := Merge(testUserFull, req.User, req.GetFieldMask())
		assert.Nil(t, err)
		validate(t, testUserFull, req)
	})

	t.Run("nil mask", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: nil,
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details", "Tags."},
			},
		}

		err := Merge(testUserFull, req.User, nil)
		assert.Nil(t, err)
		assert.Nil(t, req.User)
	})

	t.Run("nil source", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: nil,
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details", "Tags."},
			},
		}

		err := Merge(nil, req.User, req.GetFieldMask())
		assert.Equal(t, ErrorNilSource, err)
		assert.Nil(t, req.User)
	})

	t.Run("nil dest", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: nil,
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details", "Tags."},
			},
		}

		err := Merge(testUserFull, nil, req.GetFieldMask())
		assert.Equal(t, ErrorNilDest, err)
		assert.Nil(t, req.User)
	})

	t.Run("type not match", func(t *testing.T) {
		req := &testproto.UpdateUserRequest{
			User: nil,
			FieldMask: &field_mask.FieldMask{
				Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details", "Tags."},
			},
		}

		err := Merge(testUserFull, req, req.GetFieldMask())
		assert.Equal(t, ErrorTypeNotMatch, err)
		assert.Nil(t, req.User)
	})
}

func BenchmarkFieldMask(b *testing.B) {

	b.Run("normal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := &testproto.UpdateUserRequest{
				User: new(testproto.User),
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"Id", "Avatar.OriginalUrl", "Tags", "Images", "Permissions", "Details"},
				},
			}

			err := Merge(testUserFull, req.User, req.GetFieldMask())
			assert.Nil(b, err)
		}
	})

	b.Run("all", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := &testproto.UpdateUserRequest{
				User: new(testproto.User),
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"*"},
				},
			}

			err := Merge(testUserFull, req.User, req.GetFieldMask())
			assert.Nil(b, err)
		}
	})
}
