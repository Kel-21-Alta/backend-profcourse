package pdf

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"profcourse/business/users"
	"strconv"
)

type GeneratePDF struct {}

func buildHeading(m pdf.Maroto, user users.Domain) {
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Laporan Perkembangan User", props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Left,
				Color: getDarkPurpleColor(),
			})
		})
	})
	m.Row(10, func() {
		m.Col(2, func() {
			m.Text("Nama: ", props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Left,
				Color: getDarkPurpleColor(),
			})
		})
		m.Col(6, func() {
			m.Text(user.Name, props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Left,
				Color: getDarkPurpleColor(),
			})
		})
	})
	m.Row(10, func() {
		m.Col(2, func() {
			m.Text("Email: ", props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Left,
				Color: getDarkPurpleColor(),
			})
		})
		m.Col(6, func() {
			m.Text(user.Email, props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Left,
				Color: getDarkPurpleColor(),
			})
		})
	})
}
func buildCourseList(m pdf.Maroto, course []users.Course) {
	tableHeadings := []string{"Kursus", "Progres", "Point"}

	contents := [][]string{}

	for _, u := range course {
		contents = append(contents, []string{u.CourseTitle, strconv.Itoa(u.Progres), strconv.Itoa(u.Score)})
	}

	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getBlueColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Kusus Yang Pernah Diikuti", props.Text{
				Top:    2,
				Size:   10,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      11,
			GridSizes: []uint{7, 3, 2},
		},
		ContentProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{7, 3, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   2,
		Line:                 false,
	})
}
func getBlueColor() color.Color {
	return color.Color{
		Red:   11,
		Green: 86,
		Blue:  173,
	}
}
func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}
func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
func (g *GeneratePDF)GeneratePDFDataReport(user users.Domain, course []users.Course) (string, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m, user)
	buildCourseList(m, course)

	var path = 	user.ID + "-" + user.Email + ".pdf"

	err := m.OutputFileAndClose(path)
	if err != nil {
		return "", err
	}
	return path, nil
}


func NewGeneratePDF() users.PDF {
	return &GeneratePDF{}
}