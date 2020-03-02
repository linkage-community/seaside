package helper

/**
 * pp: pretty-print helper
 */

import (
	"fmt"
	"strings"

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
	user += fmt.Sprintf("[%d @%s] %s", u.ID, u.ScreenName, u.Name)
	if u.AvatarFile != nil {
		user += "\n\t.Avatar: "
		user += FileToString(u.AvatarFile)
	}
	return user
}

func FileToString(f *entity.AlbumFile) string {
	text := ""
	text += fmt.Sprintf("[%s] ", f.Name)
	if v := ChooseAlbumFileVariant(f); v != nil {
		text += fmt.Sprintf("%s", v.URL)
		return text
	}
	return text + "unknown"
}

func PostToString(p entity.Post) string {
	text := ".User: "
	text += UserToString(&p.User)
	if len(strings.TrimSpace(p.Text)) == 0 {
		text += "\n.Text: empty"
	} else {
		text += "\n.Text:\n"
		text += AddLineIndent(p.Text)
	}
	text += fmt.Sprintf("\n.Application: [%d] %s bot=%v", p.Application.ID, p.Application.Name, p.Application.IsAutomated)
	if len(p.Files) != 0 {
		text += "\n.Files:"
		for i, f := range p.Files {
			text += fmt.Sprintf("\n\t#%d ", i)
			text += FileToString(f)
		}
	}
	return text
}
