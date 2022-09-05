package blob_storage

type BlobStorageRequest struct {
	StorageType          BlobStorageType
	Region               string
	Endpoint             string // for s3 compatible storage
	BucketName           string
	Key                  string
	AccessKey            string
	Passkey              string
	FileDownloadLocation string
	AzureBlobConfig      *AzureBlobConfig
}

type BlobStorageType string

const (
	BLOB_STORAGE_AZURE BlobStorageType = "AZURE"
	BLOB_STORAGE_S3                    = "S3"
	BLOB_STORAGE_GCP                   = "GCP"
	BLOB_STORAGE_MINIO                 = "MINIO"
)
