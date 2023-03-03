package services

import (
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"go.uber.org/zap"
)

type CheckStorage interface {
	Get(id int) (entity.OrderCheck, error)
	GetAll() []entity.OrderCheck
	Create(check entity.OrderCheck) (int, error)
	GetAllGeneratedChecks() ([]entity.OrderCheck, error)
	UpdateStatusPrinted(checkIds []int) error
	UpdateStatusGeneratedAndFilePath(checkId int, filePath string) error
}

type PDFStorage interface {
	GenerateCheckPDF(check entity.OrderCheck) (string, error)
}

type checkService struct {
	log          *zap.SugaredLogger
	checkStorage CheckStorage
	pdfStorage   PDFStorage
}

func NewCheckService(log *zap.SugaredLogger, checkStorage CheckStorage, pdfStorage PDFStorage) *checkService {
	return &checkService{log: log, checkStorage: checkStorage, pdfStorage: pdfStorage}
}
