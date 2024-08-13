package templatecompose
import (
    "math"

    "github.com/go-pdf/fpdf"
)

func addAt(p *fpdf.Fpdf, x, y float64, fn func(*fpdf.Tpl)) {
    t := p.CreateTemplate(fn)
    origin := fpdf.PointType{X: x, Y: y}
    _, size := t.Size()
    p.UseTemplateScaled(t, origin, size)
}

func smallSquare(t *fpdf.Tpl) {
    t.SetDrawColor(0, 0, 0)
    t.SetLineWidth(1)
    t.Rect(1, 1, 6, 6, "D")
}

func tpl2fpdf(t *fpdf.Tpl) *fpdf.Fpdf {
    return &t.Fpdf
}

func squareOfSquares(t *fpdf.Tpl) {
    for row := 0; row < 5; row++ {
        for col := 0; col < 5; col++ {
            x := float64(col) * float64(7.5)
            y := float64(row) * float64(7.5)
            addAt(tpl2fpdf(t), x, y, smallSquare)
        }
    }
}

func templateSize(t fpdf.Template) (width, height float64) {
    _, s := t.Size()
    return s.Wd, s.Ht
}

func borderTopLabel(label string, child func(*fpdf.Tpl)) func(*fpdf.Tpl) {
    return func(t *fpdf.Tpl) {
        const fontHeight float64 = 10
        const pad float64 = 2
        t.SetFont("Arial", "", fontHeight)
        childT := t.CreateTemplate(child)
        childTWidth, childTHeight := templateSize(childT)

        bodyWidth := math.Max(childTWidth, t.GetStringWidth(label))
        bodyHeight := fontHeight + childTHeight + float64(3) * pad

        t.RoundedRect(1, 1, 1 + bodyWidth, 1 + bodyHeight, 2, "1234", "D")
        t.Text(1 + pad, 1 + pad + fontHeight, label)
        t.UseTemplateScaled(childT, fpdf.PointType{X: 1 + pad, Y: 1 + 2 * pad + fontHeight}, fpdf.SizeType{Wd: childTWidth, Ht: childTHeight})
    }
}

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    addAt(p, 50, 50, borderTopLabel("Hull", squareOfSquares))
    return p
}
