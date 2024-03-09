package repository

import (
	"github.com/certified-juniors/AtomHackEarthBackend/internal/model"
)

func (r *Repository) GetDocumentsCountByStatus(status model.Status) (uint, error) {
    var count int64
    if err := r.db.DatabaseGORM.Model(&model.Document{}).Where("status = ?", status).Count(&count).Error; err != nil {
        return 0, err
    }
    return uint(count), nil
}

func (r *Repository) GetDocumentsCountByDeliveryStatus(deliveryStatus model.DeliveryStatus) (uint, error) {
    var count int64
    if err := r.db.DatabaseGORM.Model(&model.Document{}).Where("delivery_status = ?", deliveryStatus).Count(&count).Error; err != nil {
        return 0, err
    }
    return uint(count), nil
}

func (r *Repository) GetFormedDocuments(page, pageSize int, deliveryStatus model.DeliveryStatus) ([]model.Document, uint, error) {
    var documents []model.Document
    offset := (page - 1) * pageSize

    if deliveryStatus != "" {
        if deliveryStatus == model.DeliveryStatusSuccess{
            if err := r.db.DatabaseGORM.Where("delivery_status = ?", deliveryStatus).Order("received_time DESC").Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
                return nil, 0, err
            }
        } else if deliveryStatus == model.DeliveryStatusPending{
            if err := r.db.DatabaseGORM.Where("delivery_status = ?", deliveryStatus).Order("sent_time DESC").Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
                return nil, 0, err
            }
        }

        total, err := r.GetDocumentsCountByDeliveryStatus(deliveryStatus)
        if err != nil {
            return nil, 0, err
        }

        return documents, total, nil
    }

    if err := r.db.DatabaseGORM.Where("status = ?", model.StatusFormed).Order("sent_time DESC").Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
        return nil, 0, err
    }

    total, err := r.GetDocumentsCountByStatus(model.StatusFormed)
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


