package blob_storage

import (
	"fmt"
	"testing"
)

func TestBlobStorageService(t *testing.T) {

	t.SkipNow()
	t.Run("command exec", func(t *testing.T) {
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

}
