package user_test

import (
	"auctionsystem/internal/common"
	"auctionsystem/internal/user"
	"auctionsystem/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		name     string
		request  user.GetRequestSchema
		expected user.GetResponseSchema
		wantErr  bool
		errMsg   string
	}{
		{
			name: "Success - User found",
			request: user.GetRequestSchema{
				ID: 1,
			},
			expected: user.GetResponseSchema{
				ResponseSchema: common.ResponseSchema{
					Code: 0,
					Msg:  "success",
				},
				Name: "golrice",
			},
			wantErr: false,
		},
		{
			name: "Failure - User not found",
			request: user.GetRequestSchema{
				ID: 2,
			},
			expected: user.GetResponseSchema{
				ResponseSchema: common.ResponseSchema{
					Code: 1,
					Msg:  "user not found",
				},
			},
			wantErr: true,
			errMsg:  "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化 Service 并注入 Mock
			service := user.NewUserService(testutil.NewMockUserRepository(), 2)

			// 调用被测方法
			resp, err := service.Get(tt.request)

			// 错误断言
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			// 返回值断言
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Name, resp.Name)
			assert.Equal(t, tt.expected.Code, resp.Code)
		})
	}
}
