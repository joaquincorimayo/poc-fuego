package adapter

import (
	"github.com/joaquincorimayo/poc-fuego/domain/model"
	"github.com/joaquincorimayo/poc-fuego/domain/ports"
)

type ReportService struct {
	reportRepository ports.Repository
}

func NewReportService(reportRepository ports.Repository) *ReportService {
	return &ReportService{
		reportRepository: reportRepository,
	}
}

func (s *ReportService) FindByCode(code string) (*model.Report, error) {
	return s.reportRepository.FindByCode(code)
}

func (s *ReportService) Save(report *model.Report) error {
	return s.reportRepository.Save(report)
}
