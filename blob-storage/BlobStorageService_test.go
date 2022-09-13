package blob_storage

import (
	"fmt"
	"testing"
)

func TestBlobStorageService(t *testing.T) {

	t.SkipNow()
	t.Run("Azure command exec", func(t *testing.T) {
		var azureBlobBaseConfig *AzureBlobBaseConfig
		azureBlobBaseConfig = &AzureBlobBaseConfig{
			AccountKey:        "",
			AccountName:       "",
			Enabled:           true,
			BlobContainerName: "logs",
		}

		request := &BlobStorageRequest{
			StorageType:         BLOB_STORAGE_AZURE,
			SourceKey:           "devtron/cd-artifacts/110/110.zip",
			DestinationKey:      "job-artifact.zip",
			AzureBlobBaseConfig: azureBlobBaseConfig,
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		_, _, err := blobStorageServiceImpl.Get(request)
		fmt.Println(err)
	})
	t.Run("S3 command exec", func(t *testing.T) {
		awsS3BaseConfig := &AwsS3BaseConfig{
			AccessKey:   "",
			Passkey:     "",
			EndpointUrl: "devtron-minio.devtroncd:9000",
			BucketName:  "devtron-ci-log",
			Region:      "us-east-2",
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		request := &BlobStorageRequest{
			StorageType:     BLOB_STORAGE_S3,
			SourceKey:       "devtron/cd-artifacts/110/110.zip",
			DestinationKey:  "job-artifact.zip",
			AwsS3BaseConfig: awsS3BaseConfig,
		}
		err := blobStorageServiceImpl.PutWithCommand(request)
		fmt.Println(err)
	})

	t.Run("Gcp Upload Command exec", func(t *testing.T) {
		gcpConfig := &GcpBlobBaseConfig{
			CredentialFileJsonData: "",
			BucketName:             "kb-devtron-log",
			//BucketName: "kb-devtron-wo-version",
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		request := &BlobStorageRequest{
			StorageType:       BLOB_STORAGE_GCP,
			SourceKey:         "",
			DestinationKey:    "",
			GcpBlobBaseConfig: gcpConfig,
		}
		err := blobStorageServiceImpl.PutWithCommand(request)
		fmt.Println(err)
	})

	t.Run("Gcp Download Command exec", func(t *testing.T) {
		gcpConfig := &GcpBlobBaseConfig{
			CredentialFileJsonData: "",
			BucketName:             "kb-devtron-log",
			//BucketName: "kb-devtron-wo-version",
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		request := &BlobStorageRequest{
			StorageType:       BLOB_STORAGE_GCP,
			SourceKey:         "sample.go",
			DestinationKey:    "sample.go",
			GcpBlobBaseConfig: gcpConfig,
		}
		success, totalBytes, err := blobStorageServiceImpl.Get(request)
		fmt.Println(success, totalBytes, err)
	})

}
