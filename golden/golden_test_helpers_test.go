package golden

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var updateGolden = flag.Bool("update", false, "обновить golden-файлы")

func assertGolden(t *testing.T, name string, got string) {
	t.Helper()

	path := filepath.Join("testdata", name+".golden")
	if *updateGolden {
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatalf("не удалось обновить golden-файл %s: %v", path, err)
		}
	}

	wantBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("не удалось прочитать golden-файл %s: %v", path, err)
	}

	want := strings.TrimRight(string(wantBytes), "\r\n")
	if got != want {
		t.Errorf("несовпадение с golden-файлом %s\nОжидалось:\n%s\nПолучено:\n%s", path, want, got)
	}
}
