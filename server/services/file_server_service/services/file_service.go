package services

import (
	"file_server_service/models"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/jung-kurt/gofpdf"
	// "github.com/unidoc/unioffice/document"
)

// Initialize a validator
var validate = validator.New()

// ValidateLetterDocument validates the LetterDocument struct
func ValidateLetterDocument(letter models.LetterDocument) error {
	return validate.Struct(letter)
}

// using unioffice
// GenerateDocx creates a .docx file from the LetterDocument data
// func GenerateDocxOld(letter models.LetterDocument) (string, error) {
// 	// Create a new document
// 	doc := document.New()
// 	defer doc.Close()

// 	// Add content to the document
// 	doc.AddParagraph().AddRun().AddText("Subject: " + letter.Subject)
// 	doc.AddParagraph().AddRun().AddText("Receiver Name: " + letter.ReceiverName)
// 	doc.AddParagraph().AddRun().AddText("Content: " + letter.Content)
// 	doc.AddParagraph().AddRun().AddText("Sender Name: " + letter.SenderName)

// 	fileName := letter.FileName
// 	filePath := filepath.Join("public", "docs", fileName)

// 	// Save the document
// 	if err := doc.SaveToFile(filePath); err != nil {
// 		return "", err
// 	}

// 	return fileName, nil
// }

// generate document using gopdf
func GenerateDocx(letter models.LetterDocument) (string, error) {
	// Create a new PDF document (A4 page size)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Add content to the PDF
	pdf.Cell(0, 10, "Subject: "+letter.Subject)
	pdf.Ln(10) // Line break

	pdf.SetFont("Arial", "", 14)
	pdf.Cell(0, 10, "Receiver Name: "+letter.ReceiverName)
	pdf.Ln(10) // Line break

	pdf.Cell(0, 10, "Content: "+letter.Content)
	pdf.Ln(10) // Line break

	pdf.Cell(0, 10, "Sender Name: "+letter.SenderName)
	pdf.Ln(10) // Line break

	// Define file name and path
	fileName := letter.FileName
	filePath := filepath.Join("public", "docs", fileName)

	// Save the PDF document
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
