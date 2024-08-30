package dockerOperations

import (
	"context"
	"github.com/devtron-labs/common-lib/utils/bean"
	"os"
	"testing"
)

const (
	PrivateEcrRegistryImageTest    = "PrivateEcrRegistryImage"
	PrivateHarborRegistryImageTest = "PrivateHarborRegistryImage"
	PrivateAzureRegistryImageTest  = "PrivateAzureRegistryImage"
	PrivateGcrRegistryImageTest    = "PrivateGcrRegistryImage"
	PrivateDockerRegistryImageTest = "PrivateDockerRegistryImage"
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
			name: "PublicDockerImageSingleArch",
			args: args{
				ctx:        context.Background(),
				image:      "logiqai/flash-brew-coffee:brew.v3.10.8",
				dockerAuth: nil,
			},
			want:    "sha256:993708ab5922ef7a2eec8b3bcf93822e58da8c4657bad73888f6bd92173ee1ad",
			wantErr: false,
		},
		{
			name: "PublicDockerImageMultiArch",
			args: args{
				ctx:        context.Background(),
				image:      "postgres:12.20-bullseye",
				dockerAuth: nil,
			},
			want:    "sha256:e1c0ba2f2a0bb8d1976c904d55ff7c817fcd5e922a938a05bb1698a6688028dd",
			wantErr: false,
		},
		{
			name: "PublicDockerImageMultiArchFromEcrRegistry",
			args: args{
				ctx:        context.Background(),
				image:      "public.ecr.aws/nginx/nginx:1.27.1-alpine3.20-perl",
				dockerAuth: nil,
			},
			want:    "sha256:8e42c4262557e07ef930a7568dd061661c94e57b1c0df75c68440ba94f29007b",
			wantErr: false,
		},
		{
			name: PrivateDockerRegistryImageTest,
			args: args{
				ctx:        context.Background(),
				image:      "prakash1001/gryffindor:6a824121-5-207",
				dockerAuth: &bean.DockerAuthConfig{IsRegistryPrivate: true},
			},
			want:    "sha256:11fc6035647ae0bd5b2b58518540bc72780435cb08293f744a48d9a93d2772d0",
			wantErr: false,
		},
		{
			name: "PublicBigDockerImageSingleArchAround3GB",
			args: args{
				ctx:        context.Background(),
				image:      "cimg/android:2024.08.1-node", // this image is above 4gb
				dockerAuth: nil,
			},
			want:    "sha256:dd376d1eb25e9402f6e3d126e766d529eaa817e707fc962603555cd9fda8b1cc",
			wantErr: false,
		},
		{
			name: PrivateEcrRegistryImageTest,
			args: args{
				ctx:        context.Background(),
				image:      "445808685819.dkr.ecr.ap-south-1.amazonaws.com/test:10f752d3-682-24528",
				dockerAuth: &bean.DockerAuthConfig{IsRegistryPrivate: true, RegistryType: bean.RegistryTypeEcr},
			},
			want:    "sha256:b116c920fba7845ab3721c0355eefe10dd0803277ff9c7616543b1e6b2982459",
			wantErr: false,
		},
		{
			name: PrivateHarborRegistryImageTest,
			args: args{
				ctx:        context.Background(),
				image:      "stage-harbor.devtron.info/devtron-test/pk:3fdcd758-1-40",
				dockerAuth: &bean.DockerAuthConfig{IsRegistryPrivate: true},
			},
			want:    "sha256:14c62c4fa09f02b3131ea4e6c48c152ac83b4ee05a7eec9d544105b7feb2a40a",
			wantErr: false,
		},
		{
			name: PrivateAzureRegistryImageTest,
			args: args{
				ctx:        context.Background(),
				image:      "devtroninc.azurecr.io/test:b0d3a379-1738-41271",
				dockerAuth: &bean.DockerAuthConfig{IsRegistryPrivate: true},
			},
			want:    "sha256:5b2b8575668ca6475e572989b7377d917482999160dadbe1e19f401aa54af0eb",
			wantErr: false,
		},
		{
			name: PrivateGcrRegistryImageTest,
			args: args{
				ctx:        context.Background(),
				image:      "asia-southeast1-docker.pkg.dev/centering-cable-321111/devtron-test/test:6a824121-232-180",
				dockerAuth: &bean.DockerAuthConfig{IsRegistryPrivate: true, RegistryType: bean.RegistryTypeGcr},
			},
			want:    "sha256:9750bc4ebbc3a45fecec39107b2c3b16bff603ee14bd6b565bc723cd085960b5",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		dockerHubUsername, dockerHubPass := os.Getenv("DOCKER_USERNAME"), os.Getenv("DOCKER_PASSWORD")
		ecrAccessKey, ecrSecretAccessKey, ecrRegion := os.Getenv("ECR_ACCESS_KEY"), os.Getenv("ECR_SECRET_ACCESS_KEY"), os.Getenv("ECR_REGION")
		azureUsername, azurePassword := os.Getenv("AZURE_USERNAME"), os.Getenv("AZURE_PASSWORD")
		harborUsername, harborPassword := "admin", os.Getenv("HARBOR_PASSWORD")
		gcrCredsJson := os.Getenv("GCR_CRED_JSON")
		switch tt.name {
		case PrivateEcrRegistryImageTest:
			tt.args.dockerAuth.AccessKeyEcr = ecrAccessKey
			tt.args.dockerAuth.SecretAccessKeyEcr = ecrSecretAccessKey
			tt.args.dockerAuth.EcrRegion = ecrRegion
		case PrivateHarborRegistryImageTest:
			tt.args.dockerAuth.Username = harborUsername
			tt.args.dockerAuth.Password = harborPassword
		case PrivateAzureRegistryImageTest:
			tt.args.dockerAuth.Username = azureUsername
			tt.args.dockerAuth.Password = azurePassword
		case PrivateGcrRegistryImageTest:
			tt.args.dockerAuth.CredentialFileJsonGcr = gcrCredsJson
		case PrivateDockerRegistryImageTest:
			tt.args.dockerAuth.Username = dockerHubUsername
			tt.args.dockerAuth.Password = dockerHubPass
		}
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
