package images

import (
	"github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/core/images"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type Image struct {
	ClientImage client.Image
	Config      ocispec.Image
	Image       images.Image
	Manifest    *ocispec.Manifest
}
