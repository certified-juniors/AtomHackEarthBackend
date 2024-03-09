package repository

import (
	"github.com/certified-juniors/AtomHackEarthBackend/internal/model"
)

func (r *Repository) GetFormedDocuments(page, pageSize int) ([]model.Document, error) {
    var documents []model.Document
    offset := (page - 1) * pageSize

    if err := r.db.DatabaseGORM.Where("status = ?", model.StatusFormed).Order("sent_time DESC").Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
        return nil, err
    }

    return documents, nil
}

func (r *Repository) GetDocumentByID(docID uint) (*model.Document, error) {
    var document model.Document
    if err := r.db.DatabaseGORM.Preload("Files").Where("status != ?", model.StatusDeleted).First(&document, docID).Error; err != nil {
        return nil, err
    }
    return &document, nil
}

func (r *Repository) CreateDocument(doc *model.Document) (uint, error) {
	if err := r.db.DatabaseGORM.Create(doc).Error; err != nil {
		return 0, err
	}

	return doc.ID, nil
}


