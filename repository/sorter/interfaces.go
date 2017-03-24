package sorter

import (
	"github.com/dpb587/metalink/repository"
)

type Sorter interface {
	Less(repository.BlobReceipt, repository.BlobReceipt) bool
}
