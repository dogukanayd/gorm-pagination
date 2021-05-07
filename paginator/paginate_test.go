package paginator

import (
	"fmt"
	"os"
	"testing"
)

const (
	AppURL = "test.app"
	Path   = "/test-path"
)

func TestConfig_GetAppURL(t *testing.T) {
	t.Run("it_should_return_app_url_from_environment", func(t *testing.T) {
		_ = os.Setenv("APP_URL", AppURL)

		c := Config{}

		au := c.GetAppURL()

		if au != AppURL {
			t.Errorf("GetAppURL() want: %s, got: %s", AppURL, au)
		}
	})

	t.Run("it_should_return_app_url_from_config", func(t *testing.T) {
		AppURL := "test.app"

		c := Config{
			AppURL: AppURL,
		}

		au := c.GetAppURL()

		if au != AppURL {
			t.Errorf("GetAppURL() want: %s, got: %s", AppURL, au)
		}
	})
}

func TestResult_GetLastPage(t *testing.T) {
	t.Run("it_should_return_result_of_total_divided_by_per_page", func(t *testing.T) {
		r := Result{
			Total:   10,
			PerPage: 2,
		}

		lastPage := r.GetLastPage()

		if lastPage != 5 {
			t.Errorf("GetLastPage() want: %d, got: %d", 5, lastPage)
		}
	})
}

func TestConfig_GetPageURL(t *testing.T) {
	t.Run("it_should_get_page_url", func(t *testing.T) {
		c := Config{
			AppURL:  AppURL,
			PerPage: 10,
			Path:    Path,
		}

		pageURL := c.GetPageURL(10)
		want := fmt.Sprintf("%s%s?page=%d&per_page=%d", c.GetAppURL(), c.Path, 10, c.PerPage)

		if want != pageURL {
			t.Errorf("GetPageURL() want: %s, got: %s", want, pageURL)
		}
	})
}

func TestConfig_PreviousPageURL(t *testing.T) {
	t.Run("it_should_return_previous_page_url_when_page_number_greater_than_one", func(t *testing.T) {
		c := Config{
			AppURL:  AppURL,
			PerPage: 10,
			Path:    Path,
			Page:    10,
		}

		previousPage := c.PreviousPageURL()
		want := fmt.Sprintf("%s%s?page=%d&per_page=%d", c.GetAppURL(), c.Path, 9, c.PerPage)

		if want != previousPage {
			t.Errorf("PreviousPageURL() want: %s, got: %s", want, previousPage)
		}
	})
}
