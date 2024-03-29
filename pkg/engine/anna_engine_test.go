package engine_test

import (
	"html/template"
	"log"
	"os"
	"slices"
	"testing"

	"github.com/acmpesuecc/anna/pkg/engine"
	"github.com/acmpesuecc/anna/pkg/parser"
)

func TestRenderTags(t *testing.T) {
	e := engine.Engine{
		Templates:   make(map[template.URL]parser.TemplateData),
		TagsMap:     make(map[string][]parser.TemplateData),
		ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	e.LayoutConfig.BaseURL = "example.org"

	fileOutPath := "../../test/engine/render_tags/"

	e.TagsMap["blogs"] = []parser.TemplateData{
		{
			CompleteURL: "posts/file1.html",
			Frontmatter: parser.Frontmatter{
				Title: "file1",
				Tags:  []string{"blogs"},
			},
		},
		{
			CompleteURL: "posts/file2.html",
			Frontmatter: parser.Frontmatter{
				Title: "file2",
				Tags:  []string{"blogs", "tech"},
			},
		},
	}

	e.TagsMap["tech"] = []parser.TemplateData{
		{
			CompleteURL: "posts/file2.html",
			Frontmatter: parser.Frontmatter{
				Title: "file2",
				Tags:  []string{"blogs", "tech"},
			},
		},
		{
			CompleteURL: "posts/file3.html",
			Frontmatter: parser.Frontmatter{
				Title: "file3",
				Tags:  []string{"tech"},
			},
		},
	}

	templ, err := template.ParseFiles(TestDirPath+"render_tags/tags_template.html", TestDirPath+"render_tags/tags_subpage_template.html")
	if err != nil {
		t.Errorf("%v", err)
	}
	e.RenderTags(fileOutPath, templ)

	t.Run("render tag.html", func(t *testing.T) {

		got_tags_file, err := os.ReadFile(TestDirPath + "render_tags/rendered/tags.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		want_tags_file, err := os.ReadFile(TestDirPath + "render_tags/want_tags.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		if !slices.Equal(got_tags_file, want_tags_file) {
			t.Errorf("The expected and generated tags.html can be found in test/engine/render_tags/rendered/")
		}
	})

	t.Run("render tag-subpage.html", func(t *testing.T) {

		got_blogs_file, err := os.ReadFile(TestDirPath + "render_tags/rendered/tags/blogs.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		want_blogs_file, err := os.ReadFile(TestDirPath + "render_tags/want_blogs_tags.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		if !slices.Equal(got_blogs_file, want_blogs_file) {
			t.Errorf("The expected and generated blogs.html tag-subpage can be found in test/engine/render_tags/rendered/tags/")
		}

		got_tech_file, err := os.ReadFile(TestDirPath + "render_tags/rendered/tags/tech.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		want_tech_file, err := os.ReadFile(TestDirPath + "render_tags/want_tech_tags.html")
		if err != nil {
			t.Errorf("%v", err)
		}

		if !slices.Equal(got_tech_file, want_tech_file) {
			t.Errorf("The expected and generated tech.html tag-subpage can be found in test/engine/render_tags/rendered/tags/")
		}
	})
}

/*
func TestGenerateSitemap(t *testing.T) {
	t.Run("render sitemap.xml", func(t *testing.T) {
		engine := engine.Engine{
			Templates:   make(map[template.URL]parser.TemplateData),
			TagsMap:     make(map[string][]parser.TemplateData),
			ErrorLogger: log.New(os.Stderr, "TEST ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		}

		t1 := parser.TemplateData{
			FilenameWithoutExtension: "index",
			Frontmatter: parser.Frontmatter{
				Date: "2024-02-23",
			},
		}

		t2 := parser.TemplateData{
			FilenameWithoutExtension: "about",
			Frontmatter: parser.Frontmatter{
				Date: "2024-02-23",
			},
		}

		t3 := parser.TemplateData{
			FilenameWithoutExtension: "research",
			Frontmatter: parser.Frontmatter{
				Date: "2024-02-23",
			},
		}

		engine.LayoutConfig.BaseURL = "example.org"
		// setting up engine
		engine.Templates["index"] = t1
		engine.Templates["about"] = t2
		engine.Templates["research"] = t3

		engine.GenerateSitemap(TestDirPath + "sitemap/got_sitemap.xml")

		got_sitemap, err := os.ReadFile(TestDirPath + "sitemap/got_sitemap.xml")
		if err != nil {
			t.Errorf("Error in reading the contents of got_sitemap.xml")
		}

		want_sitemap, err := os.ReadFile(TestDirPath + "sitemap/want_sitemap.xml")
		if err != nil {
			t.Errorf("Error in reading the contents of _sitemap.xml")
		}

		// Remove spaces and whitespace characters

		// got_sitemap = []byte(strings.ReplaceAll(string(got_sitemap), " ", ""))
		// want_sitemap = []byte(strings.ReplaceAll(string(want_sitemap), " ", ""))
		// // Replace all tabs
		// got_sitemap = []byte(strings.ReplaceAll(string(got_sitemap), "\t", ""))
		// want_sitemap = []byte(strings.ReplaceAll(string(want_sitemap), "\t", ""))

		// got_sitemap = []byte(strings.ReplaceAll(string(got_sitemap), "\n", ""))
		// want_sitemap = []byte(strings.ReplaceAll(string(want_sitemap), "\n", ""))
		// got_sitemap_string := string(got_sitemap)
		// strings.TrimFunc(got_sitemap_string, func(r rune) bool {

		// })

		if strings.Compare(string(got_sitemap), string(want_sitemap)) != 0 {
			t.Errorf("The expected and generated sitemap can be found in test/layout/sitemap/")
		}
	})
}
*/
