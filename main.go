package main

import (
  "fmt"
  "github.com/jung-kurt/gofpdf"
  "image/color"
  "flag"
  "time"
)

type PDFOption func(*gofpdf.Fpdf)
// returns PDFOption
func FillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}

// helper function
// without this, we would need to do something like
// 	pdf.SetFillColor(c.R, c.G, c.B)
func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}
type PDF struct {
  fpdf *gofpdf.Fpdf
  x, y float64
}
// WAYS TO MOVE
// regular move
func (p *PDF) Move(xDelta, yDelta float64) {
	p.x, p.y = p.x+xDelta, p.y+yDelta
	p.fpdf.MoveTo(p.x, p.y)
}

// absolut move
func (p *PDF) MoveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) Text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

// takes a slice  of gofpdf
func (p *PDF) Polygon(pts []gofpdf.PointType, opts ...PDFOption) {

	for _, opt := range opts {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {

  name := flag.String("name", "Sebastian", "the name of the person who completed more than 40 hours of community service")
  flag.Parse()

  fpdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
  w, h := fpdf.GetPageSize()
  // // we print the width and height to not get an error
  // fmt.Printf("width=%v", "height=%v\n", w, h)
  fpdf.AddPage()
  pdf := PDF{
		fpdf: fpdf,
	}

  skyblue1 := color.RGBA{61, 186 ,213, 250}
	skyblue2 := color.RGBA{41, 166, 193, 220}

  // BANNERS
  // pdf.SetFillColor(61,186,213)
  pdf.Polygon([]gofpdf.PointType{
    {0, 0},
    {0, h / 9.0},
    {w - (w / 6.0), 0},
  }, FillColor(skyblue2))
  // pdf.SetFillColor(41, 166, 193)
  pdf.Polygon([]gofpdf.PointType{
		{w / 6.0, 0},
		{w, 0},
		{w, h / 9.0},
	}, FillColor(skyblue1))
	pdf.Polygon([]gofpdf.PointType{
		{w, h},
		{w, h - h/8.0},
		{w / 6, h},
	}, FillColor(skyblue2))
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - h/8.0},
		{w - (w / 6), h},
	}, FillColor(skyblue1))

// FONTS
fpdf.SetFont("times", "B", 50)
fpdf.SetTextColor(50, 50, 50)
pdf.MoveAbs(0, 100)
_, lineHt := fpdf.GetFontSize()
fpdf.WriteAligned(0, lineHt, "Certificate of Appreciation", gofpdf.AlignCenter)
pdf.Move(0, lineHt*2.0)

fpdf.SetFont("arial", "", 28)
_, lineHt = fpdf.GetFontSize()
fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
pdf.Move(0, lineHt*2.0)

fpdf.SetFont("times", "B", 42)
_, lineHt = fpdf.GetFontSize()
fpdf.WriteAligned(0, lineHt, *name, gofpdf.AlignCenter)
pdf.Move(0, lineHt*1.75)

	fpdf.SetFont("arial", "", 22)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt*1.5, "For extensive hours of volunteer service and going beyond the immediate duties to provide help to those who need the most help in our community.", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*4.5)

	fpdf.ImageOptions("images/FOODONATE.png", w/2.0-60.0, pdf.y+30.0, 130.0, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	pdf.Move(0, 65.0)
	fpdf.SetFillColor(100, 100, 100)
	fpdf.Rect(60.0, pdf.y, 240.0, 1.0, "F")
	fpdf.Rect(490.0, pdf.y, 240.0, 1.0, "F")

	fpdf.SetFont("arial", "", 12)
	pdf.Move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.MoveAbs(60.0+105.0, pdf.y)
	pdf.Text("Date")
	pdf.MoveAbs(490.0+60.0, pdf.y)
	pdf.Text("board of directors")
	pdf.MoveAbs(60.0, pdf.y-lineHt/1.5)
	fpdf.SetFont("times", "", 22)
	_, lineHt = fpdf.GetFontSize()
	pdf.Move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	yr, mo, day := time.Now().Date()
	dateStr := fmt.Sprintf("%d/%d/%d", mo, day, yr)
	fpdf.CellFormat(240.0, lineHt, dateStr, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	pdf.MoveAbs(490.0, pdf.y)
	sig, err := gofpdf.SVGBasicFileParse("images/sig.svg")
	if err != nil {
		panic(err)
	}
	pdf.Move(0, -(sig.Ht*.45 - lineHt))
	fpdf.SVGBasicWrite(&sig, 0.5)


  // pdf.MoveTo(0, 0)
  // pdf.SetFont("arial", "B", 30)
  // // to get exact fontsize
  // _, lineHt := pdf.GetFontSize()
  // pdf.SetTextColor(255, 0, 0)
  // pdf.Text(0, lineHt, "Hello, world")
  // pdf.MoveTo(0, lineHt*2.0)
  //
  //
  // pdf.SetFont("times", "", 18)
  // pdf.SetTextColor(100, 100, 100)
  // _, lineHt = pdf.GetFontSize()
  // pdf.MultiCell(0, lineHt, "Foodonate would like to extend the highest appreciation, on behalf of every member of the team, to", gofpdf.BorderNone, gofpdf.AlignRight, false)
  //
  // // Shapes
  // pdf.SetFillColor(0, 255, 0)
  // pdf.SetDrawColor(0, 0, 255)
  // pdf.Rect(10, 100, 100, 100, "FD")
  // pdf.SetFillColor(100, 200, 200)
  // pdf.Polygon([]gofpdf.PointType{
  //   {110, 250},
  //   {160, 300},
  //   {110, 350},
  //   {60, 300},
  //   {70, 230},
  // }, "F")
  //
  // // printing Foodonate logo
  // pdf.ImageOptions("images/FOODONATE.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
  //   ReadDpi: true,
  //   }, 0, "")

  // // Calling the grid
  // // drawGrid(pdf)
  err = fpdf.OutputFileAndClose("cert.pdf")
  if err != nil {
    panic(err)
  }
}
// making grid to help me find the positions
func drawGrid(pdf *gofpdf.Fpdf) {
  w, h := pdf.GetPageSize()
  pdf.SetFont("courier", "", 12)
  pdf.SetTextColor(80, 80, 80)
  pdf.SetDrawColor(200, 200, 200)
  for x := 0.0; x < w; x = x + (w / 20.0) {
    pdf.SetTextColor(200, 200, 200)
    pdf.Line(x, 0, x, h)
    _, lineHt := pdf.GetFontSize()
    pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
  }
  for y := 0.0; y < h; y = y + (w / 20.0) {
    pdf.SetTextColor(80, 80, 80)
    pdf.Line(0, y, w, y)
    pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
  }
}
