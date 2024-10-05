/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
			SourceKey:         "/Users/kripanshbanga/Desktop/out_self_cm.yaml",
			DestinationKey:    "abcd/efg/out_self_cm.yaml",
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
			SourceKey:         "abcd/efg/out_self_cm.yaml",
			DestinationKey:    "out_self_cm.yaml",
			GcpBlobBaseConfig: gcpConfig,
		}
		success, totalBytes, err := blobStorageServiceImpl.Get(request)
		fmt.Println(success, totalBytes, err)
	})

	t.Run("S3 Upload with session", func(t *testing.T) {
		awsS3BaseConfig := &AwsS3BaseConfig{
			AccessKey:   "",
			Passkey:     "",
			EndpointUrl: "",
			BucketName:  "deepak-bucket1234",
			Region:      "ap-south-1",
		}
		blobStorageServiceImpl := NewBlobStorageServiceImpl(nil)
		request := &BlobStorageRequest{
			StorageType:     BLOB_STORAGE_S3,
			SourceKey:       "/shivamnagar409/latest.txt",
			DestinationKey:  "/latest.txt",
			AwsS3BaseConfig: awsS3BaseConfig,
		}
		err := blobStorageServiceImpl.UploadToBlobWithSession(request)
		fmt.Println(err)
	})

}
