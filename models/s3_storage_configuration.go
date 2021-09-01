// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3StorageConfiguration s3 storage configuration
//
// swagger:model s3-storage-configuration
type S3StorageConfiguration struct {

	// aws access key id
	AwsAccessKeyID string `json:"aws_access_key_id,omitempty"`

	// aws secret access key
	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty"`

	// bucket host
	BucketHost string `json:"bucket_host,omitempty"`

	// bucket name
	BucketName string `json:"bucket_name,omitempty"`

	// bucket port
	BucketPort int32 `json:"bucket_port,omitempty"`
}

// Validate validates this s3 storage configuration
func (m *S3StorageConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3StorageConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3StorageConfiguration) UnmarshalBinary(b []byte) error {
	var res S3StorageConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
