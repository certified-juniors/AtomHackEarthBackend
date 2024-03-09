package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/certified-juniors/AtomHackEarthBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// GetFormedDocuments возвращает сформированные документы.
// @Summary Возвращает сформированные документы.
// @Description Возвращает список сформированных документов с учетом параметров page и pageSize.
// @Tags Документы
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param pageSize query int false "Размер страницы" default(10)
// @Success 200 {array} model.Document "Успешный ответ"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/formed [get]
func (h *Handler) GetFormedDocuments(c *gin.Context) {
    // Получаем параметры из URL
    page, _ := strconv.Atoi(c.Query("page"))
    pageSize, _ := strconv.Atoi(c.Query("pageSize"))

    // Получаем сформированные документы
    documents, err := h.r.GetFormedDocuments(page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve documents: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, documents)
}

// GetDocumentByID получает документ по его ID.
// @Summary Получает документ по ID.
// @Description Получает документ из репозитория по указанному ID.
// @Tags Документы
// @Accept json
// @Produce json
// @Param docID path int true "ID документа"
// @Success 200 {object} model.Document "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/{docID} [get]
func (h *Handler) GetDocumentByID(c *gin.Context) {
	docID, err := strconv.Atoi(c.Param("docID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return
	}

	doc, err := h.r.GetDocumentByID(uint(docID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doc)
}

// AcceptDocument принимает новый документ.
// @Summary Принимает новый документ.
// @Description Принимает новый документ с параметрами id, title, owner, createdAt, payload и files.
// @Tags Документы
// @Accept json
// @Produce json
// @Param id formData int true "ID документа"
// @Param title formData string true "Заголовок документа"
// @Param owner formData string true "Владелец документа"
// @Param createdAt formData string true "Дата и время создания документа в формате RFC3339"
// @Param payload formData string true "Payload документа"
// @Param files formData file true "Файлы, прикрепленные к документу"
// @Success 200 {object} model.AcceptDocument "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/send-to-earth [post]
func (h *Handler) AcceptDocument(c *gin.Context) {
	acceptID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get id from request"})
		return
	}

	title := c.PostForm("title")
	owner := c.PostForm("owner")

	createdAtStr := c.PostForm("createdAt")
	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse createdAt"})
		return
	}

	payload := c.PostForm("payload")

	receivedTime := time.Now()
	deliveryStatus := model.DeliveryStatusSuccess

	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) > 0 {
		// Файлы присутствуют, выполняем загрузку файлов
		doc := model.Document{
			AcceptID:        uint(acceptID),
			Title:           title,
			Owner:           owner,
			ReceivedTime:    &receivedTime,
			Status:          model.StatusFormed,
			CreatedAt:       createdAt,
			DeliveryStatus:  &deliveryStatus,
			Payload:         payload,
		}

		docID, err := h.r.CreateDocument(&doc)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fileIDs, err := h.r.UploadFiles(uint(docID), files)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload files: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"fileIDs": fileIDs})
	} else {
		doc := model.Document{
			AcceptID:        uint(acceptID),
			Title:           title,
			Owner:           owner,
			ReceivedTime:    &receivedTime,
			Status:          model.StatusFormed,
			CreatedAt:       createdAt,
			DeliveryStatus:  &deliveryStatus,
			Payload:         payload,
		}
		docID, err := h.r.CreateDocument(&doc)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"docID": docID})
	}
}




