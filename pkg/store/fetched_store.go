package store

import (
	"fmt"
	"log"
	"strings"
	"time"

	"job-seek/pkg/request"

	excelize "github.com/xuri/excelize/v2"
)

var (
	defaultStorePath = "./outputs/fetched_job.xlsx"
	headerRow        = map[string]string{
		"A": "PostId",
		"B": "PostTitle",
		"C": "PayRange",
		"D": "PostUrl",
		"E": "Post-Role",
		"F": "Expire-Date",
		"G": "Company-Name",
		"H": "Company-Industry",
		"I": "Company-GroupSize",
		"J": "Company-JobPosted",
		"K": "Company-SeekReferenceId",
		"L": "Company-Url",
		"M": "Company-Linkedin",
		"N": "Company-Description",
		"O": "Company-Specialties",
		"P": "Company-HeadQuarters",
		"Q": "Company-Location",
		"R": "Post-HittedKeywords",
		"S": "Post-Score",
		"T": "Post-DebugText",
	}
)

func CreateOrGetStoredFile(storePath string) *excelize.File {
	f, err := excelize.OpenFile(storePath)
	if err != nil {
		log.Println(err)
		f = excelize.NewFile()
		f.SaveAs(storePath)
	}
	return f

}

func CreateRecordSheet(f *excelize.File, sheetName string) {
	index, _ := f.GetSheetIndex(sheetName)
	if index == -1 {
		index, _ = f.NewSheet(sheetName)
	}
	f.SetActiveSheet(index)
	f.SetSheetView(sheetName, -1, nil)
}

func CreateFetchedJobStore() (*excelize.File, string, int) {

	storePath := defaultStorePath
	fetchedDate := time.Now().Format("2006-01-02T1504")
	sheetName := fmt.Sprintf("fetched_job_%s", fetchedDate)
	f := CreateOrGetStoredFile(storePath)
	CreateRecordSheet(f, sheetName)
	if idx, _ := f.GetSheetIndex("Sheet1"); idx != -1 {
		f.DeleteSheet("Sheet1")
	}

	initIndex := WriteRecordHeader(f, sheetName)
	if err := f.Save(); err != nil {
		log.Println(err)
	}

	return f, sheetName, initIndex
}

func WriteRecordHeader(f *excelize.File, sheetName string) int {

	intialIndex := 2
	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "top",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	defaultStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "top",
			WrapText:   true,
		},
	})

	f.SetColStyle(sheetName, "A:AA", defaultStyle)
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", intialIndex), fmt.Sprintf("AA%d", intialIndex), style)

	// headerRow := []string{"PostId", "PostTitle", "PayRange", "PostUrl", "Post-Role", "Expire-Date", "Company-Name", "Company-Industry", "Company-GroupSize", "Company-JobPosted", "Company-SeekReferenceId", "Company-Url", "Company-Linkedin", "Company-Description", "Company-Specialties", "Company-HeadQuarters", "Company-Location", "Post-HittedKeywords", "Post-Score", "Post-DebugText"}

	for key, header := range headerRow {
		cellIndex := fmt.Sprintf("%s%d", key, intialIndex)
		f.SetCellValue(sheetName, cellIndex, header)

	}

	f.SetColWidth(sheetName, "A", "S", 18)
	f.SetColWidth(sheetName, "T", "U", 500)

	return intialIndex

}

func WriteRecordToPage(f *excelize.File, sheetName string, record *request.SeekPostDetails, index int) {
	mapping := map[string]string{
		"PostId":                  record.PostId,
		"PostTitle":               record.PostTitle,
		"PayRange":                record.PayRange,
		"PostUrl":                 record.PostUrl,
		"Post-Role":               fmt.Sprintf("%s, %s", record.Role, record.WorkType),
		"Expire-Date":             record.ExpiringDate,
		"Company-Name":            record.CompanyDetails.Name,
		"Company-Industry":        record.CompanyDetails.Industry,
		"Company-GroupSize":       record.CompanyDetails.GroupSize,
		"Company-JobPosted":       fmt.Sprintf("%d", record.CompanyDetails.JobPosted),
		"Company-SeekReferenceId": record.CompanyDetails.ReferenceId,
		"Company-Url":             record.CompanyDetails.Url,
		"Company-Linkedin":        record.CompanyDetails.Linkedin,
		"Company-Description":     record.CompanyDetails.Description,
		"Company-Specialties":     record.CompanyDetails.Specialties,
		"Company-HeadQuarters":    record.CompanyDetails.HeadQuarters,
		"Company-Location":        record.CompanyDetails.Locations,
		"Post-HittedKeywords":     strings.Join(record.HittedKeywords, ","),
		"Post-Score":              fmt.Sprintf("%d", record.Score),
		"Post-DebugText":          record.DebugText,
	}

	for key, value := range headerRow {
		// fmt.Println(key, value, mapping[value])
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", key, index), mapping[value])
	}

}

func Save(f *excelize.File) error {
	return f.SaveAs(defaultStorePath)
}
