package helper

/**
 * pp: pretty-print helper
 */

import (
	"fmt"
	"strings"
	"time"

	"github.com/linkage-community/wetsuit/entity"
)

func MapStringPerOneLine(s string, a func(s string) string) string {
	r := ""
	for _, s := range strings.Split(s, "\n") {
		r += a(s) + "\n"
	}
	return r[:len(r)-1]
}
func AddLineIndent(s string) string {
	return MapStringPerOneLine(s, func(s string) string { return "\t" + s })
}

func ChooseAlbumFileVariant(a *entity.AlbumFile) *entity.AlbumFileVariant {
	for _, v := range a.Variants {
		if v.Type == "image" {
			return &v
		}
	}
	return nil
}

func UserToString(u *entity.User) string {
	user := ""
	user += fmt.Sprintf("%s @%s [%d]", u.Name, u.ScreenName, u.ID)
	return user
}

func FileToString(f *entity.AlbumFile) string {
	text := fmt.Sprintf("%s ", f.Name)
	if v := ChooseAlbumFileVariant(f); v != nil {
		return text + fmt.Sprintf("[%s]", v.URL)
	}
	return text + "unknown"
}

func PostToString(p entity.Post) string {
	text := ""
	text += "Author:\t" + UserToString(&p.User)
	if t, err := time.Parse(time.RFC3339, p.CreatedAt); err == nil {
		text += "\nDate:\t" + t.Local().Format("2006/01/02 15:04:05 MST")
	}
	text += fmt.Sprintf("\nBy:\t%s [%d bot=%v]", p.Application.Name, p.Application.ID, p.Application.IsAutomated)
	if len(p.Files) != 0 {
		for i, f := range p.Files {
			text += fmt.Sprintf("\nFile#%d:\t%s", i, FileToString(f))
		}
	}

	if body := strings.TrimSpace(p.Text); len(body) > 0 {
		text += "\n\n"
		text += AddLineIndent(body)
		text += "\n"
	}
	return text
}
