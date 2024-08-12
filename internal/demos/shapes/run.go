package shapes
import "github.com/go-pdf/fpdf"

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    p.SetFillColor(255, 0, 0)
    p.SetDrawColor(0, 0, 255)
    p.SetLineWidth(5)
    p.Circle(100, 100, 25, "DF")


    p.SetFillColor(50, 200, 100)
    p.SetDrawColor(0, 0, 0)
    p.SetLineWidth(2)
    p.RoundedRect(50, 200, 100, 50, 10, "1234", "DF")


    isoTriangle := func(cx, cy, w, h float64, style string) {
        hr := h / 2.0
        wr := w / 2.0
        p.MoveTo(cx, cy - hr)
        p.LineTo(cx + wr, cy + hr)
        p.LineTo(cx - wr, cy + hr)
        p.ClosePath()
        p.DrawPath(style)
    }

    p.SetFillColor(100, 255, 255)
    p.SetDrawColor(0, 0, 0)
    p.SetLineWidth(2)
    isoTriangle(100, 400, 50, 75, "DF")
    return p
}
