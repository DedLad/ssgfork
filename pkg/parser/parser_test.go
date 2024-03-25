package parser_test

import (
	"html/template"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/acmpesuecc/anna/pkg/parser"
)

func TestReadMdDir(t *testing.T) {
	t.Run("reading markdown files and rendering without drafts", func(t *testing.T) {
		p := parser.Parser{
			Templates:   make(map[template.URL]parser.TemplateData),
			TagsMap:     make(map[string][]parser.TemplateData),
			ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		}
		p.RenderDrafts = false

		inpBaseDirFS := os.DirFS("../../test/input")
		p.ReadMdDir("../../test/input", inpBaseDirFS)

		got_parsed_files := len(p.MdFilesName)
		want_parsed_files := 1

		if got_parsed_files != want_parsed_files {
			t.Errorf("got %v, want %v", got_parsed_files, want_parsed_files)
		}

	})

	t.Run("reading all markdown files inluding drafts", func(t *testing.T) {
		p := parser.Parser{
			Templates:   make(map[template.URL]parser.TemplateData),
			TagsMap:     make(map[string][]parser.TemplateData),
			ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		}
		p.RenderDrafts = true

		inpBaseDirFS := os.DirFS("../../test/input")
		p.ReadMdDir("../../test/input", inpBaseDirFS)

		got_parsed_files := len(p.MdFilesName)
		want_parsed_files := 2
		if got_parsed_files != want_parsed_files {
			t.Errorf("got %v, want %v", got_parsed_files, want_parsed_files)
		}
	})
}

func TestAddFileandRender(t *testing.T) {
	got_parser := parser.Parser{
		Templates:   make(map[template.URL]parser.TemplateData),
		TagsMap:     make(map[string][]parser.TemplateData),
		ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	t.Run("parsing data out of one markdown post", func(t *testing.T) {
		inputMd, err := os.ReadFile("../../test/md_inp.md")
		if err != nil {
			t.Errorf("%v", err)
		}
		want_parser := parser.Parser{
			Templates:   make(map[template.URL]parser.TemplateData),
			TagsMap:     make(map[string][]parser.TemplateData),
			ErrorLogger: got_parser.ErrorLogger,
		}
		sample_frontmatter, sample_body, parseSuccess := got_parser.ParseMarkdownContent(string(inputMd))
		if !parseSuccess {
			return
		}

		filename := "testpost.md"
		fileurl := "testpost.html"
		want_parser.MdFilesName = append(want_parser.MdFilesName, filename)
		want_parser.MdFilesPath = append(want_parser.MdFilesPath, filename)
		want_page := parser.TemplateData{
			URL:                      template.URL(fileurl),
			Date:                     want_parser.DateParse(sample_frontmatter.Date).Unix(),
			FilenameWithoutExtension: "testpost",
			Frontmatter:              sample_frontmatter,
			Body:                     template.HTML(sample_body),
			// TODO: Test the Layout contained in the struct
			// Layout: parser.LayoutConfig,
		}

		want_parser.Templates[template.URL("testpost.md")] = want_page
		for _, tag := range sample_frontmatter.Tags {
			want_parser.TagsMap[tag] = append(want_parser.TagsMap[tag], want_page)
		}

		if sample_frontmatter.Type == "post" {
			want_parser.Posts = append(want_parser.Posts, want_page)
		}

		got_parser.AddFileAndRender("", filename, sample_frontmatter, sample_body)

		if !reflect.DeepEqual(got_parser, want_parser) {
			t.Errorf("want %v; \ngot %v", want_parser, got_parser)
		}
	})
}

func TestParseMarkdownContent(t *testing.T) {
	p := parser.Parser{
		Templates:   make(map[template.URL]parser.TemplateData),
		TagsMap:     make(map[string][]parser.TemplateData),
		ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	t.Run("render markdown files to html", func(t *testing.T) {
		inputMd, err := os.ReadFile("../../test/md_inp.md")
		if err != nil {
			t.Errorf("%v", err)
		}

		_, body_got, parseSuccess := p.ParseMarkdownContent(string(inputMd))

		if parseSuccess {

			body_want, err := os.ReadFile("../../test/html_want_output.html")
			if err = os.WriteFile("../../test/html_got_output.html", []byte(body_got), 0666); err != nil {
				t.Errorf("%v", err)
			}
			if string(body_want) != body_got {
				t.Errorf("%v\nThe expected and generated html can be found in test/", err)
			}
		}
	})

	t.Run("parse frontmatter from markdown files", func(t *testing.T) {
		inputMd, err := os.ReadFile("../../test/md_inp.md")
		if err != nil {
			t.Errorf("%v", err)
		}

		frontmatter_got, _, parseSuccess := p.ParseMarkdownContent(string(inputMd))

		if !parseSuccess {
			frontmatter_want := parser.Frontmatter{
				Title:       "Markdown Test",
				Date:        "2024-03-23",
				Description: "File containing markdown to test the SSG",
				Type:        "post",
			}

			if !reflect.DeepEqual(frontmatter_got, frontmatter_want) {
				t.Errorf("got %v, \nwant %v", frontmatter_got, frontmatter_want)
			}
		}
	})
}
