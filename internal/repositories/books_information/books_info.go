package booksinformation

import (
	"context"
	"database/sql"

	"fanchann/library/internal/models/domain"
)

type IBooksInformation interface {
	Insert(ctx context.Context, tx *sql.Tx, dataBooks *domain.Books_Information) domain.Books_Information
	Delete(ctx context.Context, tx *sql.Tx, dataId *domain.Books_Information) error
}
