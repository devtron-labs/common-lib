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
			StorageType:     BLOB_STORAGE_AZURE,
			SourceKey:       "devtron/cd-artifacts/110/110.zip",
			DestinationKey:  "job-artifact.zip",
			AzureBlobConfig: azureBlobBaseConfig,
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		_, _, err := blobStorageServiceImpl.Get(request)
		fmt.Println(err)
	})
	t.Run("S3 command exec", func(t *testing.T) {
		awsS3BaseConfig := &AwsS3BaseConfig{
			AccessKey:   "ZFRSVmt2WGZQMERuRVdTZ1BjY0VjMzFxamRJPQo",
			Passkey:     "QTdpSGY2RmR2OXQ3eVRod1Azcllhbzk3a0U4PQo",
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

}
