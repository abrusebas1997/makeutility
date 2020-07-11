package main

import "github.com/jung-kurt/gofpdf"

func main() {
  pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
  w, h := pdf.GetPageSize()
  // we print the width and height to not get an error
  fmt.Printf("width=%v", "height=%v\n", w, h)
  pdf.AddPage()



  pdf.AddPage()
  pdf.SetFont("arial", "B", 38)
  pdf.Cell(40, 10, "Hello, world")
  err := pdf.OutputFileAndClose("p1.pdf")
  if err != nil {
    panic(err)
  }
}
