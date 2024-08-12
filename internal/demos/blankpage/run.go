package blankpage

import "github.com/go-pdf/fpdf"

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    return p
}
