package repository

import (
	"strings"

	"github.com/certified-juniors/AtomHackEarthBackend/internal/model"
)

func (r *Repository) GetDocumentsFormedCount(status model.Status, deliveryStatus model.DeliveryStatus, ownerOrTitle string) (uint, error) {
    var count int64
    query := r.db.DatabaseGORM.Model(&model.Document{})

    if status != "" {
        query = query.Where("LOWER(status) = LOWER(?)", status)
    }

    if deliveryStatus != "" {
        query = query.Where("LOWER(delivery_status) = LOWER(?)", deliveryStatus)
    }

    if ownerOrTitle != "" {
        query = query.Where("LOWER(owner) LIKE ? OR LOWER(title) LIKE ?", "%"+strings.ToLower(ownerOrTitle)+"%", "%"+strings.ToLower(ownerOrTitle)+"%")
    }

    if err := query.Count(&count).Error; err != nil {
        return 0, err
    }

    return uint(count), nil
}

func (r *Repository) GetFormedDocuments(page, pageSize int, deliveryStatus model.DeliveryStatus, ownerOrTitle string) ([]model.Document, uint, error) {
    var documents []model.Document
    offset := (page - 1) * pageSize

    query := r.db.DatabaseGORM

    if deliveryStatus != "" {
        query = query.Where("LOWER(delivery_status) = LOWER(?)", deliveryStatus)

        if deliveryStatus == model.DeliveryStatusSuccess {
            query = query.Order("received_time DESC")
        } else if deliveryStatus == model.DeliveryStatusPending {
            query = query.Order("sent_time DESC")
        }
    } else {
        query = query.Where("LOWER(status) = LOWER(?)", model.StatusFormed).Order("sent_time DESC")
    }

    if ownerOrTitle != "" {
        ownerOrTitle = strings.ToLower(ownerOrTitle)
        query = query.Where("LOWER(owner) LIKE ? OR LOWER(title) LIKE ?", "%"+ownerOrTitle+"%", "%"+ownerOrTitle+"%")
    }

    if err := query.Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
        return nil, 0, err
    }

    total, err := r.GetDocumentsFormedCount(model.StatusFormed, deliveryStatus, ownerOrTitle)
    if err != nil {
        return nil, 0, err
    }

    return documents, total, nil
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


