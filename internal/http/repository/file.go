package repository

import (
	"fmt"
	"mime/multipart"

	"github.com/certified-juniors/AtomHackEarthBackend/internal/model"
)

func (r *Repository) GetFilesByDocumentID(docID uint) ([]model.File, error) {
	var files []model.File
	if err := r.db.DatabaseGORM.Where("document_id = ?", docID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *Repository) UploadFiles(docID uint, files []*multipart.FileHeader) ([]uint, error) {
	var fileIDs []uint

	for _, fileHeader := range files {
		// Получаем файл из заголовка запроса
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Получаем размер файла
		fileSize := fileHeader.Size
		fileName := fileHeader.Filename

		// Загружаем файл в хранилище MinIO
		objectName := fmt.Sprintf("documents/%d/%s", docID, fileName)
		fileURL, err := r.mc.UploadFile(objectName, file, fileSize)
		if err != nil {
			return nil, err
		}

		// Сохраняем запись файла в базе данных
		newFile := model.File{
			Path:       fileURL,
			DocumentID: docID,
		}
		if err := r.db.DatabaseGORM.Create(&newFile).Error; err != nil {
			return nil, fmt.Errorf("failed to save file path in database: %w", err)
		}

		// Добавляем ID файла в список
		fileIDs = append(fileIDs, newFile.ID)
	}

	return fileIDs, nil
}

func (r *Repository) UploadFile(docID uint, file multipart.File, fileSize int64, fileName string) (uint, error) {
	var document model.Document
	if err := r.db.DatabaseGORM.First(&document, docID).Error; err != nil {
		return 0, fmt.Errorf("failed to find document with ID %d: %w", docID, err)
	}

	objectName := fmt.Sprintf("documents/%d/%s", docID, fileName)

	fileURL, err := r.mc.UploadFile(objectName, file, fileSize)
	if err != nil {
		return 0, err
	}

	newFile := model.File{
		Path:       fileURL,
		DocumentID: docID,
	}
	// Сохраняем новую запись файла в базе данных
	if err := r.db.DatabaseGORM.Create(&newFile).Error; err != nil {
		return 0, fmt.Errorf("failed to save file path in database: %w", err)
	}

	return newFile.ID, nil
}

