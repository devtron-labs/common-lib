package blob_storage

type BlobStorageRequest struct {
	StorageType         BlobStorageType
	SourceKey           string
	DestinationKey      string
	AwsS3BaseConfig     *AwsS3BaseConfig
	AzureBlobBaseConfig *AzureBlobBaseConfig
	GcpBlobBaseConfig   *GcpBlobBaseConfig
}

type BlobStorageS3Config struct {
	AccessKey                  string `json:"accessKey"`
	Passkey                    string `json:"passkey"`
	EndpointUrl                string `json:"endpointUrl"`
	IsInSecure                 bool   `json:"isInSecure"`
	CiLogBucketName            string `json:"ciLogBucketName"`
	CiLogRegion                string `json:"ciLogRegion"`
	CiLogBucketVersioning      bool   `json:"ciLogBucketVersioning"`
	CiCacheBucketName          string `json:"ciCacheBucketName"`
	CiCacheRegion              string `json:"ciCacheRegion"`
	CiCacheBucketVersioning    bool   `json:"ciCacheBucketVersioning"`
	CiArtifactBucketName       string `json:"ciArtifactBucketName"`
	CiArtifactRegion           string `json:"ciArtifactRegion"`
	CiArtifactBucketVersioning bool   `json:"ciArtifactBucketVersioning"`
}

func (b *BlobStorageS3Config) GetBlobStorageBaseS3Config(blobStorageObjectType string) *AwsS3BaseConfig {
	awsS3BaseConfig := &AwsS3BaseConfig{
		AccessKey:         b.AccessKey,
		Passkey:           b.Passkey,
		EndpointUrl:       b.EndpointUrl,
		IsInSecure:        b.IsInSecure,
	}
	switch blobStorageObjectType {
	case BlobStorageObjectTypeCache:
		awsS3BaseConfig.BucketName = b.CiCacheBucketName
		awsS3BaseConfig.Region = b.CiCacheRegion
		awsS3BaseConfig.VersioningEnabled = b.CiCacheBucketVersioning
		return awsS3BaseConfig
	case BlobStorageObjectTypeLog:
		awsS3BaseConfig.BucketName = b.CiLogBucketName
		awsS3BaseConfig.Region = b.CiLogRegion
		awsS3BaseConfig.VersioningEnabled = b.CiLogBucketVersioning
		return awsS3BaseConfig
	case BlobStorageObjectTypeArtifact:
		awsS3BaseConfig.BucketName = b.CiArtifactBucketName
		awsS3BaseConfig.Region = b.CiArtifactRegion
		awsS3BaseConfig.VersioningEnabled = b.CiArtifactBucketVersioning
		return awsS3BaseConfig
	default:
		return nil
	}
}

type AwsS3BaseConfig struct {
	AccessKey         string `json:"accessKey"`
	Passkey           string `json:"passkey"`
	EndpointUrl       string `json:"endpointUrl"`
	IsInSecure        bool   `json:"isInSecure"`
	BucketName        string `json:"bucketName"`
	Region            string `json:"region"`
	VersioningEnabled bool   `json:"versioningEnabled"`
}

type AzureBlobConfig struct {
	Enabled               bool   `json:"enabled"`
	AccountName           string `json:"accountName"`
	BlobContainerCiLog    string `json:"blobContainerCiLog"`
	BlobContainerCiCache  string `json:"blobContainerCiCache"`
	BlobContainerArtifact string `json:"blobStorageArtifact"`
	AccountKey            string `json:"accountKey"`
}

func (b *AzureBlobConfig) GetBlobStorageBaseAzureConfig(blobStorageObjectType string) *AzureBlobBaseConfig {
	azureBlobBaseConfig := &AzureBlobBaseConfig{
		Enabled:     b.Enabled,
		AccountName: b.AccountName,
		AccountKey:  b.AccountKey,
	}
	switch blobStorageObjectType {
	case BlobStorageObjectTypeCache:
		azureBlobBaseConfig.BlobContainerName = b.BlobContainerCiCache
		return azureBlobBaseConfig
	case BlobStorageObjectTypeLog:
		azureBlobBaseConfig.BlobContainerName = b.BlobContainerCiLog
		return azureBlobBaseConfig
	case BlobStorageObjectTypeArtifact:
		azureBlobBaseConfig.BlobContainerName = b.BlobContainerArtifact
		return azureBlobBaseConfig
	default:
		return nil
	}
}

type AzureBlobBaseConfig struct {
	Enabled           bool   `json:"enabled"`
	AccountName       string `json:"accountName"`
	AccountKey        string `json:"accountKey"`
	BlobContainerName string `json:"blobContainerName"`
}

type GcpBlobConfig struct {
	CredentialFileJsonData string `json:"credentialFileData"`
	CacheBucketName        string `json:"ciCacheBucketName"`
	LogBucketName          string `json:"logBucketName"`
	ArtifactBucketName     string `json:"artifactBucketName"`
}

func (b *GcpBlobConfig) GetBlobStorageBaseGcpConfig(blobStorageObjectType string) *GcpBlobBaseConfig {
	gcpBlobBaseConfig := &GcpBlobBaseConfig{
		CredentialFileJsonData: b.CredentialFileJsonData
	}
	switch blobStorageObjectType {
	case BlobStorageObjectTypeCache:
		gcpBlobBaseConfig.BucketName = b.CacheBucketName
		return gcpBlobBaseConfig
	case BlobStorageObjectTypeLog:
		gcpBlobBaseConfig.BucketName = b.LogBucketName
		return gcpBlobBaseConfig
	case BlobStorageObjectTypeArtifact:
		gcpBlobBaseConfig.BucketName = b.ArtifactBucketName
		return gcpBlobBaseConfig
	default:
		return nil
	}
}

type GcpBlobBaseConfig struct {
	BucketName             string `json:"bucketName"`
	CredentialFileJsonData string `json:"credentialFileData"`
}

type BlobStorageType string

const (
	BLOB_STORAGE_AZURE            BlobStorageType = "AZURE"
	BLOB_STORAGE_S3                               = "S3"
	BLOB_STORAGE_GCP                              = "GCP"
	BLOB_STORAGE_MINIO                            = "MINIO"
	BlobStorageObjectTypeCache                    = "cache"
	BlobStorageObjectTypeArtifact                 = "artifact"
	BlobStorageObjectTypeLog                      = "log"
)
