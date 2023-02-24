package services

import (
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

type CheckStorage interface {
	Get(id string) entity.ICheck
	GetAll() []entity.ICheck
	Create(check entity.ICheck) error
}

type PDFStorage interface {
	GenerateCheckPDF(check entity.ICheck) error
}

type checkService struct {
	checkStorage CheckStorage
	pdfStorage   PDFStorage
}

func NewCheckService(checkStorage CheckStorage, pdfStorage PDFStorage) *checkService {
	return &checkService{checkStorage: checkStorage, pdfStorage: pdfStorage}
}
