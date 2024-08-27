package dockerOperations

import (
	"context"
	"github.com/devtron-labs/common-lib/utils/bean"
	"testing"
)

func TestGetImageDigestByImage(t *testing.T) {
	type args struct {
		ctx        context.Context
		image      string
		dockerAuth *bean.DockerAuthConfig
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test1_PublicDockerImageSingleArch",
			args: args{
				ctx:        context.Background(),
				image:      "logiqai/flash-brew-coffee:brew.v3.10.8",
				dockerAuth: nil,
			},
			want:    "sha256:993708ab5922ef7a2eec8b3bcf93822e58da8c4657bad73888f6bd92173ee1ad",
			wantErr: false,
		},
		{
			name: "Test2_PublicDockerImageMultiArch",
			args: args{
				ctx:        context.Background(),
				image:      "postgres:12.20-bullseye",
				dockerAuth: nil,
			},
			want:    "sha256:e1c0ba2f2a0bb8d1976c904d55ff7c817fcd5e922a938a05bb1698a6688028dd",
			wantErr: false,
		},
		{
			name: "Test3_PublicDockerImageMultiArchFromEcrRegistry",
			args: args{
				ctx:        context.Background(),
				image:      "public.ecr.aws/nginx/nginx:1.27.1-alpine3.20-perl",
				dockerAuth: nil,
			},
			want:    "sha256:8e42c4262557e07ef930a7568dd061661c94e57b1c0df75c68440ba94f29007b",
			wantErr: false,
		},
		{
			name: "Test4_PrivateDockerImageSingleArch",
			args: args{
				ctx:   context.Background(),
				image: "prakash1001/gryffindor:6a824121-5-207",
				dockerAuth: &bean.DockerAuthConfig{
					Username: "", //username and password removed for security purpose
					Password: "",
				},
			},
			want:    "sha256:11fc6035647ae0bd5b2b58518540bc72780435cb08293f744a48d9a93d2772d0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetImageDigestByImage(tt.args.ctx, tt.args.image, tt.args.dockerAuth)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImageDigestByImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetImageDigestByImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
