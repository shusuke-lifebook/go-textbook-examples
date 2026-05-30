package main

import (
	"ch13-bookmark-app/internal/repository"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("リクエスト受信", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := sql.Open("sqlite", "bookmarks.db")
	if err != nil {
		slog.Error("DB接続失敗", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.New(db)
	if err := repo.InitTable(); err != nil {
		slog.Error("テーブル作成失敗", "error", err)
		os.Exit(1)
	}

	// TODO サーバー起動の処理

}
