package data

import "github.com/kyoya67/portfolio/backend/internal/model"

var Works = []model.Work{
	{
		ID:          "portfolio-site",
		Title:       "Portfolio Site",
		Description: "Go, Lambda, DynamoDB, Cognito を学ぶための個人ポートフォリオサイト",
		Body:        "個人ポートフォリオサイトの概要。",
		Technologies: []string{
			"Go",
			"TypeScript",
			"Next.js",
			"AWS",
		},
		Published:  true,
		CreatedAt:  "2026-06-24T00:00:00Z",
		UpdatedAt:  "2026-06-24T00:00:00Z",
		ProjectURL: "https://example.com",
	},
}

var Articles = []model.Article{
	{
		ID:          "zenn-go-lambda",
		Source:      "zenn",
		SourceID:    "12345",
		Title:       "Go Lambda の学習メモ",
		URL:         "https://zenn.dev/example/articles/go-lambda",
		Excerpt:     "Go Lambda を個人開発で扱うためのメモ。",
		Tags:        []string{"go", "lambda", "aws"},
		PublishedAt: "2026-06-24T00:00:00Z",
		FetchedAt:   "2026-06-24T00:00:00Z",
		IsPublished: true,
	},
	{
		ID:          "qiita-dynamodb",
		Source:      "qiita",
		SourceID:    "67890",
		Title:       "DynamoDB のアクセスパターン整理",
		URL:         "https://qiita.com/example/items/dynamodb",
		Excerpt:     "DynamoDB は読み方から設計する。",
		Tags:        []string{"dynamodb", "aws"},
		PublishedAt: "2026-06-23T00:00:00Z",
		FetchedAt:   "2026-06-24T00:00:00Z",
		IsPublished: true,
	},
}

var Profile = model.Profile{
	ID:        "main",
	Name:      "Kyoya",
	Bio:       "26歳のソフトウェアエンジニア。Go と AWS を学習中。",
	Skills:    []string{"Go", "TypeScript", "AWS", "DynamoDB"},
	GitHub:    "https://github.com/Kyoya67",
	UpdatedAt: "2026-06-24T00:00:00Z",
}
