package services

import (
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

type CheckStorage interface {
	Get(id int) (entity.OrderCheck, error)
	GetAll() []entity.OrderCheck
	Create(check entity.OrderCheck) error
	GetAllGeneratedChecks() ([]entity.OrderCheck, error)
	UpdateStatusPrinted(checkIds []int) error
}

type PDFStorage interface {
	GenerateCheckPDF(check entity.OrderCheck) error
}

type checkService struct {
	checkStorage CheckStorage
	pdfStorage   PDFStorage
}

func NewCheckService(checkStorage CheckStorage, pdfStorage PDFStorage) *checkService {
	return &checkService{checkStorage: checkStorage, pdfStorage: pdfStorage}
}
