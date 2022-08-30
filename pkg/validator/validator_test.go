package validator

import (
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestValidator_Module(t *testing.T) {
	i := i18n.New("en")
	i.LoadLanguages("./locales_test", "en", "tr")
	v := New(i)
	assert.NotEqual(t, nil, v)
	t.Run("ValidateStruct", func(t *testing.T) {
		t.Run("should return no errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "test",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 0, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
	})
	t.Run("ConnectCustom", func(t *testing.T) {
		v.ConnectCustom()
		t.Run("should return no errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"username"`
			}
			test := Test{
				Name: "test",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(
				t,
				0,
				len(errors),
			)
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"username"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(
				t,
				1,
				len(errors),
			)
		})
	})
	t.Run("RegisterTagName", func(t *testing.T) {
		v.RegisterTagName()
		t.Run("should return empty string", func(t *testing.T) {
			type Test struct {
				Name string `json:"-" validate:"required"`
			}
			test := Test{
				Name: "test",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(
				t,
				0,
				len(errors),
			)
		})
		t.Run("should return no errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "test",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 0, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
	})

	t.Run("translateErrorMessage", func(t *testing.T) {
		t.Run("should return no errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "test",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 0, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
		t.Run("should return errors", func(t *testing.T) {
			type Test struct {
				Name string `json:"name" validate:"required"`
			}
			test := Test{
				Name: "",
			}
			errors := v.ValidateStruct(test)
			assert.Equal(t, 1, len(errors))
		})
	})

	t.Run("CustomValidator", func(t *testing.T) {
		t.Run("validateUserName", func(t *testing.T) {
			t.Run("should return no errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"username"`
				}
				test := Test{
					Name: "test",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 0, len(errors))
			})
			t.Run("should return errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"username"`
				}
				test := Test{
					Name: "",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 1, len(errors))
			})
		})
		t.Run("validatePassword", func(t *testing.T) {
			t.Run("should return no errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"password"`
				}
				test := Test{
					Name: "Tat%at334_3s",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 0, len(errors))
			})
			t.Run("should return errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"password"`
				}
				test := Test{
					Name: "",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 1, len(errors))
			})
		})
		t.Run("validateLocale", func(t *testing.T) {
			t.Run("should return no errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"locale"`
				}
				test := Test{
					Name: "test",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 0, len(errors))
			})
			t.Run("should return errors", func(t *testing.T) {
				type Test struct {
					Name string `json:"name" validate:"locale"`
				}
				test := Test{
					Name: "",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 1, len(errors))
			})
		})
		t.Run("validateObjectId", func(t *testing.T) {
			t.Run("should return errors", func(t *testing.T) {
				type Test struct {
					Id string `json:"name" validate:"object_id"`
				}
				test := Test{
					Id: "5e9f8f8f8f8f8f8f8f8f8f8",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 1, len(errors))
			})
			t.Run("should return errors", func(t *testing.T) {
				type Test struct {
					Id string `json:"name" validate:"object_id"`
				}
				test := Test{
					Id: "",
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 1, len(errors))
			})
			t.Run("should return no errors", func(t *testing.T) {
				type Test struct {
					Id string `json:"name" validate:"object_id"`
				}
				test := Test{
					Id: primitive.NewObjectID().Hex(),
				}
				errors := v.ValidateStruct(test)
				assert.Equal(t, 0, len(errors))
			})
		})
	})
}
