package main

import (
  "fmt"
  "github.com/jung-kurt/gofpdf"
)

func main() {
  pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
  w, h := pdf.GetPageSize()
  // we print the width and height to not get an error
  fmt.Printf("width=%v", "height=%v\n", w, h)
  pdf.AddPage()

  pdf.MoveTo(0, 0)
  pdf.SetFont("arial", "B", 30)
  // to get exact fontsize
  _, lineHt := pdf.GetFontSize()
  pdf.SetTextColor(255, 0, 0)
  pdf.Text(0, lineHt, "Hello, world")
  pdf.MoveTo(0, lineHt*2.0)


  pdf.SetFont("times", "", 18)
  pdf.SetTextColor(100, 100, 100)
  _, lineHt = pdf.GetFontSize()
  pdf.MultiCell(0, lineHt, "Foodonate would like to extend the highest appreciation, on behalf of every member of the team, to", gofpdf.BorderNone, gofpdf.AlignRight, false)

  // Shapes
  pdf.SetFillColor(0, 255, 0)
  pdf.SetDrawColor(0, 0, 255)
  pdf.Rect(10, 100, 100, 100, "FD")
  pdf.SetFillColor(100, 200, 200)
  pdf.Polygon([]gofpdf.PointType{
    {110, 250},
    {160, 300},
    {110, 350},
    {600, 300},
  }, "F")

  err := pdf.OutputFileAndClose("p1.pdf")
  if err != nil {
    panic(err)
  }
}
