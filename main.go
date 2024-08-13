package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/go-pdf/fpdf"

    demo_blankpage "github.com/krelinga/go-pdf-playground/internal/demos/blankpage"
    demo_helloworld "github.com/krelinga/go-pdf-playground/internal/demos/helloworld"
    demo_shapes "github.com/krelinga/go-pdf-playground/internal/demos/shapes"
    demo_templates "github.com/krelinga/go-pdf-playground/internal/demos/templates"
    demo_fontwidth "github.com/krelinga/go-pdf-playground/internal/demos/fontwidth"
    demo_templatecompose "github.com/krelinga/go-pdf-playground/internal/demos/templatecompose"
)

const outputDir = "/host/test/go-pdf-playground"

type demo struct {
    Name string
    Run func() *fpdf.Fpdf
}

var demos = []*demo{
    &demo{
        Name: "blankpage",
        Run: demo_blankpage.Run,
    },
    &demo{
        Name: "helloworld",
        Run: demo_helloworld.Run,
    },
    &demo{
        Name: "shapes",
        Run: demo_shapes.Run,
    },
    &demo{
        Name: "templates",
        Run: demo_templates.Run,
    },
    &demo{
        Name: "fontwidth",
        Run: demo_fontwidth.Run,
    },
    &demo{
        Name: "templatecompose",
        Run: demo_templatecompose.Run,
    },
}

func runDemos(demos []*demo, outputDir string) error {
    // Prep the output directory.
    if err := os.RemoveAll(outputDir); err != nil {
        return fmt.Errorf("error removing %s: %w", outputDir, err)
    }
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        return fmt.Errorf("error creating %s: %w", outputDir, err)
    }

    // Run each of the demos.
    for _, d := range demos {
        p := d.Run()
        outPath := filepath.Join(outputDir, d.Name) + ".pdf"
        if err := p.OutputFileAndClose(outPath); err != nil {
            return fmt.Errorf("Error generating file %s: %w", outPath, err)
        }
    }

    return nil
}

func main() {
    if err := runDemos(demos, outputDir); err != nil {
        log.Fatalln(err)
    }
}
