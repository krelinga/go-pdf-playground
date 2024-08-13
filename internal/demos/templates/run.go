package templates
import "github.com/go-pdf/fpdf"

func point(x, y float64) fpdf.PointType {
    return fpdf.PointType{
        X: 0,
        Y: 0,
    }
}

var origin = point(0, 0)

func size(w, h float64) fpdf.SizeType {
    return fpdf.SizeType{
        Wd: w,
        Ht: h,
    }
}

func inches(i float64) float64 {
    return 72.0 * i
}

var rectTemplateSize = size(inches(3), inches(3))

func rectTemplate(p *fpdf.Fpdf) fpdf.Template {
    return p.CreateTemplate(func(t *fpdf.Tpl) {
        t.SetFillColor(50, 200, 100)
        t.SetDrawColor(0, 0, 0)
        t.SetLineWidth(2)
        t.RoundedRect(1, 1, 100, 50, 10, "1234", "DF")
    })
}

func Run() *fpdf.Fpdf {
    p := fpdf.New("P", "pt", "Letter", "")
    p.AddPage()
    p.SetFillColor(0, 0, 0)
    rect := rectTemplate(p)
    p.UseTemplate(rect)
    point, size := rect.Size()
    // I couldn't get UseTemplateScaled() to work with a template that is not
    // copuled to this instance of PDF (i.e. one created with fpdf.CreateTpl()).
    p.UseTemplateScaled(rect, point.Transform(25, 25), size)

    // If this is black (and it is) then it proves that templates do save &
    // restore the PDF state, so I can use these as a sort of push / pop
    // equivalent.
    p.Circle(50, 500, 10, "F")
    return p
}
