package fontwidth
import "github.com/go-pdf/fpdf"

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    p.SetFont("Arial", "", 12)
    text := "Some Text"
    // 50, 50 is at the bottom left of the text.
    p.Text(50, 50, text)
    w := p.GetStringWidth(text)
    // Should be more-or-less directly under the text.
    p.Line(50, 50 + 2, 50 + w, 50 + 2)

    // This demonstrates that text just overflows a cell that is too small.
    p.Text(50, 100, "Cell that is too small")
    p.MoveTo(50, 105)
    p.Cell(50, 14, "Here is some very very very long text.")

    return p
}
