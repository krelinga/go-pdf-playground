package helloworld

import "github.com/go-pdf/fpdf"

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    p.SetFont("Arial", "B", 16)
    p.Cell(72, 100, "Hello, world")
    return p
}
