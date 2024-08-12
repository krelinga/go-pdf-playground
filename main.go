package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/go-pdf/fpdf"

    demo_blankpage "github.com/krelinga/go-pdf-playground/internal/demos/blankpage"
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
}

func runDemos(demos []*demo, outputDir string) error {
    // Prep the output directory.
    if err := os.RemoveAll(outputDir); err != nil {
        return fmt.Errorf("error removing %s: %w", outputDir, err)
    }
    if err := os.MkdirAll(outputDir, 755); err != nil {
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
